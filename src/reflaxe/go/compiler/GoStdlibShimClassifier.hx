package reflaxe.go.compiler;

#if macro
import haxe.macro.Type;

class GoStdlibShimClassifier {
	public static function needsIoHelperSurface(classType:ClassType, fieldName:String, isIoInputHelperMethodName:String->Bool,
			isIoOutputHelperMethodName:String->Bool):Bool {
		if (classType.pack.join(".") != "haxe.io") {
			return false;
		}
		return switch (classType.name) {
			case "Input", "BytesInput", "BufferInput", "StringInput":
				isIoInputHelperMethodName(fieldName);
			case "Output", "BytesOutput":
				isIoOutputHelperMethodName(fieldName);
			case _:
				false;
		};
	}

	public static function requiredGroupsForClass(classType:ClassType):Array<String> {
		var pack = classType.pack.join(".");
		if (pack == "haxe.io") {
			return switch (classType.name) {
				case "BufferInput", "Bytes", "BytesBuffer", "Input", "Output", "Encoding", "BytesInput", "BytesOutput", "Eof", "StringInput":
					["io"];
				case "Path":
					["stdlib_symbols"];
				case _:
					[];
			};
		}

		if (pack == "haxe.ds") {
			return switch (classType.name) {
				case "IntMap", "StringMap", "ObjectMap", "EnumValueMap", "List":
					["ds"];
				case "BalancedTree":
					["stdlib_symbols"];
				case _:
					[];
			};
		}

		if (pack == "sys" && classType.name == "Http") {
			return ["http"];
		}

		if (pack == "sys" && classType.name == "FileSystem") {
			return ["filesystem"];
		}

		if ((pack == "" && classType.name == "Sys") || (pack == "sys.io" && (classType.name == "File" || classType.name == "Process"))) {
			return ["sys"];
		}

		if (pack == "sys.net" && (classType.name == "Host" || classType.name == "Socket" || classType.name == "UdpSocket")) {
			return ["net_socket"];
		}

		if ((pack == "haxe.atomic"
			&& (classType.name == "AtomicInt" || classType.name == "AtomicBool" || classType.name == "AtomicObject"))
			|| (pack == "haxe.atomic._AtomicInt" && classType.name == "AtomicInt_Impl_")
			|| (pack == "haxe.atomic._AtomicBool" && classType.name == "AtomicBool_Impl_")
			|| (pack == "haxe.atomic._AtomicObject" && classType.name == "AtomicObject_Impl_")) {
			return ["atomic"];
		}

		if ((pack == "" && classType.name == "EReg")
			|| (pack == "haxe" && (classType.name == "Serializer" || classType.name == "Unserializer"))) {
			return ["regex_serializer"];
		}

		if (((pack == ""
			&& (classType.name == "Std" || classType.name == "Date" || classType.name == "Math" || classType.name == "Type" || classType.name == "Reflect"
				|| classType.name == "Xml" || classType.name == "UnicodeString"))
			|| (pack == "_UnicodeString" && classType.name == "UnicodeString_Impl_"))
			|| (pack == "haxe.crypto"
				&& (classType.name == "Base64" || classType.name == "Md5" || classType.name == "Sha1" || classType.name == "Sha224"
					|| classType.name == "Sha256"))
			|| (pack == "haxe.xml" && (classType.name == "Parser" || classType.name == "Printer"))
			|| (pack == "haxe.zip" && (classType.name == "Compress" || classType.name == "Uncompress"))) {
			return ["stdlib_symbols"];
		}

		return [];
	}

	public static function requiredGroupsForEnum(enumType:EnumType):Array<String> {
		var pack = enumType.pack.join(".");
		if (pack == "haxe.io" && enumType.name == "Error") {
			return ["io"];
		}
		if (pack == "haxe.ds" && enumType.name == "Option") {
			return ["stdlib_symbols"];
		}
		return [];
	}
}
#end
