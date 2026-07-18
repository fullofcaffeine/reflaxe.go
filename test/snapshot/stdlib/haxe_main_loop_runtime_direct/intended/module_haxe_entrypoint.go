package main

func haxe__EntryPoint_addThread(f func()) {
	haxe__EntryPoint_mutex.__hx_this.acquire()
	haxe__EntryPoint_threadCount = int(int32((haxe__EntryPoint_threadCount + 1)))
	haxe__EntryPoint_mutex.__hx_this.release()
	haxe__EntryPoint_mainThread.__hx_this.get_events().__hx_this.promise()
	sys__thread__Thread_create(func() {
		if f != nil {
			f()
		}
		haxe__EntryPoint_mutex.__hx_this.acquire()
		haxe__EntryPoint_threadCount = int(int32((haxe__EntryPoint_threadCount - 1)))
		haxe__EntryPoint_mutex.__hx_this.release()
		haxe__EntryPoint_mainThread.__hx_this.get_events().__hx_this.runPromised(func() {
		})
	})
}

var haxe__EntryPoint_mainThread *sys__thread__Thread = sys__thread__Thread_current()

var haxe__EntryPoint_mutex *sys__thread__Mutex = New_sys__thread__Mutex()

func haxe__EntryPoint_run() {
	haxe__EntryPoint_mainThread.__hx_this.get_events().__hx_this.loop()
}

func haxe__EntryPoint_runInMainThread(f func()) {
	if f != nil {
		haxe__EntryPoint_mainThread.__hx_this.get_events().__hx_this.run(f)
	}
}

var haxe__EntryPoint_threadCount int = 0

func haxe__EntryPoint_wakeup() {
	haxe__EntryPoint_mainThread.__hx_this.get_events().__hx_this.run(func() {
	})
}
