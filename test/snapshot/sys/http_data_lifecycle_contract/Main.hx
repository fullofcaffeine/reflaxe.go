import haxe.io.Bytes;
import haxe.io.BytesOutput;

private class DataOutput extends BytesOutput {
	public final events:Array<String> = [];
	public var closeCount(default, null):Int = 0;

	public function new() {
		super();
	}

	override public function prepare(size:Int):Void {
		events.push("prepare:" + size);
		super.prepare(size);
	}

	override public function writeBytes(bytes:Bytes, pos:Int, len:Int):Int {
		events.push("write:" + bytes.sub(pos, len).toString());
		return super.writeBytes(bytes, pos, len);
	}

	override public function close():Void {
		closeCount++;
		events.push("close");
		super.close();
	}
}

class Main {
	static function main() {
		var request = new haxe.Http("data:text/plain,hello%20world");
		var output = new DataOutput();
		request.onStatus = function(status) output.events.push("status:" + status);
		request.onError = function(message) output.events.push("error:" + message);
		request.customRequest(false, output);

		Sys.println("events=" + output.events.join(">"));
		Sys.println("body=" + output.getBytes().toString());
		Sys.println("closeCount=" + output.closeCount);
	}
}
