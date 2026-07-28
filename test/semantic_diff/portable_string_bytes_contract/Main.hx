import haxe.io.Bytes;
import haxe.io.Encoding;

/**
	What
	Proves the source-visible string and byte behavior that any native Go-backed
	carrier must preserve.

	Why
	A raw Go string is not nullable and uses byte offsets, while a naked
	`[]byte` cannot preserve the `BytesData` integer view and its aliasing rules.

	How
	The same source runs through the Haxe interpreter and reflaxe.go. It checks
	Unicode-scalar operations, null string behavior, byte aliases, cached-view
	invalidation, overlapping copy, independent slices, and encoding round trips.
**/
class Main {
	static function byteValues(value:Bytes):String {
		return [for (index in 0...value.length) value.get(index)].join(",");
	}

	static function main() {
		final text = "A😀éZ";
		Sys.println("string.length=" + text.length);
		Sys.println("string.chars=" + text.charAt(0) + ":" + text.charAt(1) + ":" + text.charAt(2) + ":" + text.charAt(9));
		Sys.println("string.codes=" + text.charCodeAt(1) + ":" + text.charCodeAt(-1) + ":" + text.charCodeAt(9));
		Sys.println("string.slices=" + text.substring(1, 3) + ":" + text.substr(-2, 2) + ":" + text.split("").join("|"));

		final absent:String = null;
		final nullText = "null";
		Sys.println("string.null=" + (absent == null) + ":" + (absent == nullText) + ":" + (absent + "!"));
		final sameText = text;
		Sys.println("string.value=" + (sameText == text) + ":" + (("A" + "😀éZ") == text));

		final source = Bytes.ofString("abcd", Encoding.UTF8);
		final sourceAlias = source;
		sourceAlias.set(0, "A".code);
		Sys.println("bytes.object-alias=" + source.toString());

		// Materializing the native byte cache must not disconnect the public data.
		source.toString();
		final data = source.getData();
		data[1] = "B".code;
		Sys.println("bytes.data-to-object=" + source.toString());

		final dataAlias = Bytes.ofData(data);
		dataAlias.set(2, "C".code);
		Sys.println("bytes.object-to-data=" + source.toString() + ":" + data[2]);

		final copied = source.sub(0, source.length);
		copied.set(0, "x".code);
		Sys.println("bytes.sub-copy=" + source.toString() + ":" + copied.toString());

		final overlap = Bytes.ofString("12345");
		overlap.blit(1, overlap, 0, 4);
		Sys.println("bytes.overlap=" + overlap.toString());

		final emoji = Bytes.ofString("😀é", Encoding.UTF8);
		Sys.println("bytes.utf8=" + emoji.toHex() + ":" + emoji.getString(0, emoji.length, Encoding.UTF8));

		final raw = Bytes.ofString("hé", Encoding.RawNative);
		Sys.println("bytes.raw-default=" + byteValues(raw) + ":" + raw.getString(0, raw.length, Encoding.RawNative));

		final empty = Bytes.alloc(0);
		final missing:Null<Bytes> = null;
		Sys.println("bytes.empty-null=" + empty.length + ":" + (missing == null));
	}
}
