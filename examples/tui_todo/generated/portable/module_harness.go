package main

import "examples_tui_todo_portable/hxrt"

func Harness_assertContract(runtime profile__TodoRuntime) *string {
	app := New_app__TodoApp(runtime)
	Harness_runBaseline(app)
	baseline := app.baselineSignature()
	if !hxrt.StringEqualStringPtr(baseline, hxrt.StringFromLiteral("open=1,done=1,total=2")) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("baseline drift: "), baseline))
	}
	added := app.addMany(Harness_batchTitles(), 3)
	if (added != 2) || (app.totalCount() != 4) {
		hxrt.Throw(hxrt.StringFromLiteral("batch add drift"))
	}
	diag := app.diagnostics()
	if !hxrt.StringEqualStringPtr(diag, hxrt.StringFromLiteral("p1=1,completed=1")) {
		hxrt.Throw(hxrt.StringFromLiteral("missing diagnostics"))
	}
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("OK "), runtime.profileId())
}

func Harness_batchTitles() *hxrt.Array {
	out := hxrt.NewArray()
	out.Push(hxrt.StringFromLiteral("Ship generated-go sync"))
	out.Push(hxrt.StringFromLiteral("Add binary matrix"))
	return out
}

func Harness_run(runtime profile__TodoRuntime) *string {
	app := New_app__TodoApp(runtime)
	baselineView := Harness_runBaseline(app)
	baseline := app.baselineSignature()
	added := app.addMany(Harness_batchTitles(), 3)
	extras := hxrt.StringConcatAny(hxrt.StringFromLiteral("batch_add="), added)
	extras = hxrt.StringConcatStringPtr(extras, hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(",diag="), app.diagnostics()))
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("profile="), runtime.profileId()), hxrt.StringFromLiteral("\nbaseline=")), baseline), hxrt.StringFromLiteral("\nbaseline_view:\n")), baselineView), hxrt.StringFromLiteral("\nfinal_view:\n")), app.render()), hxrt.StringFromLiteral("\nextras=")), extras)
}

func Harness_runBaseline(app *app__TodoApp) *string {
	app.add(hxrt.StringFromLiteral("Write profile docs"), 2)
	app.add(hxrt.StringFromLiteral("Backfill regression snapshots"), 1)
	app.toggle(2)
	app.tag(1, hxrt.StringFromLiteral("docs"))
	app.tag(2, hxrt.StringFromLiteral("tests"))
	return app.render()
}
