package main

import "examples_pulseforge_portable/hxrt"

type I_app__core__PulsePipeline interface {
	run(frames []*app__core__PulseIngressFrame) *app__core__PulseReport
	ingest(frames []*app__core__PulseIngressFrame, capacity int) *app__core__PulseIngestResult
	aggregate(enriched []*app__core__PulseEnrichedEvent) map[string]any
	findSourceAggregate(summaries []*app__core__PulseSourceAggregate, source *string) *app__core__PulseSourceAggregate
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
	alerts := self.collectAlerts(enriched, 20)
	runtimeScore := self.runtime.stageScore(parsed, enriched, alerts, ingest.backpressureEvents)
	alertDigest := self.alertToken(alerts)
	return New_app__core__PulseReport(self.runtime.profileId(), self.runtime.variantId(), self.runtime.capabilityId(), ingest.receivedCount, len(ingest.acceptedFrames), ingest.backpressureEvents, len(parsed), len(enriched), len(func(hx_obj_13 map[string]any) []*app__core__PulseSourceAggregate {
		hx_field_14 := hx_obj_13["sources"]
		if hx_field_14 == nil {
			var hx_zero_15 []*app__core__PulseSourceAggregate
			return hx_zero_15
		}
		return hx_field_14.([]*app__core__PulseSourceAggregate)
	}(aggregates)), func(hx_obj_16 map[string]any) int {
		hx_field_17 := hx_obj_16["totalValue"]
		if hx_field_17 == nil {
			var hx_zero_18 int
			return hx_zero_18
		}
		return hx_field_17.(int)
	}(aggregates), func(hx_obj_19 map[string]any) int {
		hx_field_20 := hx_obj_19["totalWeighted"]
		if hx_field_20 == nil {
			var hx_zero_21 int
			return hx_zero_21
		}
		return hx_field_20.(int)
	}(aggregates), func(hx_obj_22 map[string]any) *string {
		hx_field_23 := hx_obj_22["summary"]
		if hx_field_23 == nil {
			var hx_zero_24 *string
			return hx_zero_24
		}
		return hx_field_23.(*string)
	}(aggregates), len(alerts), alertDigest, runtimeScore)
}

func (self *app__core__PulsePipeline) ingest(frames []*app__core__PulseIngressFrame, capacity int) *app__core__PulseIngestResult {
	var hx_if_25 int
	if capacity <= 0 {
		hx_if_25 = 1
	} else {
		hx_if_25 = capacity
	}
	boundedCapacity := hx_if_25
	queue := []*app__core__PulseIngressFrame{}
	queueHead := 0
	accepted := []*app__core__PulseIngressFrame{}
	backpressureEvents := 0
	_g := 0
	for _g < len(frames) {
		frame := frames[_g]
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
	sourceSummaries := []*app__core__PulseSourceAggregate{}
	totalValue := 0
	totalWeighted := 0
	_g := 0
	for _g < len(enriched) {
		entry := enriched[_g]
		_g = int(int32((_g + 1)))
		totalValue = int(int32((hxrt.Int32Wrap(totalValue) + hxrt.Int32Wrap(entry.event.value))))
		totalWeighted = int(int32((hxrt.Int32Wrap(totalWeighted) + hxrt.Int32Wrap(entry.weightedValue))))
		source := entry.event.source
		bucket := self.findSourceAggregate(sourceSummaries, source)
		if bucket == nil {
			bucket = New_app__core__PulseSourceAggregate(source)
			sourceSummaries = append(sourceSummaries, bucket)
		}
		bucket.record(entry)
	}
	digest := hxrt.StringFromLiteral("")
	_g_1 := 0
	for _g_1 < len(sourceSummaries) {
		summary := sourceSummaries[_g_1]
		_g_1 = int(int32((_g_1 + 1)))
		if !hxrt.StringEqualStringPtr(digest, hxrt.StringFromLiteral("")) {
			digest = hxrt.StringConcatStringPtr(digest, hxrt.StringFromLiteral(","))
		}
		digest = hxrt.StringConcatStringPtr(digest, summary.summaryToken())
	}
	hx_obj_26 := map[string]any{}
	hx_obj_26["sources"] = sourceSummaries
	hx_obj_26["totalValue"] = totalValue
	hx_obj_26["totalWeighted"] = totalWeighted
	hx_obj_26["summary"] = digest
	return hx_obj_26
}

func (self *app__core__PulsePipeline) findSourceAggregate(summaries []*app__core__PulseSourceAggregate, source *string) *app__core__PulseSourceAggregate {
	_g := 0
	for _g < len(summaries) {
		summary := summaries[_g]
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(summary.source, source) {
			return summary
		}
	}
	return nil
}

func (self *app__core__PulsePipeline) collectAlerts(enriched []*app__core__PulseEnrichedEvent, weightedThreshold int) []*app__core__PulseAlert {
	alerts := []*app__core__PulseAlert{}
	_g := 0
	for _g < len(enriched) {
		entry := enriched[_g]
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
