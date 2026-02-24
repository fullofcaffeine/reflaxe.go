import go.Chan;
import go.Go;
import go.Select;

class Main {
	static inline var STOP_TOKEN = "__stop__";
	static inline var EMPTY_TOKEN = "__empty__";

	static function worker(jobs:Chan<String>, results:Chan<String>):Void {
		while (true) {
			var job = jobs.recvOr(STOP_TOKEN);
			if (job == STOP_TOKEN) {
				return;
			}
			results.send(job);
		}
	}

	static function main() {
		var workerCount = 3;
		var tasks:Array<String> = ["alpha", "beta", "gamma", "delta"];
		var jobs:Chan<String> = Go.newChan(tasks.length + workerCount);
		var results:Chan<String> = Go.newChan(tasks.length);

		for (task in tasks) {
			jobs.send(task);
		}
		for (_ in 0...workerCount) {
			jobs.send(STOP_TOKEN);
		}

		for (_ in 0...workerCount) {
			Go.spawn(function() {
				worker(jobs, results);
			});
		}

		var received = 0;
		while (received < tasks.length) {
			var value = results.recvOr(EMPTY_TOKEN);
			if (value == EMPTY_TOKEN) {
				continue;
			}
			received++;
		}

		var selectGate:Chan<Int> = Go.newChan(1);
		var firstTry = switch (Select.send(selectGate, 5)) {
			case Sent:
				true;
			case Defaulted:
				false;
		}
		var secondTry = switch (Select.send(selectGate, 6)) {
			case Sent:
				true;
			case Defaulted:
				false;
		}
		var firstRecv = switch (Select.recv(selectGate)) {
			case Received(value):
				value;
			case Defaulted:
				-1;
		}
		var secondRecv = switch (Select.recv(selectGate)) {
			case Received(value):
				value;
			case Defaulted:
				99;
		}

		var left:Chan<String> = Go.newChan(1);
		var right:Chan<String> = Go.newChan(1);
		right.send("right");
		var recv2 = switch (Select.recv2(left, right)) {
			case First(value):
				"left:" + value;
			case Second(value):
				"right:" + value;
			case Defaulted:
				"none";
		}

		var send2a:Chan<Int> = Go.newChan(1);
		var send2b:Chan<Int> = Go.newChan(1);
		var send2 = switch (Select.send2(send2a, 11, send2b, 22)) {
			case FirstSent:
				"a";
			case SecondSent:
				"b";
			case Defaulted:
				"none";
		}
		var send2Values = send2a.recvOr(-1) + "," + send2b.recvOr(-1);

		Sys.println("worker.count=" + received);
		Sys.println("select.trySend=" + firstTry + "," + secondTry);
		Sys.println("select.recvOr=" + firstRecv + "," + secondRecv);
		Sys.println("select.recv2=" + recv2);
		Sys.println("select.send2=" + send2 + " values=" + send2Values);
	}
}
