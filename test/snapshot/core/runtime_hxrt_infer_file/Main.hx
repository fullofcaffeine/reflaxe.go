import haxe.io.Bytes;
import sys.io.File;
import sys.io.FileSeek;

class Main {
	static function main() {
		var path = "runtime_hxrt_infer_file.bin";
		File.saveBytes(path, Bytes.ofHex("00ff80"));
		var output = File.update(path, true);
		output.seek(1, SeekBegin);
		output.writeByte(7);
		output.close();
		var input = File.read(path, true);
		input.readByte();
		input.close();
	}
}
