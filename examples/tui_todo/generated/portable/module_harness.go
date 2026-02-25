package main

import "examples_tui_todo_portable/hxrt"

func Harness_assertContract(runtime profile__TodoRuntime) *string {
	app := New_app__TodoApp(runtime)
	Harness_runBaseline(app)
	baseline := app.baselineSignature()
	if !hxrt.StringEqualStringPtr(baseline, hxrt.StringFromLiteral("open=1,done=1,total=2")) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("baseline drift: "), baseline))
		var hx_throw_zero_1 *string
		return hx_throw_zero_1
	}
	if runtime.supportsBatchAdd() {
		added := app.addMany(Harness_batchTitles(), 3)
		if (added != 2) || (app.totalCount() != 4) {
			hxrt.Throw(hxrt.StringFromLiteral("batch add drift"))
			var hx_throw_zero_2 *string
			return hx_throw_zero_2
		}
	} else {
		if app.totalCount() != 2 {
			hxrt.Throw(hxrt.StringFromLiteral("portable total drift"))
			var hx_throw_zero_3 *string
			return hx_throw_zero_3
		}
	}
	if runtime.supportsDiagnostics() {
		diag := app.diagnostics()
		if !hxrt.StringEqualStringPtr(diag, hxrt.StringFromLiteral("p1=1,completed=1")) {
			hxrt.Throw(hxrt.StringFromLiteral("missing diagnostics"))
			var hx_throw_zero_4 *string
			return hx_throw_zero_4
		}
	}
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("OK "), runtime.profileId())
}

func Harness_batchTitles() *haxe__ds__List {
	out := New_haxe__ds__List()
	out.add(hxrt.StringFromLiteral("Ship generated-go sync"))
	out.add(hxrt.StringFromLiteral("Add binary matrix"))
	return out
}

func Harness_run(runtime profile__TodoRuntime) *string {
	app := New_app__TodoApp(runtime)
	baselineView := Harness_runBaseline(app)
	baseline := app.baselineSignature()
	extras := hxrt.StringFromLiteral("batch_add=0")
	if runtime.supportsBatchAdd() {
		added := app.addMany(Harness_batchTitles(), 3)
		extras = hxrt.StringConcatAny(hxrt.StringFromLiteral("batch_add="), added)
	}
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
