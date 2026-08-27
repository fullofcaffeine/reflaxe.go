package go;

/**
	What: A typed view of one native Go `[]string` value.

	Why: Haxe `String` uses a nullable pointer carrier in generated code, while Go
	APIs expose non-nullable string values inside `[]string`. `NativeSlice<String>`
	retains the Haxe pointer carrier and therefore cannot model that ABI exactly.

	How: The compiler stores Go string values, projects Haxe strings at writes, and
	restores the Haxe carrier at reads. `fromArray` makes the representation copy
	explicit so incompatible collection headers never alias.
**/
extern class NativeStringSlice implements ArrayAccess<String> {
	/** The current native slice length. */
	public var length(default, null):Int;

	/** Copies a portable Haxe string Array into native Go `[]string` storage. */
	public static function fromArray(value:Array<String>):NativeStringSlice;
}
