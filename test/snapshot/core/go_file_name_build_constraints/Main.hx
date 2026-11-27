/**
	Modules whose plain file name would end in a Go `_test`, GOOS or GOARCH
	suffix. Go silently leaves such files out of the build, so every module here
	must get a file name that every platform compiles.
**/
class Main {
	static function main() {
		Sys.println(unit.Test.name());
		Sys.println(tools.Linux.name());
		Sys.println(tools.Windows_amd64.name());
		Sys.println(x.Arm64.name());
		Sys.println(x.Js.name());
		Sys.println(x.Unix.name());
	}
}
