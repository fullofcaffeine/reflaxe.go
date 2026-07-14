import haxe.io.Eof;
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

	static function main() {
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

		Sys.println("detached.throws=" + throws(() -> {
			new Process("sh", ["-c", "exit 0"], true);
		}));
	}
}
