import haxe.io.Bytes;
import haxe.io.BytesBuffer;
import haxe.zip.Compress;
import haxe.zip.FlushMode;
import haxe.zip.Uncompress;

private typedef ZipProgress = {
	final done:Bool;
	final read:Int;
	final write:Int;
}

class Main {
	// Haxe 4.3.7's Eval extern misspells `write`. Keep that upstream typing
	// defect at this reference-only adapter instead of weakening production APIs.
	static inline function progress(value:Dynamic):ZipProgress {
		return cast value;
	}

	static function appendOutput(output:BytesBuffer, chunk:Bytes, length:Int):Void {
		if (length > 0)
			output.addBytes(chunk, 0, length);
	}

	static function compressProgressively(parts:Array<Bytes>, outputSize:Int):Bytes {
		var codec = new Compress(6);
		var output = new BytesBuffer();
		var calls = 0;
		var totalRead = 0;
		var totalWrite = 0;
		var done = false;

		for (index in 0...parts.length) {
			var part = parts[index];
			codec.setFlushMode(index == parts.length - 1 ? FINISH : (index == 0 ? NO : SYNC));
			var position = 0;
			do {
				var destination = Bytes.alloc(outputSize);
				var result = progress(codec.execute(part, position, destination, 0));
				calls++;
				position += result.read;
				totalRead += result.read;
				totalWrite += result.write;
				appendOutput(output, destination, result.write);
				done = result.done;
				if (!done && result.read == 0 && result.write == 0)
					throw "compression made no progress";
			} while (position < part.length || (index == parts.length - 1 && !done));
		}

		codec.close();
		var compressed = output.getBytes();
		Sys.println("compress=" + (calls > 1) + ":" + (totalRead == 23) + ":" + (totalWrite == compressed.length) + ":" + done);
		return compressed;
	}

	static function uncompressProgressively(compressed:Bytes, inputSize:Int, outputSize:Int, windowBits:Int):Bytes {
		var codec = new Uncompress(windowBits);
		codec.setFlushMode(SYNC);
		var output = new BytesBuffer();
		var calls = 0;
		var totalRead = 0;
		var totalWrite = 0;
		var done = false;
		var compressedPosition = 0;

		while (!done) {
			var remaining = compressed.length - compressedPosition;
			var partLength = remaining < inputSize ? remaining : inputSize;
			var part = compressed.sub(compressedPosition, partLength);
			var partPosition = 0;
			do {
				var destination = Bytes.alloc(outputSize);
				var result = progress(codec.execute(part, partPosition, destination, 0));
				calls++;
				partPosition += result.read;
				totalRead += result.read;
				totalWrite += result.write;
				appendOutput(output, destination, result.write);
				done = result.done;
				if (!done && result.read == 0 && result.write == 0)
					throw "decompression made no progress";
			} while (partPosition < part.length || (compressedPosition + partPosition == compressed.length && !done));
			compressedPosition += partPosition;
			if (!done && compressedPosition == compressed.length)
				throw "decompression exhausted its input";
		}

		codec.close();
		var restored = output.getBytes();
		Sys.println("uncompress="
			+ (calls > 1)
			+ ":"
			+ (totalRead == compressed.length)
			+ ":"
			+ (totalWrite == restored.length)
			+ ":"
			+ done);
		return restored;
	}

	static function offsetRoundTrip(payload:Bytes):Bool {
		var source = Bytes.alloc(payload.length + 1);
		source.set(0, 0x6a);
		source.blit(1, payload, 0, payload.length);
		var compressedDestination = Bytes.alloc(128);
		compressedDestination.set(0, 0x4a);
		compressedDestination.set(1, 0x5a);
		var compressor = new Compress(6);
		compressor.setFlushMode(FINISH);
		var compressedStep = progress(compressor.execute(source, 1, compressedDestination, 2));
		compressor.close();
		var compressed = compressedDestination.sub(2, compressedStep.write);

		var compressedSource = Bytes.alloc(compressed.length + 1);
		compressedSource.set(0, 0x7a);
		compressedSource.blit(1, compressed, 0, compressed.length);
		var restoredDestination = Bytes.alloc(payload.length + 2);
		restoredDestination.set(0, 0x2a);
		restoredDestination.set(1, 0x3a);
		var inflater = new Uncompress();
		inflater.setFlushMode(SYNC);
		var restoredStep = progress(inflater.execute(compressedSource, 1, restoredDestination, 2));
		inflater.close();

		return compressedStep.done
			&& compressedStep.read == payload.length
			&& compressedDestination.get(0) == 0x4a
			&& compressedDestination.get(1) == 0x5a
			&& restoredStep.done
			&& restoredStep.read == compressed.length
			&& restoredDestination.get(0) == 0x2a
			&& restoredDestination.get(1) == 0x3a
			&& restoredDestination.sub(2, restoredStep.write).toString() == payload.toString();
	}

	static function uncompressWithMode(compressed:Bytes, payload:Bytes, mode:FlushMode):Bool {
		var inflater = new Uncompress();
		inflater.setFlushMode(mode);
		var destination = Bytes.alloc(payload.length + 8);
		var step = progress(inflater.execute(compressed, 0, destination, 0));
		inflater.close();
		return step.done && step.read == compressed.length && destination.sub(0, step.write).toString() == payload.toString();
	}

	static function main():Void {
		var parts = [Bytes.ofString("stream-"), Bytes.ofString("safe-"), Bytes.ofString("zip-zip-zip")];
		var payload = Bytes.ofString("stream-safe-zip-zip-zip");
		var compressed = compressProgressively(parts, 5);
		var restored = uncompressProgressively(compressed, 4, 3, 0);
		Sys.println("roundtrip=" + (restored.toString() == payload.toString()));

		var wrapped = Compress.run(payload, 6);
		var raw = wrapped.sub(2, wrapped.length - 6);
		var rawRestored = uncompressProgressively(raw, 3, 2, -15);
		Sys.println("raw=" + (rawRestored.toString() == payload.toString()));
		Sys.println("offsets=" + offsetRoundTrip(payload));
		Sys.println("inflate-modes=" + uncompressWithMode(wrapped, payload, NO) + ":" + uncompressWithMode(wrapped, payload, FINISH));
	}
}
