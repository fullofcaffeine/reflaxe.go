package haxe;

/**
	Single event registered with `haxe.MainLoop` on `haxe.go`.

	What
	Stores one callback, its blocking flag, priority, and next scheduled run time.

	Why
	Upstream `MainLoop` keeps a linked list that a target-specific entry point
	drives. Go already has a runtime-backed event loop, so this wrapper keeps the
	Haxe-facing event object but schedules dispatch through `haxe.EntryPoint`.

	How
	A dispatch callback checks whether the event is still active, waits until
	`nextRun` if `delay()` moved it into the future, calls the user callback, and
	requeues itself while the event remains active.
**/
@:allow(haxe.MainLoop)
class MainEvent {
	var f:Void->Void;
	var timer:Null<Timer>;
	var active:Bool = true;

	public var isBlocking:Bool = true;
	public var nextRun(default, null):Float;
	public var priority(default, null):Int;

	function new(f:Void->Void, priority:Int) {
		this.f = f;
		this.priority = priority;
		nextRun = -1.0;
	}

	public function delay(t:Float):Void {
		nextRun = Timer.stamp() + t;
		if (timer != null) {
			timer.stop();
			timer = null;
		}
		schedule();
	}

	public inline function call():Void {
		if (f != null) {
			f();
		}
	}

	public function stop():Void {
		if (!active) {
			return;
		}
		active = false;
		f = null;
		if (timer != null) {
			timer.stop();
			timer = null;
		}
		MainLoop.__remove(this);
	}

	function delayNow():Void {
		nextRun = -1.0;
		if (timer != null) {
			timer.stop();
			timer = null;
		}
		schedule();
	}

	function schedule():Void {
		if (!active) {
			return;
		}
		var wait = nextRun - Timer.stamp();
		if (wait <= 0) {
			EntryPoint.runInMainThread(dispatch);
		} else {
			var ms = Math.ceil(wait * 1000);
			if (ms < 1) {
				ms = 1;
			}
			timer = Timer.delay(dispatch, ms);
		}
	}

	function dispatch():Void {
		if (!active || f == null) {
			return;
		}
		var wait = nextRun - Timer.stamp();
		if (wait > 0) {
			schedule();
			return;
		}
		call();
		if (active && f != null) {
			schedule();
		}
	}
}

/**
	Main-loop facade for direct `haxe.MainLoop` use on `haxe.go`.

	What
	Supports adding callbacks, stopping them through `MainEvent`, and scheduling
	work onto the main thread.

	Why
	Direct `haxe.MainLoop` was previously blocked because it compiled without a
	real Go event-loop owner. The owner now exists in `sys.thread.EventLoop`, so
	the public facade can be source-owned again.

	How
	The facade tracks active events for `hasEvents()` / blocking semantics and uses
	`haxe.EntryPoint` to enqueue actual callback dispatch on the main event loop.
**/
class MainLoop {
	static var pending = new Array<MainEvent>();

	public static var threadCount(get, never):Int;

	static inline function get_threadCount():Int {
		return EntryPoint.threadCount;
	}

	public static function hasEvents():Bool {
		for (event in pending) {
			if (event.isBlocking) {
				return true;
			}
		}
		return false;
	}

	public static function addThread(f:Void->Void):Void {
		EntryPoint.addThread(f);
	}

	public static function runInMainThread(f:Void->Void):Void {
		EntryPoint.runInMainThread(f);
	}

	public static function add(f:Void->Void, priority = 0):MainEvent {
		if (f == null) {
			throw "Event function is null";
		}
		var event = new MainEvent(f, priority);
		pending.push(event);
		event.delayNow();
		return event;
	}

	public static function __remove(event:MainEvent):Void {
		var next = new Array<MainEvent>();
		for (candidate in pending) {
			if (candidate != event) {
				next.push(candidate);
			}
		}
		pending = next;
	}
}
