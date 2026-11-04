import CgoFeature;
import PureGoFeature;

class Main {
	static function main():Void {
		if (Registry.selected != "ok")
			throw "selected recovery did not run";
		Sys.println(Registry.selected);
	}
}
