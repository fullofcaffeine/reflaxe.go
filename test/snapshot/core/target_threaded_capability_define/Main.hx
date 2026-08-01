#if !target.threaded
#error "haxe.go must declare target.threaded before ordinary dependencies are typed"
#end
class Main {
	static function main() {
		trace("target-threaded");
	}
}
