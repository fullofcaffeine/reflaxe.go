#if macro
import haxe.macro.Context;
import haxe.macro.Type;
import reflaxe.go.compiler.GoCarrierContractIndex;
#end

/**
	Contract for the classes that `GoCompiler.methodUsesCarrierOptionalAbi` walks
	for inherited carrier contracts.

	Why
	`methodUsesCarrierOptionalAbi` walks only the project classes that
	`GoCarrierContractIndex` returns for a method name. A class left out loses a
	carrier signature, and Go then rejects the interface method set. A class kept
	in vain only costs time. BeadsHX has no such contract at all, so its compile
	cannot catch a lost class.

	What
	Pins, for real typed classes, that the collected names equal the old per-name
	walk for every class and name: a contract reached through a parent interface
	or a superclass counts; a method without an optional argument, or an
	optional-argument method declared outside any interface, does not. Also pins
	that interfaces are never returned, and that an appended, replaced, shortened
	or reset class list is read again.

	How
	After the initialization macros, loads the fixture classes, compares the
	index with a copy of the walk it replaced, and checks the lists the index
	returns.
**/
class GoCarrierContractIndexContract {
	#if macro
	public static function check():Void {
		Context.onAfterInitMacros(run);
	}

	static function require(value:Bool, message:String):Void {
		if (!value)
			Context.fatalError("GoCarrierContractIndexContract: " + message, Context.currentPos());
	}

	static function load(path:String):ClassType {
		return switch (Context.getType(path)) {
			case TInst(ref, _): ref.get();
			case _: throw "not a class: " + path;
		};
	}

	static function names(classes:Array<ClassType>):String {
		return [for (source in classes) source.name].join(",");
	}

	static function fullName(classType:ClassType):String {
		return classType.pack.length == 0 ? classType.name : classType.pack.join(".") + "." + classType.name;
	}

	/** The replaced per-name test: `declaresOptionalArgumentMethod` without its cache. */
	static function declaresOptionalArgumentMethod(classType:ClassType, methodName:String):Bool {
		for (field in classType.fields.get()) {
			if (field.name != methodName)
				continue;
			switch (Context.follow(field.type)) {
				case TFun(args, _):
					for (arg in args)
						if (arg.opt)
							return true;
				case _:
			}
		}
		return false;
	}

	/** The replaced per-name walk: `classDeclaresCarrierContract(root, false)`. */
	static function oldContract(root:ClassType, methodName:String):Bool {
		final visitedInterfaces = new Map<String, Bool>();
		function interfaceDeclares(interfaceType:ClassType):Bool {
			final interfaceName = fullName(interfaceType);
			if (visitedInterfaces.exists(interfaceName))
				return false;
			visitedInterfaces.set(interfaceName, true);
			if (declaresOptionalArgumentMethod(interfaceType, methodName))
				return true;
			for (entry in interfaceType.interfaces)
				if (interfaceDeclares(entry.t.get()))
					return true;
			return false;
		}
		function visitClass(current:ClassType):Bool {
			for (entry in current.interfaces)
				if (interfaceDeclares(entry.t.get()))
					return true;
			return current.superClass != null && visitClass(current.superClass.t.get());
		}
		return visitClass(root);
	}

	static function expect(index:GoCarrierContractIndex, classes:Array<ClassType>, methodName:String, expected:String):Void {
		final found = names(index.contractClasses(classes, methodName));
		require(found == expected, "method " + methodName + " gives [" + found + "], expected [" + expected + "]");
	}

	static function run():Void {
		final module = "GoCarrierContractIndexContractFixture";
		final base = load(module);
		final sub = load(module + ".CarrierSubContract");
		final loose = load(module + ".CarrierLooseContract");
		final optional = load(module + ".CarrierOptionalContract");
		final child = load(module + ".CarrierChildContract");
		final other = load(module + ".CarrierOtherContract");
		final index = new GoCarrierContractIndex();

		// The collected names are exactly the names the old walk accepts.
		final methodNames = ["get", "plain", "extra", "other", "missing", "new"];
		for (source in [base, sub, loose, optional, child, other]) {
			final collected = index.contractMethodNames(source);
			for (methodName in methodNames)
				require(collected.contains(methodName) == oldContract(source, methodName),
					source.name
					+ "."
					+ methodName
					+ ": index "
					+ collected.contains(methodName)
					+ ", old walk "
					+ oldContract(source, methodName));
		}
		// Through a parent interface, through a superclass, and never outside an interface.
		require(names([base]) == "GoCarrierContractIndexContractFixture", "fixture name");
		require(index.contractMethodNames(base).join(",") == "extra,get", "base collects " + index.contractMethodNames(base).join(","));
		require(index.contractMethodNames(sub).join(",") == "other,extra,get", "subclass collects " + index.contractMethodNames(sub).join(","));
		require(index.contractMethodNames(loose).length == 0, "a method outside an interface is no contract");

		// Interfaces are never returned, even one whose parent has the contract.
		final classes = [optional, child, base, sub, loose, other];
		expect(index, classes, "get", "GoCarrierContractIndexContractFixture,CarrierSubContract");
		expect(index, classes, "extra", "GoCarrierContractIndexContractFixture,CarrierSubContract");
		expect(index, classes, "other", "CarrierSubContract");
		expect(index, classes, "plain", "");
		expect(index, classes, "missing", "");

		// Appended classes are read on the next call.
		final growing = [loose];
		expect(index, growing, "get", "");
		growing.push(base);
		expect(index, growing, "get", "GoCarrierContractIndexContractFixture");
		growing.push(sub);
		expect(index, growing, "other", "CarrierSubContract");
		expect(index, growing, "get", "GoCarrierContractIndexContractFixture,CarrierSubContract");

		// A different array starts again; so does a shorter one.
		expect(index, [sub], "get", "CarrierSubContract");
		final shrinking = [base, sub];
		expect(index, shrinking, "other", "CarrierSubContract");
		shrinking.resize(1);
		expect(index, shrinking, "other", "");
		expect(index, shrinking, "get", "GoCarrierContractIndexContractFixture");

		// `reset` forgets the list, so a replaced entry is read again.
		final replaced = [base];
		expect(index, replaced, "get", "GoCarrierContractIndexContractFixture");
		replaced[0] = loose;
		index.reset();
		expect(index, replaced, "get", "");
	}
	#end
}
