package main

import (
	"fmt"
	"strings"
)

const (
	alertWeightedThreshold = 20
	ingestQueueCapacity    = 3
)

type pulseIngressFrame struct {
	Sequence int
	Source   string
	Value    int
	Region   string
}

type pulseEvent struct {
	ID     int
	Source string
	Region string
	Value  int
}

type pulseEnrichedEvent struct {
	Event         pulseEvent
	Severity      int
	WeightedValue int
}

func (e pulseEnrichedEvent) shouldAlert(weightedThreshold int) bool {
	return e.WeightedValue >= weightedThreshold
}

type pulseAlert struct {
	EventID       int
	Source        string
	Region        string
	Severity      int
	WeightedValue int
	Reason        string
}

func pulseAlertFromEnriched(entry pulseEnrichedEvent) pulseAlert {
	label := "warning"
	if entry.Severity >= 3 {
		label = "critical"
	}
	return pulseAlert{
		EventID:       entry.Event.ID,
		Source:        entry.Event.Source,
		Region:        entry.Event.Region,
		Severity:      entry.Severity,
		WeightedValue: entry.WeightedValue,
		Reason:        label,
	}
}

type pulseSourceAggregate struct {
	Source        string
	Count         int
	TotalValue    int
	TotalWeighted int
	MaxValue      int
	MaxSeverity   int
}

func newPulseSourceAggregate(source string) *pulseSourceAggregate {
	return &pulseSourceAggregate{
		Source: source,
	}
}

func (s *pulseSourceAggregate) record(entry pulseEnrichedEvent) {
	s.Count++
	s.TotalValue += entry.Event.Value
	s.TotalWeighted += entry.WeightedValue
	if entry.Event.Value > s.MaxValue {
		s.MaxValue = entry.Event.Value
	}
	if entry.Severity > s.MaxSeverity {
		s.MaxSeverity = entry.Severity
	}
}

func (s *pulseSourceAggregate) summaryToken() string {
	return fmt.Sprintf("%s:%d/%d/%d/sev%d", s.Source, s.Count, s.TotalValue, s.TotalWeighted, s.MaxSeverity)
}

type pulseIngestResult struct {
	ReceivedCount     int
	AcceptedFrames    []pulseIngressFrame
	BackpressureEvent int
}

type pulseReport struct {
	Profile            string
	Variant            string
	Capability         string
	IngestReceived     int
	IngestAccepted     int
	IngestBackpressure int
	ParseEvents        int
	EnrichEvents       int
	AggregateSources   int
	AggregateTotal     int
	AggregateWeighted  int
	AggregateSummary   string
	AlertCount         int
	AlertEvents        string
	RuntimeScore       int
}

func (r pulseReport) lines() []string {
	return []string{
		"pulseforge.profile=" + r.Profile,
		"pulseforge.variant=" + r.Variant,
		"runtime.capability=" + r.Capability,
		"ingest.received=" + fmt.Sprintf("%d", r.IngestReceived),
		"ingest.accepted=" + fmt.Sprintf("%d", r.IngestAccepted),
		"ingest.backpressure=" + fmt.Sprintf("%d", r.IngestBackpressure),
		"parse.events=" + fmt.Sprintf("%d", r.ParseEvents),
		"enrich.events=" + fmt.Sprintf("%d", r.EnrichEvents),
		"aggregate.sources=" + fmt.Sprintf("%d", r.AggregateSources),
		"aggregate.total=" + fmt.Sprintf("%d", r.AggregateTotal),
		"aggregate.weighted_total=" + fmt.Sprintf("%d", r.AggregateWeighted),
		"aggregate.summary=" + r.AggregateSummary,
		"alert.count=" + fmt.Sprintf("%d", r.AlertCount),
		"alert.events=" + r.AlertEvents,
		"runtime.score=" + fmt.Sprintf("%d", r.RuntimeScore),
	}
}

func (r pulseReport) render() string {
	return strings.Join(r.lines(), "\n")
}

type pulseRuntime interface {
	profileID() string
	variantID() string
	capabilityID() string
	parse(frames []pulseIngressFrame, workerCount int) []pulseEvent
	enrich(events []pulseEvent, workerCount int) []pulseEnrichedEvent
	stageScore(parsed []pulseEvent, enriched []pulseEnrichedEvent, alerts []pulseAlert, backpressureEvents int) int
	parseWorkers() int
	enrichWorkers() int
}

type coreRuntime struct {
	profile string
}

func (r coreRuntime) profileID() string {
	return r.profile
}

func (r coreRuntime) variantID() string {
	return "core"
}

func (r coreRuntime) capabilityID() string {
	return "core_loop"
}

func (r coreRuntime) parseWorkers() int {
	return 1
}

func (r coreRuntime) enrichWorkers() int {
	return 1
}

func (r coreRuntime) parse(frames []pulseIngressFrame, _ int) []pulseEvent {
	parsed := make([]pulseEvent, 0, len(frames))
	for _, frame := range frames {
		parsed = append(parsed, pulseParse(frame))
	}
	return parsed
}

func (r coreRuntime) enrich(events []pulseEvent, _ int) []pulseEnrichedEvent {
	enriched := make([]pulseEnrichedEvent, 0, len(events))
	for _, event := range events {
		enriched = append(enriched, pulseEnrich(event))
	}
	return enriched
}

func (r coreRuntime) stageScore(parsed []pulseEvent, enriched []pulseEnrichedEvent, alerts []pulseAlert, backpressureEvents int) int {
	score := 0
	for _, entry := range enriched {
		score += entry.WeightedValue
	}
	score += len(alerts) * 5
	score -= backpressureEvents * 2
	score += len(parsed)
	return score
}

type goNativeRuntime struct {
	profile string
}

func (r goNativeRuntime) profileID() string {
	return r.profile
}

func (r goNativeRuntime) variantID() string {
	return "go_native"
}

func (r goNativeRuntime) capabilityID() string {
	return "chan_fanout_select"
}

func (r goNativeRuntime) parseWorkers() int {
	return 3
}

func (r goNativeRuntime) enrichWorkers() int {
	return 2
}

func (r goNativeRuntime) parse(frames []pulseIngressFrame, workerCount int) []pulseEvent {
	if len(frames) == 0 {
		return []pulseEvent{}
	}

	workers := normalizeWorkers(workerCount)
	inbox := make(chan pulseIngressFrame, len(frames))
	out := make(chan pulseEvent, len(frames))
	done := make(chan int, workers)

	for _, frame := range frames {
		inbox <- frame
	}
	close(inbox)

	for i := 0; i < workers; i++ {
		go func() {
			processed := 0
			for frame := range inbox {
				out <- pulseParse(frame)
				processed++
			}
			done <- processed
		}()
	}

	waitForWorkers(done, workers)

	return orderParsed(drainEvents(out, len(frames)), len(frames))
}

func (r goNativeRuntime) enrich(events []pulseEvent, workerCount int) []pulseEnrichedEvent {
	if len(events) == 0 {
		return []pulseEnrichedEvent{}
	}

	workers := normalizeWorkers(workerCount)
	inbox := make(chan pulseEvent, len(events))
	out := make(chan pulseEnrichedEvent, len(events))
	done := make(chan int, workers)

	for _, event := range events {
		inbox <- event
	}
	close(inbox)

	for i := 0; i < workers; i++ {
		go func() {
			processed := 0
			for event := range inbox {
				out <- pulseEnrich(event)
				processed++
			}
			done <- processed
		}()
	}

	waitForWorkers(done, workers)

	return orderEnriched(drainEnriched(out, len(events)), len(events))
}

func (r goNativeRuntime) stageScore(parsed []pulseEvent, enriched []pulseEnrichedEvent, alerts []pulseAlert, backpressureEvents int) int {
	inbox := make(chan int, len(enriched))
	score := 0

	for _, entry := range enriched {
		select {
		case inbox <- entry.WeightedValue:
			score += 2
		default:
			score -= 25
		}
	}

	for i := 0; i < len(enriched); i++ {
		select {
		case value := <-inbox:
			score += value
		default:
			score += <-inbox
		}
	}

	close(inbox)
	score += len(alerts) * 7
	score -= backpressureEvents * 3
	score += len(parsed)
	return score
}

func normalizeWorkers(workerCount int) int {
	if workerCount <= 0 {
		return 1
	}
	return workerCount
}

func waitForWorkers(done chan int, workers int) {
	for i := 0; i < workers; i++ {
		<-done
	}
	close(done)
}

func drainEvents(out chan pulseEvent, expected int) []pulseEvent {
	parsed := make([]pulseEvent, 0, expected)
	for i := 0; i < expected; i++ {
		parsed = append(parsed, <-out)
	}
	close(out)
	return parsed
}

func drainEnriched(out chan pulseEnrichedEvent, expected int) []pulseEnrichedEvent {
	enriched := make([]pulseEnrichedEvent, 0, expected)
	for i := 0; i < expected; i++ {
		enriched = append(enriched, <-out)
	}
	close(out)
	return enriched
}

func orderParsed(items []pulseEvent, expected int) []pulseEvent {
	byID := make(map[int]pulseEvent, len(items))
	for _, item := range items {
		byID[item.ID] = item
	}

	ordered := make([]pulseEvent, 0, expected)
	for id := 1; id <= expected; id++ {
		if event, ok := byID[id]; ok {
			ordered = append(ordered, event)
		}
	}
	return ordered
}

func orderEnriched(items []pulseEnrichedEvent, expected int) []pulseEnrichedEvent {
	byID := make(map[int]pulseEnrichedEvent, len(items))
	for _, item := range items {
		byID[item.Event.ID] = item
	}

	ordered := make([]pulseEnrichedEvent, 0, expected)
	for id := 1; id <= expected; id++ {
		if event, ok := byID[id]; ok {
			ordered = append(ordered, event)
		}
	}
	return ordered
}

func pulseParse(frame pulseIngressFrame) pulseEvent {
	source := normalizeToken(frame.Source, "unknown")
	region := normalizeToken(frame.Region, "global")
	value := frame.Value
	if value < 0 {
		value = 0
	}
	return pulseEvent{
		ID:     frame.Sequence,
		Source: source,
		Region: region,
		Value:  value,
	}
}

func pulseEnrich(event pulseEvent) pulseEnrichedEvent {
	severity := severityFor(event.Value)
	weightedValue := event.Value*severity + regionBoost(event.Region)
	return pulseEnrichedEvent{
		Event:         event,
		Severity:      severity,
		WeightedValue: weightedValue,
	}
}

func severityFor(value int) int {
	if value >= 12 {
		return 3
	}
	if value >= 8 {
		return 2
	}
	return 1
}

func regionBoost(region string) int {
	// Keep parity with current Haxe->Go PulseForge output contract.
	// Region boost is effectively neutral in the generated example pipeline today.
	_ = region
	return 0
}

func normalizeToken(value, fallback string) string {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return fallback
	}
	return trimmed
}

func baselineFrames() []pulseIngressFrame {
	return []pulseIngressFrame{
		{Sequence: 1, Source: "edge", Value: 3, Region: "iad"},
		{Sequence: 2, Source: "api", Value: 7, Region: "sfo"},
		{Sequence: 3, Source: "db", Value: 11, Region: "fra"},
		{Sequence: 4, Source: "edge", Value: 2, Region: "iad"},
		{Sequence: 5, Source: "auth", Value: 13, Region: "gru"},
		{Sequence: 6, Source: "worker", Value: 5, Region: "sfo"},
		{Sequence: 7, Source: "api", Value: 9, Region: "fra"},
		{Sequence: 8, Source: "db", Value: 4, Region: "iad"},
	}
}

func cloneFrames(frames []pulseIngressFrame) []pulseIngressFrame {
	out := make([]pulseIngressFrame, len(frames))
	copy(out, frames)
	return out
}

func ingest(frames []pulseIngressFrame, capacity int) pulseIngestResult {
	boundedCapacity := capacity
	if boundedCapacity <= 0 {
		boundedCapacity = 1
	}
	queue := make([]pulseIngressFrame, 0, len(frames))
	queueHead := 0
	accepted := make([]pulseIngressFrame, 0, len(frames))
	backpressureEvents := 0

	for _, frame := range frames {
		if len(queue)-queueHead >= boundedCapacity {
			backpressureEvents++
			accepted = append(accepted, queue[queueHead])
			queueHead++
		}
		queue = append(queue, frame)
	}

	for queueHead < len(queue) {
		accepted = append(accepted, queue[queueHead])
		queueHead++
	}

	return pulseIngestResult{
		ReceivedCount:     len(frames),
		AcceptedFrames:    accepted,
		BackpressureEvent: backpressureEvents,
	}
}

func aggregate(enriched []pulseEnrichedEvent) (sources []*pulseSourceAggregate, totalValue int, totalWeighted int, summary string) {
	bySource := make(map[string]*pulseSourceAggregate, len(enriched))
	sourceKeys := make([]string, 0, len(enriched))

	for _, entry := range enriched {
		totalValue += entry.Event.Value
		totalWeighted += entry.WeightedValue

		source := entry.Event.Source
		bucket, ok := bySource[source]
		if !ok {
			bucket = newPulseSourceAggregate(source)
			bySource[source] = bucket
			sourceKeys = append(sourceKeys, source)
		}
		bucket.record(entry)
	}

	sources = make([]*pulseSourceAggregate, 0, len(sourceKeys))
	digest := make([]string, 0, len(sourceKeys))
	for _, source := range sourceKeys {
		item := bySource[source]
		if item == nil {
			continue
		}
		sources = append(sources, item)
		digest = append(digest, item.summaryToken())
	}
	summary = strings.Join(digest, ",")
	return sources, totalValue, totalWeighted, summary
}

func collectAlerts(enriched []pulseEnrichedEvent, weightedThreshold int) []pulseAlert {
	alerts := make([]pulseAlert, 0, len(enriched))
	for _, entry := range enriched {
		if entry.shouldAlert(weightedThreshold) {
			alerts = append(alerts, pulseAlertFromEnriched(entry))
		}
	}
	return alerts
}

func alertToken(alerts []pulseAlert) string {
	if len(alerts) == 0 {
		return "none"
	}
	ids := make([]string, 0, len(alerts))
	for _, alert := range alerts {
		ids = append(ids, fmt.Sprintf("%d", alert.EventID))
	}
	return strings.Join(ids, ",")
}

func runReport(runtime pulseRuntime, frames []pulseIngressFrame) pulseReport {
	ingestResult := ingest(frames, ingestQueueCapacity)
	parsed := runtime.parse(ingestResult.AcceptedFrames, runtime.parseWorkers())
	enriched := runtime.enrich(parsed, runtime.enrichWorkers())
	aggregates, totalValue, totalWeighted, aggregateSummary := aggregate(enriched)
	alerts := collectAlerts(enriched, alertWeightedThreshold)
	runtimeScore := runtime.stageScore(parsed, enriched, alerts, ingestResult.BackpressureEvent)
	alertDigest := alertToken(alerts)

	return pulseReport{
		Profile:            runtime.profileID(),
		Variant:            runtime.variantID(),
		Capability:         runtime.capabilityID(),
		IngestReceived:     ingestResult.ReceivedCount,
		IngestAccepted:     len(ingestResult.AcceptedFrames),
		IngestBackpressure: ingestResult.BackpressureEvent,
		ParseEvents:        len(parsed),
		EnrichEvents:       len(enriched),
		AggregateSources:   len(aggregates),
		AggregateTotal:     totalValue,
		AggregateWeighted:  totalWeighted,
		AggregateSummary:   aggregateSummary,
		AlertCount:         len(alerts),
		AlertEvents:        alertDigest,
		RuntimeScore:       runtimeScore,
	}
}

func runScripted(runtime pulseRuntime) string {
	return runReport(runtime, cloneFrames(baselineFrames())).render()
}

func newRuntime(profile, variant string) (pulseRuntime, error) {
	switch variant {
	case "core":
		return coreRuntime{profile: profile}, nil
	case "go_native":
		return goNativeRuntime{profile: profile}, nil
	default:
		return nil, fmt.Errorf("unsupported variant: %s", variant)
	}
}
