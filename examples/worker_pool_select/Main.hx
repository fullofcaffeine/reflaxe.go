import go.Chan;
import go.Go;

class Main {
	// Note: Chan<T> receive values are currently erased to `any` in portable/gopher lanes.
	// We keep `Dynamic` confined to this channel boundary to avoid invalid Go type assertions.
	static function worker(jobs:Chan<Dynamic>, results:Chan<Dynamic>):Void {
		while (true) {
			var job = jobs.recvOr(null);
			if (job == null) {
				return;
			}
			results.send(job);
		}
	}

	static function main() {
		var workerCount = 3;
		var tasks:Array<String> = ["alpha", "beta", "gamma", "delta"];
		var jobs:Chan<Dynamic> = Go.newChan(tasks.length + workerCount);
		var results:Chan<Dynamic> = Go.newChan(tasks.length);

		for (task in tasks) {
			jobs.send(task);
		}
		for (_ in 0...workerCount) {
			jobs.send(null);
		}

		for (_ in 0...workerCount) {
			Go.spawn(function() {
				worker(jobs, results);
			});
		}

		var received = 0;
		while (received < tasks.length) {
			var value = results.recvOr(null);
			if (value == null) {
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
