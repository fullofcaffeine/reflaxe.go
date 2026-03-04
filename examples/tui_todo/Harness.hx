import app.TodoApp;
import profile.TodoRuntime;

class Harness {
	static function batchTitles():Array<String> {
		var out = [];
		out.push("Ship generated-go sync");
		out.push("Add binary matrix");
		return out;
	}

	static function runBaseline(app:TodoApp):String {
		app.add("Write profile docs", 2);
		app.add("Backfill regression snapshots", 1);
		app.toggle(2);
		app.tag(1, "docs");
		app.tag(2, "tests");
		return app.render();
	}

	public static function run(runtime:TodoRuntime):String {
		var app = new TodoApp(runtime);
		var baselineView = runBaseline(app);
		var baseline = app.baselineSignature();

		var added = app.addMany(batchTitles(), 3);
		var extras = "batch_add=" + added;
		extras += ",diag=" + app.diagnostics();

		return "profile=" + runtime.profileId() + "\nbaseline=" + baseline + "\nbaseline_view:\n" + baselineView + "\nfinal_view:\n" + app.render()
			+ "\nextras=" + extras;
	}

	public static function assertContract(runtime:TodoRuntime):String {
		var app = new TodoApp(runtime);
		runBaseline(app);

		var baseline = app.baselineSignature();
		if (baseline != "open=1,done=1,total=2") {
			throw "baseline drift: " + baseline;
		}

		var added = app.addMany(batchTitles(), 3);
		if (added != 2 || app.totalCount() != 4) {
			throw "batch add drift";
		}

		var diag = app.diagnostics();
		if (diag != "p1=1,completed=1") {
			throw "missing diagnostics";
		}

		return "OK " + runtime.profileId();
	}
}
