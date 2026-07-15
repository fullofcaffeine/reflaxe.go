package main

import "snapshot/hxrt"

func sys__db__Mysql_connect(params map[string]any) sys__db__Connection {
	hxrt.Throw(New_haxe__exceptions__NotImplementedException(hxrt.StringFromLiteral("Not implemented for this platform"), nil, func() map[string]any {
		hx_obj_18 := map[string]any{}
		hx_obj_18["fileName"] = hxrt.StringFromLiteral("sys/db/Mysql.hx")
		hx_obj_18["lineNumber"] = 34
		hx_obj_18["className"] = hxrt.StringFromLiteral("sys.db.Mysql")
		hx_obj_18["methodName"] = hxrt.StringFromLiteral("connect")
		return hx_obj_18
	}()))
	var hx_throw_zero_19 sys__db__Connection
	return hx_throw_zero_19
	return nil
}
