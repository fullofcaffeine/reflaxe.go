package main

import "examples_tui_todo_portable/hxrt"

var InteractiveCli_STATE_FILE *string = hxrt.StringFromLiteral(".tui_todo_state.txt")

func InteractiveCli_clearState() {
	hxrt.TryCatch(func() {
		sys__io__File_saveContent(hxrt.StringFromLiteral(".tui_todo_state.txt"), hxrt.StringFromLiteral(""))
	}, func(hx_caught_6 any) {
		hx_tmp := hxrt.ExceptionCaught(hx_caught_6)
		_ = hx_tmp
	})
}

func InteractiveCli_decodeTags(raw *string) []*string {
	out := []*string{}
	if hxrt.StringEqualStringPtr(raw, hxrt.StringFromLiteral("")) {
		return out
	}
	values := InteractiveCli_splitEscaped(raw, 44)
	_g := 0
	for _g < len(values) {
		tag := values[_g]
		_g = int(int32((_g + 1)))
		if !hxrt.StringEqualStringPtr(tag, hxrt.StringFromLiteral("")) {
			out = append(out, tag)
		}
	}
	return out
}

func InteractiveCli_decodeToken(raw *string) *string {
	return StringTools_replace(raw, hxrt.StringFromLiteral("_"), hxrt.StringFromLiteral(" "))
}

func InteractiveCli_encodeField(raw *string) *string {
	out := New_haxe__io__BytesBuffer()
	bytes := haxe__io__Bytes_ofString(raw)
	i := 0
	for i < bytes.length {
		code := bytes.b[i]
		if code == 92 {
			hx_arr_9 := out.b
			hx_arr_9 = append(hx_arr_9, (92 & 255))
			out.b = hx_arr_9
			hx_arr_10 := out.b
			hx_arr_10 = append(hx_arr_10, (92 & 255))
			out.b = hx_arr_10
		} else {
			if code == 9 {
				hx_arr_11 := out.b
				hx_arr_11 = append(hx_arr_11, (92 & 255))
				out.b = hx_arr_11
				hx_arr_12 := out.b
				hx_arr_12 = append(hx_arr_12, (116 & 255))
				out.b = hx_arr_12
			} else {
				if code == 10 {
					hx_arr_13 := out.b
					hx_arr_13 = append(hx_arr_13, (92 & 255))
					out.b = hx_arr_13
					hx_arr_14 := out.b
					hx_arr_14 = append(hx_arr_14, (110 & 255))
					out.b = hx_arr_14
				} else {
					if code == 44 {
						hx_arr_15 := out.b
						hx_arr_15 = append(hx_arr_15, (92 & 255))
						out.b = hx_arr_15
						hx_arr_16 := out.b
						hx_arr_16 = append(hx_arr_16, (99 & 255))
						out.b = hx_arr_16
					} else {
						hx_arr_17 := out.b
						hx_arr_17 = append(hx_arr_17, (code & 255))
						out.b = hx_arr_17
					}
				}
			}
		}
		i = int(int32((i + 1)))
	}
	return out.getBytes().toString()
}

func InteractiveCli_encodeTags(tags []*string) *string {
	out := hxrt.StringFromLiteral("")
	first := true
	_g := 0
	for _g < len(tags) {
		tag := tags[_g]
		_g = int(int32((_g + 1)))
		if !first {
			out = hxrt.StringConcatStringPtr(out, hxrt.StringFromLiteral(","))
		}
		out = hxrt.StringConcatStringPtr(out, InteractiveCli_encodeField(tag))
		first = false
	}
	return out
}

func InteractiveCli_failUsage(message *string) {
	hxrt.Println(any(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error: "), message)))
	hxrt.Println(any(hxrt.StringFromLiteral("run `help` for command syntax")))
}

func InteractiveCli_listIndex(values []*string, index int) *string {
	if (index < 0) || (index >= len(values)) {
		return hxrt.StringFromLiteral("")
	}
	return values[index]
}

func InteractiveCli_loadState(app *app__TodoApp) {
	hx_try_return_18 := false
	hxrt.TryCatch(func() {
		raw := sys__io__File_getContent(hxrt.StringFromLiteral(".tui_todo_state.txt"))
		if hxrt.StringEqualStringPtr(raw, hxrt.StringFromLiteral("")) {
			hx_try_return_18 = true
			return
		}
		lines := InteractiveCli_splitRaw(raw, 10)
		_g := 0
		for _g < len(lines) {
			line := lines[_g]
			_g = int(int32((_g + 1)))
			if hxrt.StringEqualStringPtr(line, hxrt.StringFromLiteral("")) {
				continue
			}
			fields := InteractiveCli_splitEscaped(line, 9)
			title := InteractiveCli_listIndex(fields, 0)
			priority := InteractiveCli_parsePositiveInt(InteractiveCli_listIndex(fields, 1))
			if priority < 0 {
				priority = 0
			}
			done := hxrt.StringEqualStringPtr(InteractiveCli_listIndex(fields, 2), hxrt.StringFromLiteral("1"))
			id := app.add(title, priority)
			if done {
				app.toggle(id)
			}
			tags := InteractiveCli_decodeTags(InteractiveCli_listIndex(fields, 3))
			_g_1 := 0
			for _g_1 < len(tags) {
				tag := tags[_g_1]
				_g_1 = int(int32((_g_1 + 1)))
				app.tag(id, tag)
			}
		}
	}, func(hx_caught_19 any) {
		hx_tmp := hxrt.ExceptionCaught(hx_caught_19)
		_ = hx_tmp
		hx_try_return_18 = true
		return
	})
	if hx_try_return_18 {
		return
	}
}

func InteractiveCli_parsePositiveInt(raw *string) int {
	if hxrt.StringEqualStringPtr(raw, hxrt.StringFromLiteral("")) {
		return -1
	}
	bytes := haxe__io__Bytes_ofString(raw)
	value := 0
	i := 0
	for i < bytes.length {
		code := bytes.b[i]
		if (code < 48) || (code > 57) {
			return -1
		}
		value = int(int32((hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(value) * hxrt.Int32Wrap(10))))) + hxrt.Int32Wrap(int(int32((hxrt.Int32Wrap(code) - hxrt.Int32Wrap(48))))))))
		i = int(int32((i + 1)))
	}
	return value
}

func InteractiveCli_printHelp(runtime profile__TodoRuntime) {
	hxrt.Println(any(hxrt.StringFromLiteral("commands:")))
	hxrt.Println(any(hxrt.StringFromLiteral("  help")))
	hxrt.Println(any(hxrt.StringFromLiteral("  reset")))
	hxrt.Println(any(hxrt.StringFromLiteral("  list")))
	hxrt.Println(any(hxrt.StringFromLiteral("  summary")))
	hxrt.Println(any(hxrt.StringFromLiteral("  diag")))
	hxrt.Println(any(hxrt.StringFromLiteral("  add <priority> <title_token>")))
	hxrt.Println(any(hxrt.StringFromLiteral("  toggle <id>")))
	hxrt.Println(any(hxrt.StringFromLiteral("  tag <id> <tag_token>")))
	hxrt.Println(any(hxrt.StringFromLiteral("  batch <priority> <title1_token> <title2_token>")))
	hxrt.Println(any(hxrt.StringFromLiteral("token note: use '_' instead of spaces (example: Wire_release_artifacts)")))
	hxrt.Println(any(hxrt.StringFromLiteral("state file: .tui_todo_state.txt (current directory)")))
}

func InteractiveCli_printUsage(runtime profile__TodoRuntime) {
	var v any = any(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("tui_todo command session ("), runtime.profileId()), hxrt.StringFromLiteral(")")))
	hxrt.Println(v)
	hxrt.Println(any(hxrt.StringFromLiteral("run scripted contract mode with: --scripted")))
	hxrt.Println(any(hxrt.StringFromLiteral("commands:")))
	hxrt.Println(any(hxrt.StringFromLiteral("  tui_todo reset")))
	hxrt.Println(any(hxrt.StringFromLiteral("  tui_todo help")))
	hxrt.Println(any(hxrt.StringFromLiteral("  tui_todo add 2 Write_profile_docs tag 1 docs list")))
	hxrt.Println(any(hxrt.StringFromLiteral("  tui_todo batch 3 Ship_generated_go_sync Add_binary_matrix list")))
	hxrt.Println(any(hxrt.StringFromLiteral("generated-source invocation:")))
	hxrt.Println(any(hxrt.StringFromLiteral("  go run . <command...>")))
	hxrt.Println(any(hxrt.StringFromLiteral("state file: .tui_todo_state.txt (current directory)")))
}

func InteractiveCli_run(runtime profile__TodoRuntime) {
	app := New_app__TodoApp(runtime)
	InteractiveCli_loadState(app)
	args := hxrt.SysArgs()
	if len(args) == 0 {
		InteractiveCli_printUsage(runtime)
		return
	}
	i := 0
	for i < len(args) {
		cmd := args[i]
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("reset")) {
			app = New_app__TodoApp(runtime)
			InteractiveCli_clearState()
			hxrt.Println(any(hxrt.StringFromLiteral("ok reset")))
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("help")) {
			InteractiveCli_printHelp(runtime)
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("list")) {
			var v any = any(app.render())
			hxrt.Println(v)
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("summary")) {
			var v_1 any = any(app.baselineSignature())
			hxrt.Println(v_1)
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("diag")) {
			var v_2 any = any(app.diagnostics())
			hxrt.Println(v_2)
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("add")) {
			if int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(2)))) >= len(args) {
				InteractiveCli_failUsage(hxrt.StringFromLiteral("add requires <priority> <title_token>"))
				return
			}
			priority := InteractiveCli_parsePositiveInt(args[int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(1))))])
			if priority < 0 {
				InteractiveCli_failUsage(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("invalid priority: "), args[int(int32((hxrt.Int32Wrap(i)+hxrt.Int32Wrap(1))))]))
				return
			}
			title := InteractiveCli_decodeToken(args[int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(2))))])
			app.add(title, priority)
			InteractiveCli_saveState(app)
			hxrt.Println(any(hxrt.StringFromLiteral("ok add")))
			i = int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(3))))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("toggle")) {
			if int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(1)))) >= len(args) {
				InteractiveCli_failUsage(hxrt.StringFromLiteral("toggle requires <id>"))
				return
			}
			id := InteractiveCli_parsePositiveInt(args[int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(1))))])
			if id < 0 {
				InteractiveCli_failUsage(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("invalid id: "), args[int(int32((hxrt.Int32Wrap(i)+hxrt.Int32Wrap(1))))]))
				return
			}
			if app.toggle(id) {
				InteractiveCli_saveState(app)
				hxrt.Println(any(hxrt.StringFromLiteral("ok toggle")))
			} else {
				hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("missing id: "), id)))
			}
			i = int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(2))))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("tag")) {
			if int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(2)))) >= len(args) {
				InteractiveCli_failUsage(hxrt.StringFromLiteral("tag requires <id> <tag_token>"))
				return
			}
			id_1 := InteractiveCli_parsePositiveInt(args[int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(1))))])
			if id_1 < 0 {
				InteractiveCli_failUsage(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("invalid id: "), args[int(int32((hxrt.Int32Wrap(i)+hxrt.Int32Wrap(1))))]))
				return
			}
			tag := InteractiveCli_decodeToken(args[int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(2))))])
			if app.tag(id_1, tag) {
				InteractiveCli_saveState(app)
				hxrt.Println(any(hxrt.StringFromLiteral("ok tag")))
			} else {
				hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("missing id: "), id_1)))
			}
			i = int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(3))))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("batch")) {
			if int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(3)))) >= len(args) {
				InteractiveCli_failUsage(hxrt.StringFromLiteral("batch requires <priority> <title1_token> <title2_token>"))
				return
			}
			priority_1 := InteractiveCli_parsePositiveInt(args[int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(1))))])
			if priority_1 < 0 {
				InteractiveCli_failUsage(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("invalid priority: "), args[int(int32((hxrt.Int32Wrap(i)+hxrt.Int32Wrap(1))))]))
				return
			}
			titles := []*string{}
			titles = append(titles, InteractiveCli_decodeToken(args[int(int32((hxrt.Int32Wrap(i)+hxrt.Int32Wrap(2))))]))
			titles = append(titles, InteractiveCli_decodeToken(args[int(int32((hxrt.Int32Wrap(i)+hxrt.Int32Wrap(3))))]))
			added := app.addMany(titles, priority_1)
			if added > 0 {
				InteractiveCli_saveState(app)
			}
			hxrt.Println(any(hxrt.StringConcatAny(hxrt.StringFromLiteral("ok batch added="), added)))
			i = int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(4))))
			continue
		}
		InteractiveCli_failUsage(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("unknown command: "), cmd))
		return
	}
}

func InteractiveCli_saveState(app *app__TodoApp) {
	items := app.items()
	out := hxrt.StringFromLiteral("")
	_g := 0
	for _g < len(items) {
		item := items[_g]
		_g = int(int32((_g + 1)))
		out = hxrt.StringConcatStringPtr(out, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(InteractiveCli_encodeField(item.title), hxrt.StringFromLiteral("\t")), item.priority), hxrt.StringFromLiteral("\t")), func() *string {
			var hx_if_23 *string
			if item.done {
				hx_if_23 = hxrt.StringFromLiteral("1")
			} else {
				hx_if_23 = hxrt.StringFromLiteral("0")
			}
			return hx_if_23
		}()), hxrt.StringFromLiteral("\t")), InteractiveCli_encodeTags(item.tags)), hxrt.StringFromLiteral("\n")))
	}
	sys__io__File_saveContent(hxrt.StringFromLiteral(".tui_todo_state.txt"), out)
}

func InteractiveCli_splitEscaped(raw *string, separatorCode int) []*string {
	out := []*string{}
	current := New_haxe__io__BytesBuffer()
	bytes := haxe__io__Bytes_ofString(raw)
	escaped := false
	i := 0
	for i < bytes.length {
		code := bytes.b[i]
		if escaped {
			if code == 116 {
				hx_arr_24 := current.b
				hx_arr_24 = append(hx_arr_24, (9 & 255))
				current.b = hx_arr_24
			} else {
				if code == 110 {
					hx_arr_25 := current.b
					hx_arr_25 = append(hx_arr_25, (10 & 255))
					current.b = hx_arr_25
				} else {
					if code == 99 {
						hx_arr_26 := current.b
						hx_arr_26 = append(hx_arr_26, (44 & 255))
						current.b = hx_arr_26
					} else {
						if code == 92 {
							hx_arr_27 := current.b
							hx_arr_27 = append(hx_arr_27, (92 & 255))
							current.b = hx_arr_27
						} else {
							hx_arr_28 := current.b
							hx_arr_28 = append(hx_arr_28, (code & 255))
							current.b = hx_arr_28
						}
					}
				}
			}
			escaped = false
			i = int(int32((i + 1)))
			continue
		}
		if code == 92 {
			escaped = true
			i = int(int32((i + 1)))
			continue
		}
		if code == separatorCode {
			out = append(out, current.getBytes().toString())
			current = New_haxe__io__BytesBuffer()
			i = int(int32((i + 1)))
			continue
		}
		hx_arr_30 := current.b
		hx_arr_30 = append(hx_arr_30, (code & 255))
		current.b = hx_arr_30
		i = int(int32((i + 1)))
	}
	out = append(out, current.getBytes().toString())
	return out
}

func InteractiveCli_splitRaw(raw *string, separatorCode int) []*string {
	out := []*string{}
	current := New_haxe__io__BytesBuffer()
	bytes := haxe__io__Bytes_ofString(raw)
	i := 0
	for i < bytes.length {
		code := bytes.b[i]
		if code == separatorCode {
			out = append(out, current.getBytes().toString())
			current = New_haxe__io__BytesBuffer()
		} else {
			if code != 13 {
				hx_arr_33 := current.b
				hx_arr_33 = append(hx_arr_33, (code & 255))
				current.b = hx_arr_33
			}
		}
		i = int(int32((i + 1)))
	}
	out = append(out, current.getBytes().toString())
	return out
}
