import haxe.io.Eof;
import sys.FileSystem;
import sys.io.Process;

class Main {
	static function throws(action:() -> Void):Bool {
		try {
			action();
			return false;
		} catch (_:Dynamic) {
			return true;
		}
		return false;
	}

	static function reachesEof(process:Process):Bool {
		try {
			process.stdout.readLine();
			return false;
		} catch (_:Eof) {
			return true;
		} catch (_:Dynamic) {
			return false;
		}
		return false;
	}

	static function waitForChild(seconds:String):Void {
		var waiter = new Process("sh", ["-c", "sleep " + seconds]);
		waiter.exitCode();
		waiter.close();
	}

	static function main() {
		Sys.println("startup.throws=" + throws(() -> {
			new Process("__haxe_go_missing_process__", []);
		}));

		var empty = new Process("sh", ["-c", "printf '\\n'"]);
		Sys.println("empty.line=" + (empty.stdout.readLine() == ""));
		Sys.println("empty.eof=" + reachesEof(empty));
		Sys.println("empty.code=" + empty.exitCode());
		empty.close();

		var shell = new Process("printf 'shell-form\\n'");
		Sys.println("shell.line=" + shell.stdout.readLine());
		Sys.println("shell.code=" + shell.exitCode());
		shell.close();

		var piped = new Process("sh", [
			"-c",
			"IFS= read -r line; printf 'out:%s\\n' \"$line\"; printf 'err:%s\\n' \"$line\" >&2; exit 7"
		]);
		Sys.println("pid.positive=" + (piped.getPid() > 0));
		piped.stdin.writeString("hello\n");
		piped.stdin.close();
		Sys.println("stdin.stdout=" + piped.stdout.readLine());
		Sys.println("stderr.line=" + piped.stderr.readLine());
		Sys.println("exit.code=" + piped.exitCode());
		piped.close();

		var longOutput = new Process("python3", ["-c", "print('x' * 70000)"]);
		Sys.println("long.length=" + longOutput.stdout.readLine().length);
		Sys.println("long.code=" + longOutput.exitCode());
		longOutput.close();

		var nonblocking = new Process("sh", ["-c", "sleep 0.2; exit 9"]);
		Sys.println("nonblock.running=" + (nonblocking.exitCode(false) == null));
		waitForChild("0.3");
		var nonblockingCode = nonblocking.exitCode(false);
		Sys.println("nonblock.code=" + nonblockingCode);
		nonblocking.close();

		var killed = new Process("sh", ["-c", "sleep 5"]);
		killed.kill();
		Sys.println("kill.nonzero=" + (killed.exitCode() != 0));
		killed.close();

		var marker = "tmp_process_close_marker.txt";
		if (FileSystem.exists(marker)) {
			FileSystem.deleteFile(marker);
		}
		var closing = new Process("sh", ["-c", "sleep 0.2; printf done > " + marker]);
		closing.close();
		waitForChild("0.5");
		Sys.println("close.keeps.running=" + FileSystem.exists(marker));
		if (FileSystem.exists(marker)) {
			FileSystem.deleteFile(marker);
		}

		Sys.println("detached.throws=" + throws(() -> {
			new Process("sh", ["-c", "exit 0"], true);
		}));
	}
}
