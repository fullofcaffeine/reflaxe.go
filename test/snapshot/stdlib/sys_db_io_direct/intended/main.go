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
		sys__db__Mysql_connect(func() map[string]any {
			hx_obj_3 := map[string]any{}
			hx_obj_3["host"] = hxrt.StringFromLiteral("localhost")
			hx_obj_3["user"] = hxrt.StringFromLiteral("u")
			hx_obj_3["pass"] = hxrt.StringFromLiteral("p")
			return hx_obj_3
		}())
	}, func(hx_caught_1 any) {
		e := hxrt.ExceptionCaught(hx_caught_1)
		_ = e
	})
	hxrt.TryCatch(func() {
		sys__db__Sqlite_open(hxrt.StringFromLiteral("db.sqlite"))
	}, func(hx_caught_4 any) {
		e_1 := hxrt.ExceptionCaught(hx_caught_4)
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

func New__Main__SnapConnection() *_Main__SnapConnection {
	self := &_Main__SnapConnection{}
	self.__hx_this = self
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

func New__Main__SnapResultSet(rows *hxrt.Array) *_Main__SnapResultSet {
	self := &_Main__SnapResultSet{}
	self.__hx_this = self
	self.rows = rows
	self.index = 0
	return self
}

func (self *_Main__SnapResultSet) get_length() int {
	return int(int32((hxrt.Int32Wrap(self.rows.Len()) - hxrt.Int32Wrap(self.index))))
}

func (self *_Main__SnapResultSet) get_nfields() int {
	return 1
}

func (self *_Main__SnapResultSet) hasNext() bool {
	return (self.index < self.rows.Len())
}

func (self *_Main__SnapResultSet) next() any {
	hx_post_6 := self.index
	self.index = int(int32((self.index + 1)))
	return self.rows.Get(hx_post_6)
}

func (self *_Main__SnapResultSet) results() *haxe__ds__List {
	out := New_haxe__ds__List()
	for self.hasNext() {
		out.__hx_this.add(self.next())
	}
	return out
}

func (self *_Main__SnapResultSet) getResult(n int) *string {
	return hxrt.StdString(func(hx_value_7 any) *_Main__SnapRow {
		if hx_value_7 == nil {
			var hx_zero_8 *_Main__SnapRow
			return hx_zero_8
		}
		return hx_value_7.(*_Main__SnapRow)
	}(self.rows.Get(int(int32((hxrt.Int32Wrap(self.index) - hxrt.Int32Wrap(1)))))).value)
}

func (self *_Main__SnapResultSet) getIntResult(n int) int {
	return func(hx_value_9 any) *_Main__SnapRow {
		if hx_value_9 == nil {
			var hx_zero_10 *_Main__SnapRow
			return hx_zero_10
		}
		return hx_value_9.(*_Main__SnapRow)
	}(self.rows.Get(int(int32((hxrt.Int32Wrap(self.index) - hxrt.Int32Wrap(1)))))).value
}

func (self *_Main__SnapResultSet) getFloatResult(n int) float64 {
	return (float64(func(hx_value_11 any) *_Main__SnapRow {
		if hx_value_11 == nil {
			var hx_zero_12 *_Main__SnapRow
			return hx_zero_12
		}
		return hx_value_11.(*_Main__SnapRow)
	}(self.rows.Get(int(int32((hxrt.Int32Wrap(self.index) - hxrt.Int32Wrap(1)))))).value) + 0.0)
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

func New__Main__SnapRow(value int) *_Main__SnapRow {
	self := &_Main__SnapRow{}
	self.__hx_this = self
	self.value = value
	return self
}
