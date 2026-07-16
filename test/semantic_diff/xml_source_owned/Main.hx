import haxe.xml.Parser;
import haxe.xml.Parser.XmlParserException;
import haxe.xml.Printer;

class Main {
	static function count(iterator:Iterator<Xml>):Int {
		var total = 0;
		while (iterator.hasNext()) {
			iterator.next();
			total++;
		}
		return total;
	}

	static function childNames(parent:Xml):String {
		var result = "";
		for (child in parent.elements()) {
			if (result.length != 0)
				result += ",";
			result += child.nodeName;
		}
		return result;
	}

	static function parseResult(source:String, strict:Bool):String {
		try {
			return "ok:" + Printer.print(Parser.parse(source, strict));
		} catch (error:XmlParserException) {
			return "xml-error:" + error.message + ":" + error.lineNumber + ":" + error.positionAtLine + ":" + error.position;
		} catch (error:Dynamic) {
			return "other-error:" + Std.string(error);
		}
		return "unreachable";
	}

	static function badNodeName():String {
		try {
			return Xml.createPCData("value").nodeName;
		} catch (error:Dynamic) {
			return Std.string(error);
		}
		return "unreachable";
	}

	static function main() {
		var document = Xml.createDocument();
		var left = Xml.createElement("left");
		var right = Xml.createElement("right");
		var first = Xml.createElement("first");
		var moved = Xml.createElement("moved");
		moved.addChild(Xml.createPCData("<&>"));

		left.addChild(moved);
		right.addChild(first);
		right.addChild(moved);
		document.addChild(left);
		document.addChild(right);

		right.set("z", "last");
		right.set("a", "<&\"");
		right.set("z", "updated");
		var attrIterator = right.attributes();
		var attrFirst = attrIterator.next();
		var attrSecond = attrIterator.next();
		if (Reflect.compare(attrFirst, attrSecond) > 0) {
			var swap = attrFirst;
			attrFirst = attrSecond;
			attrSecond = swap;
		}

		Sys.println("types=" + document.nodeType + "," + right.nodeType + "," + moved.firstChild().nodeType);
		Sys.println("parents=" + (moved.parent == right) + "," + count(left.iterator()) + "," + count(right.iterator()));
		Sys.println("attrs=" + attrFirst + "," + attrSecond + ":" + right.get("a") + ":" + right.get("z"));
		Sys.println("named=" + count(right.elementsNamed("moved")) + ":" + right.firstElement().nodeName);
		Sys.println("compact=" + Printer.print(document));
		Sys.println("pretty=" + StringTools.replace(Printer.print(document, true), "\n", "|"));

		var inserted = Xml.createComment(" before\n\tafter ");
		right.insertChild(inserted, 1);
		Sys.println("insert=" + count(right.iterator()) + ":" + inserted.parent.nodeName + ":" + right.removeChild(inserted) + ":" + (inserted.parent == null));

		var order = Xml.createElement("order");
		var a = Xml.createElement("a");
		var b = Xml.createElement("b");
		var c = Xml.createElement("c");
		var d = Xml.createElement("d");
		var e = Xml.createElement("e");
		order.addChild(a);
		order.addChild(b);
		order.insertChild(c, -1);
		order.insertChild(d, -99);
		order.insertChild(e, 99);
		order.insertChild(a, 4);
		var outsider = Xml.createElement("outsider");
		var foreign = Xml.createElement("foreign");
		outsider.addChild(foreign);
		Sys.println("insert-order=" + childNames(order));
		Sys.println("remove-missing=" + order.removeChild(foreign) + ":" + (foreign.parent == outsider));
		Sys.println("empty-child=" + (Xml.createElement("empty").firstChild() == null));
		Sys.println("bad-name=" + badNodeName());

		var rich = '<?xml version="1.0"?><!DOCTYPE root [<!ELEMENT root ANY>]><root a="&lt;&#65;&#x42;">x&amp;<![CDATA[<raw>]]><!-- c\n\t d --></root>';
		var parsed = Parser.parse(rich, true);
		var root = parsed.firstElement();
		Sys.println("rich.types=" + parsed.firstChild().nodeType + "," + root.firstChild().nodeType + "," + root.firstElement());
		Sys.println("rich.attr=" + root.get("a"));
		Sys.println("rich.print=" + Printer.print(parsed));

		Sys.println("unknown.loose=" + parseResult('<root a="&unknown;"/>', false));
		Sys.println("unknown.strict=" + parseResult('<root a="&unknown;"/>', true));
		Sys.println("attribute.loose=" + parseResult('<root a="x>y"/>', false));
		Sys.println("attribute.strict=" + parseResult('<root a="x>y"/>', true));
		Sys.println("duplicate=" + parseResult('<root a="1" a="2"/>', false));
		Sys.println("mismatch=" + parseResult("<root>\n<child></root>", true));
		Sys.println("unclosed=" + parseResult("<root>", true));
	}
}
