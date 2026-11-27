package x;

/**
	Control: `unix` is a build tag, but Go does not read it from file names, so
	this module keeps its plain `module_x_unix.go` name.
**/
class Unix {
	public static function name():String {
		return "x.Unix";
	}
}
