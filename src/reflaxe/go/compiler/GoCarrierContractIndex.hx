package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.Type.ClassType;

/**
	What: Finds the project classes whose declared interfaces give a method name
	an optional-argument contract, for `GoCompiler.methodUsesCarrierOptionalAbi`.

	Why: `methodUsesCarrierOptionalAbi` must know whether some other project class
	satisfies an interface method with optional arguments through a method that the
	asked class defines. It asked that question of every project class, and each
	question walked the class's interface graph and superclass chain again,
	reading each declaration with `Ref.get()`. The compiler asks for every method
	it lowers, and the structural recovery emit asks again in each recovery pass.
	On the BeadsHX bootstrap compile that was about 5,000 questions of about 210
	classes each, about 2.5 s of a 45 s compile. For almost every method name, no
	class has such a contract.

	How: For each class, one walk collects every method name that its interface
	graph declares with an optional argument. The walk is the same as the per-name
	walk it replaces: the class's interfaces and their parent interfaces, then the
	same for each superclass through the declared `superClass` links, with each
	interface visited once by its full name. A name is in the collected list
	exactly when that per-name walk finds it, because an interface declares a name
	with an optional argument exactly when one of its fields of that name is a
	function with an optional argument. The index maps each name to the classes
	that collect it, in list order, and skips interfaces as the caller did.

	These facts are fixed for a class: they read only its typed declarations, so
	the index can keep them for the whole compilation. The class list may only
	grow by appending: the index reads appended classes on the next call, and
	starts again when it is given a different array or a shorter one.
**/
class GoCarrierContractIndex {
	var indexedClasses:Null<Array<ClassType>> = null;
	var indexedCount = 0;
	final classesByMethod:Map<String, Array<ClassType>> = [];
	final optionalMethodsByInterface:Map<String, Array<String>> = [];

	public function new() {}

	/** Forget everything; a new compilation brings new class objects. */
	public function reset():Void {
		indexedClasses = null;
		indexedCount = 0;
		classesByMethod.clear();
		optionalMethodsByInterface.clear();
	}

	/**
		The classes of `classes` that are not interfaces and whose interface graph,
		through their superclass chains, declares `methodName` with an optional
		argument, in their order in `classes`.
	**/
	public function contractClasses(classes:Array<ClassType>, methodName:String):Array<ClassType> {
		if (classes != indexedClasses || classes.length < indexedCount) {
			classesByMethod.clear();
			indexedClasses = classes;
			indexedCount = 0;
		}
		while (indexedCount < classes.length) {
			final source = classes[indexedCount++];
			if (source.isInterface)
				continue;
			for (name in contractMethodNames(source)) {
				var owners = classesByMethod.get(name);
				if (owners == null) {
					owners = [];
					classesByMethod.set(name, owners);
				}
				owners.push(source);
			}
		}
		final owners = classesByMethod.get(methodName);
		return owners == null ? [] : owners;
	}

	/**
		Every method name that an interface of `source`, of a parent interface or of
		a superclass declares with an optional argument, each name once.
	**/
	public function contractMethodNames(source:ClassType):Array<String> {
		final names = new Array<String>();
		final collected = new Map<String, Bool>();
		final visitedInterfaces = new Map<String, Bool>();
		function visitInterface(interfaceType:ClassType):Void {
			final interfaceName = fullName(interfaceType);
			if (visitedInterfaces.exists(interfaceName))
				return;
			visitedInterfaces.set(interfaceName, true);
			for (name in optionalArgumentMethodNames(interfaceType))
				if (!collected.exists(name)) {
					collected.set(name, true);
					names.push(name);
				}
			for (entry in interfaceType.interfaces)
				visitInterface(entry.t.get());
		}
		var current:Null<ClassType> = source;
		while (current != null) {
			for (entry in current.interfaces)
				visitInterface(entry.t.get());
			current = current.superClass == null ? null : current.superClass.t.get();
		}
		return names;
	}

	/**
		The names of the fields of `interfaceType` whose type is a function with an
		optional argument, kept per declaration (see `GoTypeMapper.uniqueClassPath`).
	**/
	function optionalArgumentMethodNames(interfaceType:ClassType):Array<String> {
		final key = GoTypeMapper.uniqueClassPath(interfaceType);
		if (key != null) {
			final cached = optionalMethodsByInterface.get(key);
			if (cached != null)
				return cached;
		}
		final names = new Array<String>();
		for (field in interfaceType.fields.get())
			switch (Context.follow(field.type)) {
				case TFun(args, _):
					for (arg in args)
						if (arg.opt) {
							names.push(field.name);
							break;
						}
				case _:
			}
		if (key != null)
			optionalMethodsByInterface.set(key, names);
		return names;
	}

	/** The name `GoCompiler.fullClassName` gives a class: its package path and name. */
	static function fullName(classType:ClassType):String {
		return classType.pack.length == 0 ? classType.name : classType.pack.join(".") + "." + classType.name;
	}
}
#end
