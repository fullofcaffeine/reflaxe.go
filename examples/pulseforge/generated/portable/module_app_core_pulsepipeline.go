package main

import "examples_pulseforge_portable/hxrt"

type I_app__core__PulsePipeline interface {
	run(frames []*app__core__PulseIngressFrame) *app__core__PulseReport
	ingest(frames []*app__core__PulseIngressFrame, capacity int) *app__core__PulseIngestResult
	aggregate(enriched []*app__core__PulseEnrichedEvent) map[string]any
	collectAlerts(enriched []*app__core__PulseEnrichedEvent, weightedThreshold int) []*app__core__PulseAlert
	alertToken(alerts []*app__core__PulseAlert) *string
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

func (self *app__core__PulsePipeline) run(frames []*app__core__PulseIngressFrame) *app__core__PulseReport {
	ingest := self.ingest(frames, 3)
	parsed := self.runtime.parse(ingest.acceptedFrames, 1)
	enriched := self.runtime.enrich(parsed, 1)
	aggregates := self.aggregate(enriched)
	_ = aggregates
	alerts := self.collectAlerts(enriched, 20)
	runtimeScore := self.runtime.stageScore(parsed, enriched, alerts, ingest.backpressureEvents)
	_ = runtimeScore
	alertDigest := self.alertToken(alerts)
	return New_app__core__PulseReport(self.runtime.profileId(), self.runtime.variantId(), self.runtime.capabilityId(), ingest.receivedCount, len(ingest.acceptedFrames), ingest.backpressureEvents, len(parsed), len(enriched), len(func(hx_obj_3 map[string]any) []*app__core__PulseSourceAggregate {
		hx_field_4 := hx_obj_3["sources"]
		if hx_field_4 == nil {
			var hx_zero_5 []*app__core__PulseSourceAggregate
			return hx_zero_5
		}
		return hx_field_4.([]*app__core__PulseSourceAggregate)
	}(aggregates)), func(hx_obj_6 map[string]any) int {
		hx_field_7 := hx_obj_6["totalValue"]
		if hx_field_7 == nil {
			var hx_zero_8 int
			return hx_zero_8
		}
		return hx_field_7.(int)
	}(aggregates), func(hx_obj_9 map[string]any) int {
		hx_field_10 := hx_obj_9["totalWeighted"]
		if hx_field_10 == nil {
			var hx_zero_11 int
			return hx_zero_11
		}
		return hx_field_10.(int)
	}(aggregates), func(hx_obj_12 map[string]any) *string {
		hx_field_13 := hx_obj_12["summary"]
		if hx_field_13 == nil {
			var hx_zero_14 *string
			return hx_zero_14
		}
		return hx_field_13.(*string)
	}(aggregates), len(alerts), alertDigest, runtimeScore)
}

func (self *app__core__PulsePipeline) ingest(frames []*app__core__PulseIngressFrame, capacity int) *app__core__PulseIngestResult {
	var hx_if_15 int
	if capacity <= 0 {
		hx_if_15 = 1
	} else {
		hx_if_15 = capacity
	}
	boundedCapacity := hx_if_15
	_ = boundedCapacity
	queue := []*app__core__PulseIngressFrame{}
	_ = queue
	queueHead := 0
	_ = queueHead
	accepted := []*app__core__PulseIngressFrame{}
	_ = accepted
	backpressureEvents := 0
	_ = backpressureEvents
	_g := 0
	for _g < len(frames) {
		frame := frames[_g]
		_ = frame
		_g = int(int32((_g + 1)))
		if int(int32((hxrt.Int32Wrap(len(queue)) - hxrt.Int32Wrap(queueHead)))) >= boundedCapacity {
			backpressureEvents = int(int32((backpressureEvents + 1)))
			accepted = append(accepted, queue[queueHead])
			queueHead = int(int32((queueHead + 1)))
		}
		queue = append(queue, frame)
	}
	for queueHead < len(queue) {
		accepted = append(accepted, queue[queueHead])
		queueHead = int(int32((queueHead + 1)))
	}
	return New_app__core__PulseIngestResult(len(frames), accepted, backpressureEvents)
}

func (self *app__core__PulsePipeline) aggregate(enriched []*app__core__PulseEnrichedEvent) map[string]any {
	bySource := New_haxe__ds__StringMap()
	_ = bySource
	sourceKeys := []*string{}
	_ = sourceKeys
	totalValue := 0
	_ = totalValue
	totalWeighted := 0
	_ = totalWeighted
	_g := 0
	for _g < len(enriched) {
		entry := enriched[_g]
		_ = entry
		_g = int(int32((_g + 1)))
		totalValue = int(int32((hxrt.Int32Wrap(totalValue) + hxrt.Int32Wrap(entry.event.value))))
		totalWeighted = int(int32((hxrt.Int32Wrap(totalWeighted) + hxrt.Int32Wrap(entry.weightedValue))))
		source := entry.event.source
		bucket := func(hx_value_16 any) *app__core__PulseSourceAggregate {
			if hx_value_16 == nil {
				var hx_zero_17 *app__core__PulseSourceAggregate
				return hx_zero_17
			}
			return hx_value_16.(*app__core__PulseSourceAggregate)
		}(bySource.get(source))
		if bucket == nil {
			bucket = New_app__core__PulseSourceAggregate(source)
			bySource.set(source, bucket)
			sourceKeys = append(sourceKeys, source)
		}
		bucket.record(entry)
	}
	sourceSummaries := []*app__core__PulseSourceAggregate{}
	_ = sourceSummaries
	digest := hxrt.StringFromLiteral("")
	_ = digest
	index := 0
	for index < len(sourceKeys) {
		source_1 := sourceKeys[index]
		summary := func(hx_value_18 any) *app__core__PulseSourceAggregate {
			if hx_value_18 == nil {
				var hx_zero_19 *app__core__PulseSourceAggregate
				return hx_zero_19
			}
			return hx_value_18.(*app__core__PulseSourceAggregate)
		}(bySource.get(source_1))
		if summary != nil {
			sourceSummaries = append(sourceSummaries, summary)
			if !hxrt.StringEqualStringPtr(digest, hxrt.StringFromLiteral("")) {
				digest = hxrt.StringConcatStringPtr(digest, hxrt.StringFromLiteral(","))
			}
			digest = hxrt.StringConcatStringPtr(digest, summary.summaryToken())
		}
		index = int(int32((index + 1)))
	}
	hx_obj_20 := map[string]any{}
	hx_obj_20["sources"] = sourceSummaries
	hx_obj_20["totalValue"] = totalValue
	hx_obj_20["totalWeighted"] = totalWeighted
	hx_obj_20["summary"] = digest
	return hx_obj_20
}

func (self *app__core__PulsePipeline) collectAlerts(enriched []*app__core__PulseEnrichedEvent, weightedThreshold int) []*app__core__PulseAlert {
	alerts := []*app__core__PulseAlert{}
	_ = alerts
	_g := 0
	for _g < len(enriched) {
		entry := enriched[_g]
		_ = entry
		_g = int(int32((_g + 1)))
		if entry.weightedValue >= weightedThreshold {
			alerts = append(alerts, app__core__PulseAlert_fromEnriched(entry))
		}
	}
	return alerts
}

func (self *app__core__PulsePipeline) alertToken(alerts []*app__core__PulseAlert) *string {
	if len(alerts) == 0 {
		return hxrt.StringFromLiteral("none")
	}
	digest := hxrt.StringFromLiteral("")
	_ = digest
	index := 0
	for index < len(alerts) {
		if index > 0 {
			digest = hxrt.StringConcatStringPtr(digest, hxrt.StringFromLiteral(","))
		}
		digest = hxrt.StringConcatStringPtr(digest, hxrt.StdString(alerts[index].eventId))
		index = int(int32((index + 1)))
	}
	return digest
}
