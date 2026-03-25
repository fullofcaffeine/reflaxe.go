package;

import haxe.ds.List;
import StringBuf;
import sys.db.Connection;
import sys.db.Mysql;
import sys.db.ResultSet;
import sys.db.Sqlite;
import sys.io.File;
import sys.io.FileInput;
import sys.io.FileOutput;
import sys.io.FileSeek;

private class FakeResultSet implements ResultSet {
	public var length(get, never):Int;
	public var nfields(get, never):Int;

	var rows:Array<FakeRow>;
	var index:Int;

	public function new(rows:Array<FakeRow>) {
		this.rows = rows;
		this.index = 0;
	}

	function get_length():Int
		return rows.length - index;

	function get_nfields():Int
		return 2;

	public function hasNext():Bool
		return index < rows.length;

	public function next():Dynamic
		return rows[index++];

	public function results():List<Dynamic> {
		var out = new List<Dynamic>();
		while (hasNext())
			out.add(next());
		return out;
	}

	public function getResult(n:Int):String
		return Std.string(Reflect.field(rows[index - 1], n == 0 ? "id" : "name"));

	public function getIntResult(n:Int):Int
		return n == 0 ? rows[index - 1].id : rows[index - 1].name.length;

	public function getFloatResult(n:Int):Float
		return getIntResult(n) + 0.0;

	public function getFieldsNames():Null<Array<String>>
		return ["id", "name"];
}

private class FakeRow {
	public var id:Int;
	public var name:String;

	public function new(id:Int, name:String) {
		this.id = id;
		this.name = name;
	}
}

private class FakeConnection implements Connection {
	public function new() {}

	public function request(s:String):ResultSet
		return new FakeResultSet([new FakeRow(7, "go"), new FakeRow(9, "hx")]);

	public function close():Void {}

	public function escape(s:String):String
		return "[" + s + "]";

	public function quote(s:String):String
		return "'" + escape(s) + "'";

	public function addValue(s:StringBuf, v:Dynamic):Void
		s.add(Std.string(v));

	public function lastInsertId():Int
		return 41;

	public function dbName():String
		return "fake";

	public function startTransaction():Void {}

	public function commit():Void {}

	public function rollback():Void {}
}

class Main {
	static function main() {
		var path = "sys_db_io_contract.txt";
		var output:FileOutput = File.write(path, true);
		output.writeString("abc");
		trace('out.tell.1=' + output.tell());
		output.seek(1, SeekBegin);
		output.writeByte('Z'.code);
		trace('out.tell.2=' + output.tell());
		output.close();

		var input:FileInput = File.read(path, true);
		trace('in.tell.1=' + input.tell());
		trace('in.byte.1=' + input.readByte());
		trace('in.tell.2=' + input.tell());
		input.seek(1, SeekCur);
		trace('in.byte.2=' + input.readByte());
		trace('in.eof=' + input.eof());
		input.close();
		trace('content=' + File.getContent(path));
		sys.FileSystem.deleteFile(path);

		var seeks = [SeekBegin, SeekCur, SeekEnd];
		trace('seek.tags=' + [for (s in seeks) Type.enumConstructor(s)].join(','));

		var conn:Connection = new FakeConnection();
		var rs:ResultSet = conn.request('select 1');
		trace('db.name=' + conn.dbName());
		trace('db.quote=' + conn.quote("g'o"));
		trace('db.last=' + conn.lastInsertId());
		trace('db.fields=' + rs.getFieldsNames().join(','));
		trace('db.length=' + rs.length);
		trace('db.next=' + rs.next().name);
		trace('db.result=' + rs.getIntResult(0));
		trace('db.remaining=' + [for (row in rs.results()) row.name].join(','));

		try {
			Mysql.connect({host: 'localhost', user: 'u', pass: 'p'});
			trace('mysql=no-throw');
		} catch (e) {
			trace('mysql=' + Std.string(e));
		}

		try {
			Sqlite.open('db.sqlite');
			trace('sqlite=no-throw');
		} catch (e) {
			trace('sqlite=' + Std.string(e));
		}
	}
}
