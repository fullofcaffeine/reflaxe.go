package main

import "snapshot/hxrt"

func main() {
	output := sys__io__File_write(hxrt.StringFromLiteral("snapshot.txt"), true)
	output.__hx_this.writeString(hxrt.StringFromLiteral("ok"), nil)
	output.__hx_this.seek(0, sys__io__FileSeek_SeekBegin)
	output.__hx_this.writeByte(79)
	output.__hx_this.close()
	input := sys__io__File_read(hxrt.StringFromLiteral("snapshot.txt"), true)
	input.__hx_this.close()
	var conn sys__db__Connection = New__Main__SnapConnection()
	var rs sys__db__ResultSet = conn.request(hxrt.StringFromLiteral("select"))
	_ = rs
	hxrt.TryCatch(func() {
		sys__db__Mysql_connect(func() hxrt.Anonymous {
			hx_obj_4 := map[string]any{}
			hx_obj_4["host"] = hxrt.StringFromLiteral("localhost")
			hx_obj_4["user"] = hxrt.StringFromLiteral("u")
			hx_obj_4["pass"] = hxrt.StringFromLiteral("p")
			return hxrt.AnonymousFromMap(hx_obj_4)
		}())
	}, func(hx_caught_1 any) {
		hx_exception_2 := haxe__Exception_caught(hx_caught_1)
		hx_payload_3 := haxe__Exception_catchValue(hx_caught_1)
		_ = hx_payload_3
		var e *haxe__Exception = hx_recover_class_0(hx_exception_2)
		_ = e
	})
	hxrt.TryCatch(func() {
		sys__db__Sqlite_open(hxrt.StringFromLiteral("db.sqlite"))
	}, func(hx_caught_5 any) {
		hx_exception_6 := haxe__Exception_caught(hx_caught_5)
		hx_payload_7 := haxe__Exception_catchValue(hx_caught_5)
		_ = hx_payload_7
		var e_1 *haxe__Exception = hx_recover_class_0(hx_exception_6)
		_ = e_1
	})
	sys__FileSystem_deleteFile(hxrt.StringFromLiteral("snapshot.txt"))
}

type I__Main__SnapConnection interface {
	request(s *string) sys__db__ResultSet
	close()
	escape(s *string) *string
	quote(s *string) *string
	addValue(s *StringBuf, v any)
	lastInsertId() int
	dbName() *string
	startTransaction()
	commit()
	rollback()
}

type _Main__SnapConnection struct {
	__hx_this I__Main__SnapConnection
}

func (self *_Main__SnapConnection) HxFunctionReceiver() any {
	return self.__hx_this
}

func New__Main__SnapConnection() *_Main__SnapConnection {
	return hx_init__Main__SnapConnection(hx_alloc__Main__SnapConnection(nil))
}

func hx_alloc__Main__SnapConnection(owner any) *_Main__SnapConnection {
	self := &_Main__SnapConnection{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I__Main__SnapConnection)
	}
	return self
}

func hx_init__Main__SnapConnection(self *_Main__SnapConnection) *_Main__SnapConnection {
	return self
}

func (self *_Main__SnapConnection) request(s *string) sys__db__ResultSet {
	return New__Main__SnapResultSet(hxrt.NewArray(New__Main__SnapRow(3)))
}

func (self *_Main__SnapConnection) close() {
}

func (self *_Main__SnapConnection) escape(s *string) *string {
	return s
}

func (self *_Main__SnapConnection) quote(s *string) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("\""), s), hxrt.StringFromLiteral("\""))
}

func (self *_Main__SnapConnection) addValue(s *StringBuf, v any) {
	x := hxrt.StdString(v)
	s.b = hxrt.StringConcatStringPtr(s.b, hxrt.StdString(x))
}

func (self *_Main__SnapConnection) lastInsertId() int {
	return 2
}

func (self *_Main__SnapConnection) dbName() *string {
	return hxrt.StringFromLiteral("snap")
}

func (self *_Main__SnapConnection) startTransaction() {
}

func (self *_Main__SnapConnection) commit() {
}

func (self *_Main__SnapConnection) rollback() {
}

type I__Main__SnapResultSet interface {
	get_length() int
	get_nfields() int
	hasNext() bool
	next() any
	results() *haxe__ds__List
	getResult(n int) *string
	getIntResult(n int) int
	getFloatResult(n int) float64
	getFieldsNames() *hxrt.Array
}

type _Main__SnapResultSet struct {
	__hx_this I__Main__SnapResultSet
	length    int
	nfields   int
	rows      *hxrt.Array
	index     int
}

func (self *_Main__SnapResultSet) HxFunctionReceiver() any {
	return self.__hx_this
}

func New__Main__SnapResultSet(rows *hxrt.Array) *_Main__SnapResultSet {
	return hx_init__Main__SnapResultSet(hx_alloc__Main__SnapResultSet(nil), rows)
}

func hx_alloc__Main__SnapResultSet(owner any) *_Main__SnapResultSet {
	self := &_Main__SnapResultSet{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I__Main__SnapResultSet)
	}
	return self
}

func hx_init__Main__SnapResultSet(self *_Main__SnapResultSet, rows *hxrt.Array) *_Main__SnapResultSet {
	self.rows = rows
	self.index = 0
	return self
}

func (self *_Main__SnapResultSet) get_length() int {
	return int((hxrt.Int32Wrap(self.rows.Len()) - hxrt.Int32Wrap(self.index)))
}

func (self *_Main__SnapResultSet) get_nfields() int {
	return 1
}

func (self *_Main__SnapResultSet) hasNext() bool {
	return (self.index < self.rows.Len())
}

func (self *_Main__SnapResultSet) next() any {
	hx_post_8 := self.index
	self.index = int(int32((self.index + 1)))
	return self.rows.Get(hx_post_8)
}

func (self *_Main__SnapResultSet) results() *haxe__ds__List {
	out := New_haxe__ds__List()
	for self.hasNext() {
		out.__hx_this.add(self.next())
	}
	return out
}

func (self *_Main__SnapResultSet) getResult(n int) *string {
	return hxrt.StdString(hx_recover_class_1(self.rows.Get(int((hxrt.Int32Wrap(self.index) - hxrt.Int32Wrap(1))))).value)
}

func (self *_Main__SnapResultSet) getIntResult(n int) int {
	return hx_recover_class_1(self.rows.Get(int((hxrt.Int32Wrap(self.index) - hxrt.Int32Wrap(1))))).value
}

func (self *_Main__SnapResultSet) getFloatResult(n int) float64 {
	return (float64(hx_recover_class_1(self.rows.Get(int((hxrt.Int32Wrap(self.index) - hxrt.Int32Wrap(1))))).value) + 0.0)
}

func (self *_Main__SnapResultSet) getFieldsNames() *hxrt.Array {
	return hxrt.NewArray(hxrt.StringFromLiteral("value"))
}

type I__Main__SnapRow interface {
}

type _Main__SnapRow struct {
	__hx_this I__Main__SnapRow
	value     int
}

func (self *_Main__SnapRow) HxFunctionReceiver() any {
	return self.__hx_this
}

func New__Main__SnapRow(value int) *_Main__SnapRow {
	return hx_init__Main__SnapRow(hx_alloc__Main__SnapRow(nil), value)
}

func hx_alloc__Main__SnapRow(owner any) *_Main__SnapRow {
	self := &_Main__SnapRow{}
	if owner == nil {
		self.__hx_this = self
	} else {
		self.__hx_this = owner.(I__Main__SnapRow)
	}
	return self
}

func hx_init__Main__SnapRow(self *_Main__SnapRow, value int) *_Main__SnapRow {
	self.value = value
	return self
}

func hx_recover_class_0(value any) *haxe__Exception {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__exceptions__NotImplementedException:
		return source.haxe__exceptions__PosException.haxe__Exception
	case *haxe__Exception:
		return source
	case *haxe__exceptions__PosException:
		return source.haxe__Exception
	case *haxe__ValueException:
		return source.haxe__Exception
	default:
		return value.(*haxe__Exception)
	}
}

func hx_recover_class_1(value any) *_Main__SnapRow {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *_Main__SnapRow:
		return source
	default:
		return value.(*_Main__SnapRow)
	}
}

func hx_recover_class_2(value any) *haxe__io__Eof {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__io__Eof:
		return source
	default:
		return value.(*haxe__io__Eof)
	}
}

func hx_recover_class_3(value any) *haxe__ValueException {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ValueException:
		return source
	default:
		return value.(*haxe__ValueException)
	}
}

func hx_recover_class_4(value any) *haxe__ds___List__GoListIterator {
	value = hxrt.ObjectSource(value)
	if value == nil {
		return nil
	}
	switch source := value.(type) {
	case *haxe__ds___List__GoListIterator:
		return source
	default:
		return value.(*haxe__ds___List__GoListIterator)
	}
}
