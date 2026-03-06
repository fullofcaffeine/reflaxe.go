class Main {
	static function main() {
		Sys.println("trim=" + StringTools.trim("  hi  "));
		Sys.println("starts=" + StringTools.startsWith("hello", "he"));
		Sys.println("replace=" + StringTools.replace("a-b-c", "-", ":"));
		Sys.println("contains=" + StringTools.contains("banana", "nan"));
		Sys.println("ends=" + StringTools.endsWith("banana", "na"));
	}
}
