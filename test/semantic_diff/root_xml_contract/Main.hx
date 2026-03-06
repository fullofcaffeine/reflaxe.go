class Main {
	static function main() {
		var doc = Xml.createDocument();
		var root = Xml.createElement("root");
		root.set("id", "r1");
		root.set("kind", "demo");

		var before = Xml.createComment("before");
		var itemA = Xml.createElement("item");
		itemA.set("n", "1");
		itemA.addChild(Xml.createPCData("x"));
		var itemB = Xml.createElement("item");
		itemB.set("n", "2");
		itemB.addChild(Xml.createPCData("y"));

		root.addChild(before);
		root.addChild(itemA);
		root.addChild(itemB);
		doc.addChild(root);

		Sys.println("doc.type=" + doc.nodeType);
		Sys.println("root.name=" + root.nodeName);
		Sys.println("root.id=" + root.get("id"));
		Sys.println("root.exists=" + root.exists("kind"));
		Sys.println("root.firstElement=" + root.firstElement().nodeName);
		Sys.println("itemA.value=" + itemA.firstChild().nodeValue);

		var attrIter = root.attributes();
		var attrFirst = attrIter.next();
		var attrSecond = attrIter.next();
		if (Reflect.compare(attrFirst, attrSecond) > 0) {
			var tmp = attrFirst;
			attrFirst = attrSecond;
			attrSecond = tmp;
		}
		Sys.println("attrs=" + attrFirst + "=" + root.get(attrFirst) + "," + attrSecond + "=" + root.get(attrSecond));

		var namedIter = root.elementsNamed("item");
		var namedFirst = namedIter.next();
		var namedSecond = namedIter.next();
		Sys.println("named=" + namedFirst.get("n") + ":" + namedFirst.firstChild().nodeValue + "," + namedSecond.get("n") + ":"
			+ namedSecond.firstChild().nodeValue);

		var moved = Xml.createElement("moved");
		moved.addChild(Xml.createPCData("z"));
		root.insertChild(moved, 1);
		Sys.println("after.insert=" + root.firstElement().nodeName + "," + root.elementsNamed("moved").hasNext());
		Sys.println("remove.comment=" + root.removeChild(before));

		var parsed = Xml.parse("<outer><inner a=\"1\">v</inner><inner a=\"2\"/></outer>");
		var parsedRoot = parsed.firstElement();
		Sys.println("parsed.root=" + parsedRoot.nodeName);
		Sys.println("parsed.first=" + parsedRoot.firstElement().get("a") + ":" + parsedRoot.firstElement().firstChild().nodeValue);

		var parsedIter = parsedRoot.elementsNamed("inner");
		var parsedFirst = parsedIter.next();
		var parsedSecond = parsedIter.next();
		Sys.println("parsed.named=" + parsedFirst.get("a") + "," + parsedSecond.get("a"));
		Sys.println("parsed.str=" + parsed.toString());

		var parsedCData = Xml.parse("<outer><![CDATA[x]]></outer>");
		var parsedCDataChild = parsedCData.firstElement().firstChild();
		Sys.println("parsed.cdata.type=" + parsedCDataChild.nodeType);
		Sys.println("parsed.cdata.value=" + parsedCDataChild.nodeValue);
		Sys.println("parsed.cdata.str=" + parsedCData.toString());
	}
}
