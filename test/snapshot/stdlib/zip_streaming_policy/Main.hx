import haxe.io.Bytes;
import haxe.io.BytesBuffer;
import haxe.zip.Compress;
import haxe.zip.FlushMode;
import haxe.zip.Uncompress;

class Main {
	static function throws(block:Void->Void):Bool {
		try {
			block();
		} catch (_:Dynamic) {
			return true;
		}
		return false;
	}

	static function streamRoundTrip(payload:Bytes):String {
		var compressor = new Compress(6);
		compressor.setFlushMode(FlushMode.FINISH);
		var compressedOutput = new BytesBuffer();
		var sourcePosition = 0;
		var compressCalls = 0;
		var compressDone = false;
		while (!compressDone) {
			var destination = Bytes.alloc(4);
			var result = compressor.execute(payload, sourcePosition, destination, 0);
			sourcePosition += result.read;
			compressedOutput.addBytes(destination, 0, result.write);
			compressDone = result.done;
			compressCalls++;
		}
		compressor.close();

		var compressed = compressedOutput.getBytes();
		var inflater = new Uncompress();
		inflater.setFlushMode(FlushMode.SYNC);
		var restoredOutput = new BytesBuffer();
		var compressedPosition = 0;
		var inflateCalls = 0;
		var inflateDone = false;
		while (!inflateDone) {
			var destination = Bytes.alloc(3);
			var result = inflater.execute(compressed, compressedPosition, destination, 0);
			compressedPosition += result.read;
			restoredOutput.addBytes(destination, 0, result.write);
			inflateDone = result.done;
			inflateCalls++;
		}
		inflater.close();

		return '${restoredOutput.getBytes().toString() == payload.toString()}:${compressCalls > 1}:${inflateCalls > 1}';
	}

	static function flushPolicy():String {
		var fullCompressor = new Compress(6);
		var full = throws(() -> fullCompressor.setFlushMode(FlushMode.FULL));
		fullCompressor.close();

		var blockCompressor = new Compress(6);
		var block = throws(() -> blockCompressor.setFlushMode(FlushMode.BLOCK));
		blockCompressor.close();

		var fullInflater = new Uncompress();
		var inflateFull = throws(() -> fullInflater.setFlushMode(FlushMode.FULL));
		fullInflater.close();

		var blockInflater = new Uncompress();
		var inflateBlock = throws(() -> blockInflater.setFlushMode(FlushMode.BLOCK));
		blockInflater.close();

		return '$full:$block:$inflateFull:$inflateBlock';
	}

	static function lifecycleAndBounds(payload:Bytes):String {
		var closedCompressor = new Compress(6);
		closedCompressor.close();
		closedCompressor.close();
		var compressClosed = throws(() -> closedCompressor.execute(payload, 0, Bytes.alloc(16), 0));

		var closedInflater = new Uncompress();
		closedInflater.close();
		closedInflater.close();
		var inflateClosed = throws(() -> closedInflater.execute(payload, 0, Bytes.alloc(16), 0));

		var compressor = new Compress(6);
		var compressBounds = throws(() -> compressor.execute(payload, -1, Bytes.alloc(16), 0));
		var zeroCompress = compressor.execute(payload, 0, Bytes.alloc(0), 0);
		compressor.close();

		var inflater = new Uncompress();
		var inflateBounds = throws(() -> inflater.execute(payload, 0, Bytes.alloc(16), 17));
		var zeroInflate = inflater.execute(payload, 0, Bytes.alloc(0), 0);
		inflater.close();

		var zeroCapacity = zeroCompress.read == 0 && zeroCompress.write == 0 && !zeroCompress.done && zeroInflate.read == 0 && zeroInflate.write == 0
			&& !zeroInflate.done;
		return '$compressClosed:$inflateClosed:$compressBounds:$inflateBounds:$zeroCapacity';
	}

	static function main():Void {
		var payload = Bytes.ofString("stream-safe-zip");
		Sys.println("stream=" + streamRoundTrip(payload));
		Sys.println("flush=" + flushPolicy());
		Sys.println("lifecycle=" + lifecycleAndBounds(payload));
	}
}
