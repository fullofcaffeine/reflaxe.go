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
	return New_app__core__PulseReport(self.runtime.profileId(), self.runtime.variantId(), self.runtime.capabilityId(), ingest.receivedCount, ingest.acceptedFrames.Len(), ingest.backpressureEvents, parsed.Len(), enriched.Len(), func(hx_obj_28 map[string]any) *hxrt.Array {
		hx_field_29 := hx_obj_28["sources"]
		if hx_field_29 == nil {
			var hx_zero_30 *hxrt.Array
			return hx_zero_30
		}
		return hx_field_29.(*hxrt.Array)
	}(aggregates).Len(), func(hx_obj_31 map[string]any) int {
		hx_field_32 := hx_obj_31["totalValue"]
		if hx_field_32 == nil {
			var hx_zero_33 int
			return hx_zero_33
		}
		return hx_field_32.(int)
	}(aggregates), func(hx_obj_34 map[string]any) int {
		hx_field_35 := hx_obj_34["totalWeighted"]
		if hx_field_35 == nil {
			var hx_zero_36 int
			return hx_zero_36
		}
		return hx_field_35.(int)
	}(aggregates), func(hx_obj_37 map[string]any) *string {
		hx_field_38 := hx_obj_37["summary"]
		if hx_field_38 == nil {
			var hx_zero_39 *string
			return hx_zero_39
		}
		return hx_field_38.(*string)
	}(aggregates), alerts.Len(), alertDigest, runtimeScore)
}

func (self *app__core__PulsePipeline) ingest(frames *hxrt.Array, capacity int) *app__core__PulseIngestResult {
	var hx_if_40 int
	if capacity <= 0 {
		hx_if_40 = 1
	} else {
		hx_if_40 = capacity
	}
	boundedCapacity := hx_if_40
	queue := hxrt.NewArray()
	queueHead := 0
	accepted := hxrt.NewArray()
	backpressureEvents := 0
	_g := 0
	for _g < frames.Len() {
		frame := func(hx_value_41 any) *app__core__PulseIngressFrame {
			if hx_value_41 == nil {
				var hx_zero_42 *app__core__PulseIngressFrame
				return hx_zero_42
			}
			return hx_value_41.(*app__core__PulseIngressFrame)
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
		entry := func(hx_value_46 any) *app__core__PulseEnrichedEvent {
			if hx_value_46 == nil {
				var hx_zero_47 *app__core__PulseEnrichedEvent
				return hx_zero_47
			}
			return hx_value_46.(*app__core__PulseEnrichedEvent)
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
		summary := func(hx_value_49 any) *app__core__PulseSourceAggregate {
			if hx_value_49 == nil {
				var hx_zero_50 *app__core__PulseSourceAggregate
				return hx_zero_50
			}
			return hx_value_49.(*app__core__PulseSourceAggregate)
		}(sourceSummaries.Get(_g_1))
		_g_1 = int(int32((_g_1 + 1)))
		if !hxrt.StringEqualStringPtr(digest, hxrt.StringFromLiteral("")) {
			digest = hxrt.StringConcatStringPtr(digest, hxrt.StringFromLiteral(","))
		}
		digest = hxrt.StringConcatStringPtr(digest, summary.summaryToken())
	}
	hx_obj_51 := map[string]any{}
	hx_obj_51["sources"] = sourceSummaries
	hx_obj_51["totalValue"] = totalValue
	hx_obj_51["totalWeighted"] = totalWeighted
	hx_obj_51["summary"] = digest
	return hx_obj_51
}

func (self *app__core__PulsePipeline) findSourceAggregate(summaries *hxrt.Array, source *string) *app__core__PulseSourceAggregate {
	_g := 0
	for _g < summaries.Len() {
		summary := func(hx_value_52 any) *app__core__PulseSourceAggregate {
			if hx_value_52 == nil {
				var hx_zero_53 *app__core__PulseSourceAggregate
				return hx_zero_53
			}
			return hx_value_52.(*app__core__PulseSourceAggregate)
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
		entry := func(hx_value_54 any) *app__core__PulseEnrichedEvent {
			if hx_value_54 == nil {
				var hx_zero_55 *app__core__PulseEnrichedEvent
				return hx_zero_55
			}
			return hx_value_54.(*app__core__PulseEnrichedEvent)
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
		digest = hxrt.StringConcatStringPtr(digest, hxrt.StdString(func(hx_value_57 any) *app__core__PulseAlert {
			if hx_value_57 == nil {
				var hx_zero_58 *app__core__PulseAlert
				return hx_zero_58
			}
			return hx_value_57.(*app__core__PulseAlert)
		}(alerts.Get(index)).eventId))
		index = int(int32((index + 1)))
	}
	return digest
}
