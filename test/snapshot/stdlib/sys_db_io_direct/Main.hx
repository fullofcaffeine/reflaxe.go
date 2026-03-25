package;

import StringBuf;
import sys.db.Connection;
import sys.db.Mysql;
import sys.db.ResultSet;
import sys.db.Sqlite;
import sys.io.File;
import sys.io.FileInput;
import sys.io.FileOutput;
import sys.io.FileSeek;

private class SnapResultSet implements ResultSet {
	public var length(get, never):Int;
	public var nfields(get, never):Int;

	var rows:Array<SnapRow>;
	var index:Int;

	public function new(rows:Array<SnapRow>) {
		this.rows = rows;
		this.index = 0;
	}

	function get_length():Int
		return rows.length - index;

	function get_nfields():Int
		return 1;

	public function hasNext():Bool
		return index < rows.length;

	public function next():Dynamic
		return rows[index++];

	public function results():haxe.ds.List<Dynamic> {
		var out = new haxe.ds.List<Dynamic>();
		while (hasNext())
			out.add(next());
		return out;
	}

	public function getResult(n:Int):String
		return Std.string(rows[index - 1].value);

	public function getIntResult(n:Int):Int
		return rows[index - 1].value;

	public function getFloatResult(n:Int):Float
		return rows[index - 1].value + 0.0;

	public function getFieldsNames():Null<Array<String>>
		return ["value"];
}

private class SnapRow {
	public var value:Int;

	public function new(value:Int) {
		this.value = value;
	}
}

private class SnapConnection implements Connection {
	public function new() {}

	public function request(s:String):ResultSet
		return new SnapResultSet([new SnapRow(3)]);

	public function close():Void {}

	public function escape(s:String):String
		return s;

	public function quote(s:String):String
		return '"' + s + '"';

	public function addValue(s:StringBuf, v:Dynamic):Void
		s.add(Std.string(v));

	public function lastInsertId():Int
		return 2;

	public function dbName():String
		return "snap";

	public function startTransaction():Void {}

	public function commit():Void {}

	public function rollback():Void {}
}

class Main {
	static function main() {
		var output:FileOutput = File.write('snapshot.txt', true);
		output.writeString('ok');
		output.seek(0, SeekBegin);
		output.writeByte('O'.code);
		output.close();
		var input:FileInput = File.read('snapshot.txt', true);
		trace(input.tell());
		trace(input.readByte());
		trace(input.eof());
		input.close();
		var conn:Connection = new SnapConnection();
		var rs:ResultSet = conn.request('select');
		trace(conn.dbName());
		trace(rs.getIntResult(0));
		try
			Mysql.connect({host: 'localhost', user: 'u', pass: 'p'})
		catch (e)
			trace(Std.string(e));
		try
			Sqlite.open('db.sqlite')
		catch (e)
			trace(Std.string(e));
		sys.FileSystem.deleteFile('snapshot.txt');
	}
}
