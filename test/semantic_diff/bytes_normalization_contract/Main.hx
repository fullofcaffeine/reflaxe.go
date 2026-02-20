class Main {
	static function main() {
		var bytes = haxe.io.Bytes.alloc(3);
		bytes.set(0, 300);
		bytes.set(1, -1);
		bytes.set(2, 513);
		Sys.println("vals=" + bytes.get(0) + "," + bytes.get(1) + "," + bytes.get(2));

		var buffer = new haxe.io.BytesBuffer();
		buffer.addByte(260);
		buffer.addByte(-2);
		var out = buffer.getBytes();
		Sys.println("buf=" + out.get(0) + "," + out.get(1));

		var ascii = haxe.io.Bytes.alloc(3);
		ascii.set(0, 65);
		ascii.set(1, 66);
		ascii.set(2, 67);
		Sys.println("str=" + ascii.toString());
	}
}
