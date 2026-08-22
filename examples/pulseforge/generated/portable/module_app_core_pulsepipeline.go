package main

import "examples_pulseforge_portable/hxrt"

type I_app__core__PulsePipeline interface {
	run(frames *hxrt.Array) *app__core__PulseReport
	ingest(frames *hxrt.Array, capacity int) *app__core__PulseIngestResult
	aggregate(enriched *hxrt.Array) map[string]any
	findSourceAggregate(summaries *hxrt.Array, source *string) *app__core__PulseSourceAggregate
	collectAlerts(enriched *hxrt.Array, weightedThreshold int) *hxrt.Array
	alertToken(alerts *hxrt.Array) *string
}

type app__core__PulsePipeline struct {
	__hx_this I_app__core__PulsePipeline
	runtime   app__runtime__PulseRuntime
}

func New_app__core__PulsePipeline(runtime app__runtime__PulseRuntime) *app__core__PulsePipeline {
	self := &app__core__PulsePipeline{}
	self.__hx_this = self
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
	return New_app__core__PulseReport(self.runtime.profileId(), self.runtime.variantId(), self.runtime.capabilityId(), ingest.receivedCount, ingest.acceptedFrames.Len(), ingest.backpressureEvents, parsed.Len(), enriched.Len(), func(hx_obj_1 map[string]any) *hxrt.Array {
		hx_field_2 := hx_obj_1["sources"]
		if hx_field_2 == nil {
			var hx_zero_3 *hxrt.Array
			return hx_zero_3
		}
		return hx_field_2.(*hxrt.Array)
	}(aggregates).Len(), func(hx_obj_4 map[string]any) int {
		hx_field_5 := hx_obj_4["totalValue"]
		if hx_field_5 == nil {
			var hx_zero_6 int
			return hx_zero_6
		}
		return hx_field_5.(int)
	}(aggregates), func(hx_obj_7 map[string]any) int {
		hx_field_8 := hx_obj_7["totalWeighted"]
		if hx_field_8 == nil {
			var hx_zero_9 int
			return hx_zero_9
		}
		return hx_field_8.(int)
	}(aggregates), func(hx_obj_10 map[string]any) *string {
		hx_field_11 := hx_obj_10["summary"]
		if hx_field_11 == nil {
			var hx_zero_12 *string
			return hx_zero_12
		}
		return hx_field_11.(*string)
	}(aggregates), alerts.Len(), alertDigest, runtimeScore)
}

func (self *app__core__PulsePipeline) ingest(frames *hxrt.Array, capacity int) *app__core__PulseIngestResult {
	var hx_if_13 int
	if capacity <= 0 {
		hx_if_13 = 1
	} else {
		hx_if_13 = capacity
	}
	boundedCapacity := hx_if_13
	queue := hxrt.NewArray()
	queueHead := 0
	accepted := hxrt.NewArray()
	backpressureEvents := 0
	_g := 0
	for _g < frames.Len() {
		frame := func(hx_value_14 any) *app__core__PulseIngressFrame {
			if hx_value_14 == nil {
				var hx_zero_15 *app__core__PulseIngressFrame
				return hx_zero_15
			}
			return hx_value_14.(*app__core__PulseIngressFrame)
		}(frames.Get(_g))
		_g = int(int32((_g + 1)))
		if int(int32((hxrt.Int32Wrap(queue.Len()) - hxrt.Int32Wrap(queueHead)))) >= boundedCapacity {
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

func (self *app__core__PulsePipeline) aggregate(enriched *hxrt.Array) map[string]any {
	sourceSummaries := hxrt.NewArray()
	totalValue := 0
	totalWeighted := 0
	_g := 0
	for _g < enriched.Len() {
		entry := func(hx_value_19 any) *app__core__PulseEnrichedEvent {
			if hx_value_19 == nil {
				var hx_zero_20 *app__core__PulseEnrichedEvent
				return hx_zero_20
			}
			return hx_value_19.(*app__core__PulseEnrichedEvent)
		}(enriched.Get(_g))
		_g = int(int32((_g + 1)))
		totalValue = int(int32((hxrt.Int32Wrap(totalValue) + hxrt.Int32Wrap(entry.event.value))))
		totalWeighted = int(int32((hxrt.Int32Wrap(totalWeighted) + hxrt.Int32Wrap(entry.weightedValue))))
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
		summary := func(hx_value_22 any) *app__core__PulseSourceAggregate {
			if hx_value_22 == nil {
				var hx_zero_23 *app__core__PulseSourceAggregate
				return hx_zero_23
			}
			return hx_value_22.(*app__core__PulseSourceAggregate)
		}(sourceSummaries.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		if !hxrt.StringEqualStringPtr(digest, hxrt.StringFromLiteral("")) {
			digest = hxrt.StringConcatStringPtr(digest, hxrt.StringFromLiteral(","))
		}
		digest = hxrt.StringConcatStringPtr(digest, summary.summaryToken())
	}
	hx_obj_24 := map[string]any{}
	hx_obj_24["sources"] = sourceSummaries
	hx_obj_24["totalValue"] = totalValue
	hx_obj_24["totalWeighted"] = totalWeighted
	hx_obj_24["summary"] = digest
	return hx_obj_24
}

func (self *app__core__PulsePipeline) findSourceAggregate(summaries *hxrt.Array, source *string) *app__core__PulseSourceAggregate {
	_g := 0
	for _g < summaries.Len() {
		summary := func(hx_value_25 any) *app__core__PulseSourceAggregate {
			if hx_value_25 == nil {
				var hx_zero_26 *app__core__PulseSourceAggregate
				return hx_zero_26
			}
			return hx_value_25.(*app__core__PulseSourceAggregate)
		}(summaries.Get(_g))
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
		entry := func(hx_value_27 any) *app__core__PulseEnrichedEvent {
			if hx_value_27 == nil {
				var hx_zero_28 *app__core__PulseEnrichedEvent
				return hx_zero_28
			}
			return hx_value_27.(*app__core__PulseEnrichedEvent)
		}(enriched.Get(_g))
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
		digest = hxrt.StringConcatStringPtr(digest, hxrt.StdString(func(hx_value_30 any) *app__core__PulseAlert {
			if hx_value_30 == nil {
				var hx_zero_31 *app__core__PulseAlert
				return hx_zero_31
			}
			return hx_value_30.(*app__core__PulseAlert)
		}(alerts.Get(index)).eventId))
		index = int(int32((index + 1)))
	}
	return digest
}
