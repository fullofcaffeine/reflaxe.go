package go;

/**
	What: An opaque, lossless view of one native Go error interface.
	Why: go.Error preserves legacy message snapshots, while native classification
	needs the original concrete value, wrapped causes and custom matching methods.
	How: Imported extern boundaries retain the original interface in a private
	runtime carrier and export that same root. A nil interface becomes null; an
	interface containing a typed nil remains non-null. Carrier equality does not
	mean native identity. Use typed native operations to compare or classify.
**/
@:go.import("hxrt")
@:go.name("NativeErrorValue")
extern class NativeError {
	/** Calls the original Error method now; transport itself never formats. */
	@:go.name("Text")
	public function text():String;
}
