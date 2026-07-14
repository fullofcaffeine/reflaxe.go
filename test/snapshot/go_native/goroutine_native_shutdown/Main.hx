import go.Chan;
import go.Go;

class Main {
	static function main():Void {
		var blocked:Chan<Bool> = Go.newChan();
		Go.spawn(() -> {
			blocked.recv();
			Sys.println("late-native-goroutine");
		});
		Sys.println("main-only");
	}
}
