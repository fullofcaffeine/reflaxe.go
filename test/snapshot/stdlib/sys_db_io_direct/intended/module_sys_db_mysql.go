package main

import "snapshot/hxrt"

func sys__db__Mysql_connect(params map[string]any) sys__db__Connection {
	hxrt.Throw(New_haxe__exceptions__NotImplementedException(hxrt.StringFromLiteral("Not implemented for this platform"), nil, func() map[string]any {
		hx_obj_13 := map[string]any{}
		hx_obj_13["fileName"] = hxrt.StringFromLiteral("sys/db/Mysql.hx")
		hx_obj_13["lineNumber"] = 34
		hx_obj_13["className"] = hxrt.StringFromLiteral("sys.db.Mysql")
		hx_obj_13["methodName"] = hxrt.StringFromLiteral("connect")
		return hx_obj_13
	}()))
	return nil
}
