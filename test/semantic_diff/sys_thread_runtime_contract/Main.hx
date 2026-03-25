import sys.thread.ElasticThreadPool;
import sys.thread.FixedThreadPool;
import sys.thread.EventLoop;
import sys.thread.NoEventLoopException;
import sys.thread.Thread;
import sys.thread.ThreadPoolException;
import sys.thread.EventLoop.EventHandler;

class Main {
	static function sortedMessages(count:Int):String {
		var values = new Array<String>();
		for (_ in 0...count) {
			values.push(Std.string(Thread.readMessage(true)));
		}
		haxe.ds.ArraySort.sort(values, Reflect.compare);
		var buf = new StringBuf();
		for (index => value in values) {
			if (index > 0) {
				buf.add(",");
			}
			buf.add(value);
		}
		return buf.toString();
	}

	static function main() {
		var mainThread = Thread.current();

		var worker = Thread.create(function() {
			mainThread.sendMessage("worker-ready");
			var payload = Thread.readMessage(true);
			mainThread.sendMessage("worker-echo=" + Std.string(payload));
		});
		Sys.println("thread.msg1=" + Std.string(Thread.readMessage(true)));
		worker.sendMessage("ping");
		Sys.println("thread.msg2=" + Std.string(Thread.readMessage(true)));

		try {
			worker.events.progress();
			Sys.println("thread.worker_events=available");
		} catch (err:NoEventLoopException) {
			Sys.println("thread.worker_events=" + err.message);
		}

		var runLoop = new EventLoop();
		runLoop.run(function() Sys.println("loop.run=once"));
		runLoop.loop();
		Sys.println("loop.loop_after=done");

		var promisedLoop = new EventLoop();
		promisedLoop.promise();
		promisedLoop.runPromised(function() Sys.println("loop.runPromised=ok"));
		promisedLoop.loop();
		Sys.println("loop.promised_after=done");

		var loop = new EventLoop();
		var repeats = 0;
		var handler:EventHandler = cast 0;
		handler = loop.repeat(function() {
			repeats++;
			Sys.println("loop.repeat=" + repeats);
			loop.cancel(handler);
		}, 10);
		loop.loop();
		Sys.println("loop.repeat_after=" + repeats);

		Thread.create(function() {
			Thread.runWithEventLoop(function() {
				Thread.current().events.promise();
				Thread.current().events.runPromised(function() {
					mainThread.sendMessage("runWithEventLoop=ok");
				});
			});
			mainThread.sendMessage("runWithEventLoop=after");
		});
		Sys.println("thread.runWithEventLoop=" + sortedMessages(2));

		Thread.createWithEventLoop(function() {
			Thread.current().events.promise();
			Thread.current().events.runPromised(function() {
				mainThread.sendMessage("createWithEventLoop=ok");
			});
			mainThread.sendMessage("createWithEventLoop=after-job");
		});
		Sys.println("thread.createWithEventLoop=" + sortedMessages(2));

		var fixed = new FixedThreadPool(2);
		fixed.run(function() mainThread.sendMessage("fixed:b"));
		fixed.run(function() mainThread.sendMessage("fixed:a"));
		Sys.println("fixed.msgs=" + sortedMessages(2));
		fixed.shutdown();
		Sys.println("fixed.shutdown=" + fixed.isShutdown);
		try {
			fixed.run(function() {});
		} catch (err:ThreadPoolException) {
			Sys.println("fixed.error=" + err.message);
		}

		var elastic = new ElasticThreadPool(2, 0.05);
		elastic.run(function() mainThread.sendMessage("elastic:a"));
		elastic.run(function() mainThread.sendMessage("elastic:b"));
		Sys.println("elastic.msgs=" + sortedMessages(2));
		elastic.shutdown();
		Sys.println("elastic.shutdown=" + elastic.isShutdown);
		try {
			elastic.run(function() {});
		} catch (err:ThreadPoolException) {
			Sys.println("elastic.error=" + err.message);
		}
	}
}
