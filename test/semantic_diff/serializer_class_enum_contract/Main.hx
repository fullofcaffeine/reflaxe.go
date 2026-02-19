class ProbeClass {
	public var x:Int;
	public var s:String;

	public function new(x:Int, s:String) {
		this.x = x;
		this.s = s;
	}
}

enum ProbeEnum {
	NoArgs;
	One(v:Int);
	Two(a:Int, b:String);
}

class Main {
	static function safe(label:String, fn:Void->String):Void {
		try {
			Sys.println(label + "=ok:" + fn());
		} catch (e:Dynamic) {
			Sys.println(label + "=err");
		}
	}

	static function main() {
		var classValue = new ProbeClass(3, "k");
		safe("ser.class.value", function() return haxe.Serializer.run(classValue));
		safe("ser.class.token", function() return haxe.Serializer.run(haxe.Unserializer.run("cy10:ProbeClassy1:xi3y1:sy1:kg")));

		var enumNoArgs = ProbeEnum.NoArgs;
		var enumWithArgs = ProbeEnum.Two(4, "hi");
		safe("ser.enum.value.noArgs", function() return haxe.Serializer.run(enumNoArgs));
		safe("ser.enum.value.withArgs", function() return haxe.Serializer.run(enumWithArgs));
		safe("ser.enum.token.noArgs", function() return haxe.Serializer.run(haxe.Unserializer.run("wy9:ProbeEnumy6:NoArgs:0")));
		safe("ser.enum.token.withArgs", function() return haxe.Serializer.run(haxe.Unserializer.run("wy9:ProbeEnumy3:Two:2i4y2:hi")));
	}
}
