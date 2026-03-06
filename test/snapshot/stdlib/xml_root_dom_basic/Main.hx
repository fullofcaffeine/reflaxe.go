class Main {
	static function main() {
		var doc = Xml.createDocument();
		var root = Xml.createElement("root");
		root.set("id", "r1");
		root.addChild(Xml.createPCData("hello"));
		doc.addChild(root);

		var parsed = Xml.parse("<outer><inner a=\"1\">v</inner></outer>");
		Sys.println(doc.toString());
		Sys.println(parsed.firstElement().firstElement().get("a"));
		Sys.println(parsed.toString());
	}
}
