package main

import "examples_pulseforge_portable/hxrt"

type I_app__core__PulsePipeline interface {
	run(frames *hxrt.Array) *app__core__PulseReport
	ingest(frames *hxrt.Array, capacity int) *app__core__PulseIngestResult
	aggregate(enriched *hxrt.Array) hxrt.Anonymous
	findSourceAggregate(summaries *hxrt.Array, source *string) *app__core__PulseSourceAggregate
	collectAlerts(enriched *hxrt.Array, weightedThreshold int) *hxrt.Array
	alertToken(alerts *hxrt.Array) *string
}

type app__core__PulsePipeline struct {
	__hx_this I_app__core__PulsePipeline
	runtime   app__runtime__PulseRuntime
}

func (self *app__core__PulsePipeline) HxFunctionReceiver() any {
	return self.__hx_this
}

func New_app__core__PulsePipeline(runtime app__runtime__PulseRuntime) *app__core__PulsePipeline {
	return hx_init_app__core__PulsePipeline(hx_alloc_app__core__PulsePipeline(nil), runtime)
}

func hx_alloc_app__core__PulsePipeline(owner any) *app__core__PulsePipeline {
	self := &app__core__PulsePipeline{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I_app__core__PulsePipeline)
	}
	return self
}

func hx_init_app__core__PulsePipeline(self *app__core__PulsePipeline, runtime app__runtime__PulseRuntime) *app__core__PulsePipeline {
	self.runtime = runtime
	return self
}

func (self *app__core__PulsePipeline) run(frames *hxrt.Array) *app__core__PulseReport {
	ingest := self.ingest(frames, 3)
	parsed := self.runtime.parse(ingest.acceptedFrames, 1)
	enriched := self.runtime.enrich(parsed, 1)
	aggregates := self.aggregate(enriched)
	alerts := self.collectAlerts(enriched, 20)
	runtimeScore := self.runtime.stageScore(parsed, enriched, alerts, ingest.backpressureEvents)
	alertDigest := self.alertToken(alerts)
	return New_app__core__PulseReport(self.runtime.profileId(), self.runtime.variantId(), self.runtime.capabilityId(), ingest.receivedCount, ingest.acceptedFrames.Len(), ingest.backpressureEvents, parsed.Len(), enriched.Len(), func(hx_value_1 any) *hxrt.Array {
		if hx_value_1 == nil {
			var hx_zero_2 *hxrt.Array
			return hx_zero_2
		}
		return hx_value_1.(*hxrt.Array)
	}(hxrt.AnonymousGet(aggregates, "sources")).Len(), func(hx_value_3 any) int {
		if hx_value_3 == nil {
			var hx_zero_4 int
			return hx_zero_4
		}
		return hx_value_3.(int)
	}(hxrt.AnonymousGet(aggregates, "totalValue")), func(hx_value_5 any) int {
		if hx_value_5 == nil {
			var hx_zero_6 int
			return hx_zero_6
		}
		return hx_value_5.(int)
	}(hxrt.AnonymousGet(aggregates, "totalWeighted")), func(hx_value_7 any) *string {
		if hx_value_7 == nil {
			var hx_zero_8 *string
			return hx_zero_8
		}
		return hx_value_7.(*string)
	}(hxrt.AnonymousGet(aggregates, "summary")), alerts.Len(), alertDigest, runtimeScore)
}

func (self *app__core__PulsePipeline) ingest(frames *hxrt.Array, capacity int) *app__core__PulseIngestResult {
	var hx_if_9 int
	if capacity <= 0 {
		hx_if_9 = 1
	} else {
		hx_if_9 = capacity
	}
	boundedCapacity := hx_if_9
	queue := hxrt.NewArray()
	queueHead := 0
	accepted := hxrt.NewArray()
	backpressureEvents := 0
	_g := 0
	for _g < frames.Len() {
		frame := hx_recover_class_0(frames.Get(_g))
		_g = int(int32((_g + 1)))
		if int((hxrt.Int32Wrap(queue.Len()) - hxrt.Int32Wrap(queueHead))) >= boundedCapacity {
			backpressureEvents = int(int32((backpressureEvents + 1)))
			accepted.Push(queue.Get(queueHead))
			queueHead = int(int32((queueHead + 1)))
		}
		queue.Push(frame)
	}
	for queueHead < queue.Len() {
		accepted.Push(queue.Get(queueHead))
		queueHead = int(int32((queueHead + 1)))
	}
	return New_app__core__PulseIngestResult(frames.Len(), accepted, backpressureEvents)
}

func (self *app__core__PulsePipeline) aggregate(enriched *hxrt.Array) hxrt.Anonymous {
	sourceSummaries := hxrt.NewArray()
	totalValue := 0
	totalWeighted := 0
	_g := 0
	for _g < enriched.Len() {
		entry := hx_recover_class_1(enriched.Get(_g))
		_g = int(int32((_g + 1)))
		totalValue = int((hxrt.Int32Wrap(totalValue) + hxrt.Int32Wrap(entry.event.value)))
		totalWeighted = int((hxrt.Int32Wrap(totalWeighted) + hxrt.Int32Wrap(entry.weightedValue)))
		source := entry.event.source
		bucket := self.findSourceAggregate(sourceSummaries, source)
		if bucket == nil {
			bucket = New_app__core__PulseSourceAggregate(source)
			sourceSummaries.Push(bucket)
		}
		bucket.record(entry)
	}
	digest := hxrt.StringFromLiteral("")
	_g_1 := 0
	for _g_1 < sourceSummaries.Len() {
		summary := hx_recover_class_2(sourceSummaries.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		if !hxrt.StringEqualStringPtr(digest, hxrt.StringFromLiteral("")) {
			digest = hxrt.StringConcatStringPtr(digest, hxrt.StringFromLiteral(","))
		}
		digest = hxrt.StringConcatStringPtr(digest, summary.summaryToken())
	}
	hx_obj_14 := map[string]any{}
	hx_obj_14["sources"] = sourceSummaries
	hx_obj_14["totalValue"] = totalValue
	hx_obj_14["totalWeighted"] = totalWeighted
	hx_obj_14["summary"] = digest
	return hxrt.AnonymousFromMap(hx_obj_14)
}

func (self *app__core__PulsePipeline) findSourceAggregate(summaries *hxrt.Array, source *string) *app__core__PulseSourceAggregate {
	_g := 0
	for _g < summaries.Len() {
		summary := hx_recover_class_2(summaries.Get(_g))
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(summary.source, source) {
			return summary
		}
	}
	return nil
}

func (self *app__core__PulsePipeline) collectAlerts(enriched *hxrt.Array, weightedThreshold int) *hxrt.Array {
	alerts := hxrt.NewArray()
	_g := 0
	for _g < enriched.Len() {
		entry := hx_recover_class_1(enriched.Get(_g))
		_g = int(int32((_g + 1)))
		if entry.weightedValue >= weightedThreshold {
			alerts.Push(app__core__PulseAlert_fromEnriched(entry))
		}
	}
	return alerts
}

func (self *app__core__PulsePipeline) alertToken(alerts *hxrt.Array) *string {
	if alerts.Len() == 0 {
		return hxrt.StringFromLiteral("none")
	}
	digest := hxrt.StringFromLiteral("")
	index := 0
	for index < alerts.Len() {
		if index > 0 {
			digest = hxrt.StringConcatStringPtr(digest, hxrt.StringFromLiteral(","))
		}
		digest = hxrt.StringConcatStringPtr(digest, hxrt.StdString(hx_recover_class_3(alerts.Get(index)).eventId))
		index = int(int32((index + 1)))
	}
	return digest
}
