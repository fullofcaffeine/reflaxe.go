import haxe.io.Bytes;
import haxe.zip.Compress;
import haxe.zip.Entry;
import haxe.zip.FlushMode;
import haxe.zip.Tools;
import haxe.zip.Uncompress;

class Main {
	static function main() {
		var payload = Bytes.ofString("repeat-repeat-repeat-repeat-repeat");
		var levels = [-1, 0, 1, 6, 9];
		var levelResults = new Array<String>();
		for (level in levels) {
			var compressed = Compress.run(payload, level);
			levelResults.push(Uncompress.run(compressed).toString());
		}
		Sys.println("levels=" + levelResults.join("|"));

		var compressed = Compress.run(payload, 6);
		Sys.println("buffers=" + Uncompress.run(compressed, 1).toString() + ":" + Uncompress.run(compressed, 7).toString());

		var binary = Bytes.alloc(4);
		binary.set(0, 0);
		binary.set(1, 127);
		binary.set(2, 128);
		binary.set(3, 255);
		Sys.println("binary=" + Uncompress.run(Compress.run(binary, 9), 2).toHex());
		Sys.println("empty=" + Uncompress.run(Compress.run(Bytes.alloc(0), 6), 3).length);

		var compressor = new Compress(6);
		compressor.setFlushMode(FlushMode.FINISH);
		var compressedBuffer = Bytes.alloc(256);
		var compressedResult = compressor.execute(payload, 0, compressedBuffer, 0);
		compressor.close();
		// Haxe 4.3.7's eval extern misspells only this result field as `wriet`.
		// Keep that reference-harness typo out of the typed Go-side contract.
		#if eval
		var evalCompressedResult:{done:Bool, read:Int, write:Int} = cast compressedResult;
		var compressedWrite = evalCompressedResult.write;
		#else
		var compressedWrite = compressedResult.write;
		#end
		var instanceCompressed = compressedBuffer.sub(0, compressedWrite);

		var uncompressor = new Uncompress();
		uncompressor.setFlushMode(FlushMode.SYNC);
		var restoredBuffer = Bytes.alloc(256);
		var restoredResult = uncompressor.execute(instanceCompressed, 0, restoredBuffer, 0);
		uncompressor.close();
		Sys.println("instance=" + compressedResult.done + ":" + (compressedResult.read == payload.length) + ":" + restoredResult.done + ":"
			+ (restoredResult.read == instanceCompressed.length) + ":" + restoredBuffer.sub(0, restoredResult.write).toString());

		var entry:Entry = {
			fileName: "payload.txt",
			fileSize: payload.length,
			fileTime: Date.fromTime(0),
			compressed: false,
			dataSize: payload.length,
			data: payload,
			crc32: null
		};
		Tools.compress(entry, 6);
		var rawSize = entry.dataSize;
		Tools.uncompress(entry);
		Sys.println("tools=" + (rawSize > 0) + ":" + entry.compressed + ":" + entry.data.toString());
	}
}
