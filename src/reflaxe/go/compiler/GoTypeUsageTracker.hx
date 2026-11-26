package reflaxe.go.compiler;

#if (macro || reflaxe_runtime)
import haxe.ds.ReadOnlyArray;
import haxe.macro.Context;
import haxe.macro.Type;
import haxe.macro.TypedExprTools;
import reflaxe.compiler.TypeUsageTracker.TypeOrModuleType;
import reflaxe.compiler.TypeUsageTracker.TypeUsageLevel;
import reflaxe.compiler.TypeUsageTracker.TypeUsageMap;
import reflaxe.helpers.TypeHelper;

using reflaxe.helpers.ModuleTypeHelper;

/**
	One type that a module type uses, with every level seen for it so far.

	Why / What / How
	- The map keeps the first value seen for an id, so later values with the
	  same id only add levels.
	- `levels` is a bit set of `TypeUsageLevel` values.
**/
private typedef GoTypeUsageEntry = {
	final usage:TypeOrModuleType;
	var levels:Int;
}

/**
	One field of a class with its decoded body.

	Why
	The tracker and the type-usage ledger both walk every field body of a
	project class. `ClassField.expr()` builds the whole typed body again as new
	macro values on every call, so decoding it once for both halves that cost.

	What
	`field` as `ClassType.constructor`, `fields`, or `statics` returned it, and
	`body`, the value of `field.expr()`.

	How
	`GoTypeUsageTracker.readClassFields(...)` reads them; the tracker publishes
	the list it walked as `GoTypeUsageTracker.classFields`.
**/
typedef GoTrackedClassField = {
	final field:ClassField;
	final body:Null<TypedExpr>;
}

/**
	Builds Reflaxe's `TypeUsageMap` of a module type, reading each enum, typedef,
	and abstract declaration only once per compilation.

	Why
	Reflaxe's `TypeUsageTracker` names every type it records by an id from
	`ModuleTypeHelper.getUniqueId(...)`, which reads the declaration with
	`Ref.get()`. It does this for every expression node and every type argument,
	and reading an abstract builds its whole declaration again (fields,
	operators, casts, metadata). On the BeadsHX bootstrap compile that was
	213,862 reads and about 3.3 s, of which abstracts took 2.2 s, although the
	program has only 667 distinct enums, typedefs, and abstracts.

	What
	`track(moduleType)` returns the map that Reflaxe's
	`TypeUsageTracker.trackTypesInModuleType(moduleType)` returns: the same
	levels and, in each level, the same entries in the same order.
	`GoTypeUsageTrackerContract` compares the two. The rules are Reflaxe's:
	- every expression records its type at `Expression`, except `new`
	  (`Constructed`) and variable declarations (`VariableType`); a static field
	  access also records its class at `StaticAccess`;
	- a variable field records its type at `VariableDeclaration`; a method and
	  an enum constructor record their argument and return types at
	  `FunctionDeclaration`; a superclass, interface, typedef target, or followed
	  abstract records at `ExtendedFrom`;
	- recording a type records its declaration, and then, at the same level,
	  each function argument and return type, each anonymous field type, and
	  each type argument. A function or anonymous type is also an entry itself,
	  named by its printed form;
	- an abstract seen for the first time at a level also records its followed
	  underlying type at that level.
	Reflaxe adds `haxe.Function` for function types only after
	`TypeUsageTracker.init()`, which neither Reflaxe nor haxe.go calls, so
	neither tracker adds it.

	How
	A declaration id depends only on the declaration's package, module, and
	name, which never change after Haxe creates it. Haxe gives each enum,
	typedef, and abstract of a compilation its own type path, and private ones a
	module-specific path (`pack._Module.Name`), so the path that
	`GoTypeMapper.declarationPath(...)` reads without decoding names one
	declaration. `declarationIds` keeps those ids by kind and path. Classes are
	always read: type parameter classes print alike (each local `pick<V>` owns a
	`pick.V`, and BeadsHX has two different `append.T`), so a class path does not
	name one declaration. The compiler creates a new tracker for every
	compilation, so no id survives into the next compilation on a compilation
	server.
**/
class GoTypeUsageTracker {
	/** Ids of enums (`E`), typedefs (`T`), and abstracts (`A`), by kind letter and declaration path. */
	final declarationIds:Map<String, String> = [];

	/** Entries of the module type being tracked, by id. */
	var entries:Map<String, GoTypeUsageEntry> = [];

	/**
		The fields and bodies of the class the last `track(...)` call walked, in
		the order it walked them, or null when that module type was not a class.
		The type-usage ledger reads them instead of decoding the bodies again.
	**/
	public var classFields(default, null):Null<ReadOnlyArray<GoTrackedClassField>> = null;

	public function new() {}

	/**
		Read the constructor, instance fields, and static fields of a class, in
		that order, with each body decoded once.
	**/
	public static function readClassFields(classType:ClassType):Array<GoTrackedClassField> {
		final result:Array<GoTrackedClassField> = [];
		if (classType.constructor != null)
			result.push(trackedField(classType.constructor.get()));
		for (field in classType.fields.get())
			result.push(trackedField(field));
		for (field in classType.statics.get())
			result.push(trackedField(field));
		return result;
	}

	static function trackedField(field:ClassField):GoTrackedClassField {
		return {field: field, body: field.expr()};
	}

	/** Build the type-usage map of one module type. */
	public function track(moduleType:ModuleType):TypeUsageMap {
		entries = [];
		classFields = null;
		switch (moduleType) {
			case TClassDecl(classRef):
				final classType = classRef.get();
				if (classType.superClass != null)
					recordExtended(classType.superClass.t, classType.superClass.params);
				for (implemented in classType.interfaces)
					recordExtended(implemented.t, implemented.params);
				final fields = readClassFields(classType);
				for (tracked in fields)
					recordField(tracked);
				classFields = fields;
			case TEnumDecl(enumRef):
				final enumType = enumRef.get();
				// Visit constructors in the map's own key order, as Reflaxe does.
				for (name in enumType.constructs.keys()) {
					final construct = enumType.constructs.get(name);
					if (construct != null)
						recordSignature(construct.type);
				}
			case TTypeDecl(typeRef):
				recordType(typeRef.get().type, TypeUsageLevel.ExtendedFrom);
			case TAbstract(_):
				recordType(Context.followWithAbstracts(TypeHelper.fromModuleType(moduleType)), TypeUsageLevel.ExtendedFrom);
		}
		return usageMap();
	}

	function recordExtended(classRef:Ref<ClassType>, parameters:Array<Type>):Void {
		recordModuleType(TClassDecl(classRef), TypeUsageLevel.ExtendedFrom);
		for (parameter in parameters)
			recordType(parameter, TypeUsageLevel.ExtendedFrom);
	}

	/** Record a field's body first, then its declared type. */
	function recordField(tracked:GoTrackedClassField):Void {
		final field = tracked.field;
		final body = tracked.body;
		if (body != null)
			recordExpression(body);
		switch (field.kind) {
			case FVar(_, _):
				recordType(field.type, TypeUsageLevel.VariableDeclaration);
			case FMethod(_):
				recordSignature(field.type);
		}
	}

	/** Record the argument and return types of a function type, but not the function type itself. */
	function recordSignature(type:Type):Void {
		switch (type) {
			case TFun(arguments, returnType):
				for (argument in arguments)
					recordType(argument.t, TypeUsageLevel.FunctionDeclaration);
				recordType(returnType, TypeUsageLevel.FunctionDeclaration);
			case _:
		}
	}

	/** Record one expression, then its subexpressions in `TypedExprTools.iter(...)` order. */
	function recordExpression(expr:TypedExpr):Void {
		switch (expr.expr) {
			case TField(_, FStatic(classRef, _)):
				recordType(TypeHelper.fromModuleType(TClassDecl(classRef)), TypeUsageLevel.StaticAccess);
			case _:
		}
		switch (expr.expr) {
			case TNew(classRef, parameters, _):
				recordType(TInst(classRef, parameters), TypeUsageLevel.Constructed);
			case TVar(variable, _):
				recordType(variable.t, TypeUsageLevel.VariableType);
			case _:
				recordType(expr.t, TypeUsageLevel.Expression);
		}
		TypedExprTools.iter(expr, recordExpression);
	}

	/**
		Record a type's declaration, the type itself when it is a function or an
		anonymous structure, its function arguments and return type or its
		anonymous fields, and its type arguments, all at `level`.
	**/
	function recordType(type:Null<Type>, level:TypeUsageLevel):Void {
		if (type == null)
			return;
		final declaration = TypeHelper.toModuleType(type);
		if (declaration != null)
			recordModuleType(declaration, level);
		switch (type) {
			case TFun(arguments, returnType):
				record(TypeHelper.getUniqueId(type), EType(type), level);
				for (argument in arguments)
					recordType(argument.t, level);
				recordType(returnType, level);
			case TAnonymous(anonymousRef):
				record(TypeHelper.getUniqueId(type), EType(type), level);
				for (field in anonymousRef.get().fields)
					recordType(field.type, level);
			case _:
		}
		final parameters = TypeHelper.getParams(type);
		if (parameters != null)
			for (parameter in parameters)
				recordType(parameter, level);
	}

	/** Record a declaration; an abstract new at this level also records its followed type. */
	function recordModuleType(moduleType:ModuleType, level:TypeUsageLevel):Void {
		if (!record(declarationId(moduleType), EModuleType(moduleType), level))
			return;
		switch (moduleType) {
			case TAbstract(_):
				recordType(Context.followWithAbstracts(TypeHelper.fromModuleType(moduleType)), level);
			case _:
		}
	}

	/** Add `level` to the entry `id`, creating it with `usage`; true when the entry or the level is new. */
	function record(id:String, usage:TypeOrModuleType, level:Int):Bool {
		final entry = entries.get(id);
		if (entry == null) {
			entries.set(id, {usage: usage, levels: level});
			return true;
		}
		if ((entry.levels & level) != 0)
			return false;
		entry.levels |= level;
		return true;
	}

	/**
		`ModuleTypeHelper.getUniqueId(moduleType)`, read once per enum, typedef,
		or abstract declaration and every time for a class (see the class notes).
	**/
	function declarationId(moduleType:ModuleType):String {
		final key = switch (moduleType) {
			case TClassDecl(_): null;
			case TEnumDecl(enumRef): "E" + GoTypeMapper.declarationPath(enumRef);
			case TTypeDecl(typeRef): "T" + GoTypeMapper.declarationPath(typeRef);
			case TAbstract(abstractRef): "A" + GoTypeMapper.declarationPath(abstractRef);
		};
		if (key == null)
			return moduleType.getUniqueId();
		var id = declarationIds.get(key);
		if (id == null) {
			id = moduleType.getUniqueId();
			declarationIds.set(key, id);
		}
		return id;
	}

	/**
		Every level, in increasing order, each with its entries in the key order of
		`entries`. Reflaxe builds its map the same way from a map with the same ids
		added in the same order, so both maps list the entries in the same order.
	**/
	function usageMap():TypeUsageMap {
		final usage:TypeUsageMap = [];
		for (index in 0...TypeUsageLevel.LevelCount())
			usage.set(1 << index, []);
		for (id in entries.keys()) {
			final entry = entries.get(id);
			if (entry == null)
				continue;
			for (index in 0...TypeUsageLevel.LevelCount()) {
				final level = 1 << index;
				final values = usage.get(level);
				if ((entry.levels & level) != 0 && values != null)
					values.push(entry.usage);
			}
		}
		return usage;
	}
}
#end
