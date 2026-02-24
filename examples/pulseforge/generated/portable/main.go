package main

import "examples_pulseforge_portable/hxrt"

func main() {
	var runtime app__runtime__PulseRuntime = app__runtime__RuntimeFactory_create()
	pipeline := New_app__core__PulsePipeline(runtime)
	report := pipeline.run(workload())
	_ = report
	_g := 0
	_ = _g
	_g1 := report.lines()
	for _g < len(_g1) {
		line := _g1[_g]
		_ = line
		_g = int(int32((_g + 1)))
		hxrt.Println(line)
	}
}

func workload() []*app__core__PulseEvent {
	return []*app__core__PulseEvent{New_app__core__PulseEvent(1, hxrt.StringFromLiteral("edge"), 3), New_app__core__PulseEvent(2, hxrt.StringFromLiteral("api"), 7), New_app__core__PulseEvent(3, hxrt.StringFromLiteral("db"), 11), New_app__core__PulseEvent(4, hxrt.StringFromLiteral("edge"), 2), New_app__core__PulseEvent(5, hxrt.StringFromLiteral("auth"), 13), New_app__core__PulseEvent(6, hxrt.StringFromLiteral("worker"), 5)}
}
