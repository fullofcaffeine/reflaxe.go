import go.Chan;
import go.Go;

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
		var firstTry = selectGate.trySend(5);
		var secondTry = selectGate.trySend(6);
		var firstRecv = selectGate.recvOr(-1);
		var secondRecv = selectGate.recvOr(99);

		Sys.println("worker.count=" + received);
		Sys.println("select.trySend=" + firstTry + "," + secondTry);
		Sys.println("select.recvOr=" + firstRecv + "," + secondRecv);
	}
}
