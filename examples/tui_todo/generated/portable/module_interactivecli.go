package main

import "examples_tui_todo_portable/hxrt"

var InteractiveCli_STATE_FILE *string = hxrt.StringFromLiteral(".tui_todo_state.txt")

func InteractiveCli_clearState() {
	hxrt.TryCatch(func() {
		sys__io__File_saveContent(hxrt.StringFromLiteral(".tui_todo_state.txt"), hxrt.StringFromLiteral(""))
	}, func(hx_caught_5 any) {
		hx_tmp := hx_caught_5
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
			out.b = append(out.b, (92 & 255))
			out.b = append(out.b, (92 & 255))
		} else {
			if code == 9 {
				out.b = append(out.b, (92 & 255))
				out.b = append(out.b, (116 & 255))
			} else {
				if code == 10 {
					out.b = append(out.b, (92 & 255))
					out.b = append(out.b, (110 & 255))
				} else {
					if code == 44 {
						out.b = append(out.b, (92 & 255))
						out.b = append(out.b, (99 & 255))
					} else {
						out.b = append(out.b, (code & 255))
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
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error: "), message))
	hxrt.Println(hxrt.StringFromLiteral("run `help` for command syntax"))
}

func InteractiveCli_listIndex(values []*string, index int) *string {
	if (index < 0) || (index >= len(values)) {
		return hxrt.StringFromLiteral("")
	}
	return values[index]
}

func InteractiveCli_loadState(app *app__TodoApp) {
	hx_try_return_7 := false
	hxrt.TryCatch(func() {
		raw := sys__io__File_getContent(hxrt.StringFromLiteral(".tui_todo_state.txt"))
		if hxrt.StringEqualStringPtr(raw, hxrt.StringFromLiteral("")) {
			hx_try_return_7 = true
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
	}, func(hx_caught_8 any) {
		hx_tmp := hx_caught_8
		_ = hx_tmp
		hx_try_return_7 = true
		return
	})
	if hx_try_return_7 {
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
	hxrt.Println(hxrt.StringFromLiteral("commands:"))
	hxrt.Println(hxrt.StringFromLiteral("  help"))
	hxrt.Println(hxrt.StringFromLiteral("  reset"))
	hxrt.Println(hxrt.StringFromLiteral("  list"))
	hxrt.Println(hxrt.StringFromLiteral("  summary"))
	hxrt.Println(hxrt.StringFromLiteral("  diag"))
	hxrt.Println(hxrt.StringFromLiteral("  add <priority> <title_token>"))
	hxrt.Println(hxrt.StringFromLiteral("  toggle <id>"))
	hxrt.Println(hxrt.StringFromLiteral("  tag <id> <tag_token>"))
	if runtime.supportsBatchAdd() {
		hxrt.Println(hxrt.StringFromLiteral("  batch <priority> <title1_token> <title2_token>"))
	}
	hxrt.Println(hxrt.StringFromLiteral("token note: use '_' instead of spaces (example: Wire_release_artifacts)"))
	hxrt.Println(hxrt.StringFromLiteral("state file: .tui_todo_state.txt (current directory)"))
}

func InteractiveCli_printUsage(runtime profile__TodoRuntime) {
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("tui_todo command session ("), runtime.profileId()), hxrt.StringFromLiteral(")")))
	hxrt.Println(hxrt.StringFromLiteral("run scripted contract mode with: --scripted"))
	hxrt.Println(hxrt.StringFromLiteral("commands:"))
	hxrt.Println(hxrt.StringFromLiteral("  tui_todo reset"))
	hxrt.Println(hxrt.StringFromLiteral("  tui_todo help"))
	hxrt.Println(hxrt.StringFromLiteral("  tui_todo add 2 Write_profile_docs tag 1 docs list"))
	if runtime.supportsBatchAdd() {
		hxrt.Println(hxrt.StringFromLiteral("  tui_todo batch 3 Ship_generated_go_sync Add_binary_matrix list"))
	}
	hxrt.Println(hxrt.StringFromLiteral("generated-source invocation:"))
	hxrt.Println(hxrt.StringFromLiteral("  go run . <command...>"))
	hxrt.Println(hxrt.StringFromLiteral("state file: .tui_todo_state.txt (current directory)"))
}

func InteractiveCli_run(runtime profile__TodoRuntime) {
	app := New_app__TodoApp(runtime)
	InteractiveCli_loadState(app)
	args := Sys_args()
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
			hxrt.Println(hxrt.StringFromLiteral("ok reset"))
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("help")) {
			InteractiveCli_printHelp(runtime)
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("list")) {
			hxrt.Println(app.render())
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("summary")) {
			hxrt.Println(app.baselineSignature())
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("diag")) {
			hxrt.Println(app.diagnostics())
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
			hxrt.Println(hxrt.StringFromLiteral("ok add"))
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
				hxrt.Println(hxrt.StringFromLiteral("ok toggle"))
			} else {
				hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("missing id: "), id))
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
				hxrt.Println(hxrt.StringFromLiteral("ok tag"))
			} else {
				hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("missing id: "), id_1))
			}
			i = int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(3))))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("batch")) {
			if !runtime.supportsBatchAdd() {
				hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("batch not supported in "), runtime.profileId()))
				i = int(int32((i + 1)))
				continue
			}
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
			hxrt.Println(hxrt.StringConcatAny(hxrt.StringFromLiteral("ok batch added="), added))
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
			var hx_if_10 *string
			if item.done {
				hx_if_10 = hxrt.StringFromLiteral("1")
			} else {
				hx_if_10 = hxrt.StringFromLiteral("0")
			}
			return hx_if_10
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
				current.b = append(current.b, (9 & 255))
			} else {
				if code == 110 {
					current.b = append(current.b, (10 & 255))
				} else {
					if code == 99 {
						current.b = append(current.b, (44 & 255))
					} else {
						if code == 92 {
							current.b = append(current.b, (92 & 255))
						} else {
							current.b = append(current.b, (code & 255))
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
		current.b = append(current.b, (code & 255))
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
				current.b = append(current.b, (code & 255))
			}
		}
		i = int(int32((i + 1)))
	}
	out = append(out, current.getBytes().toString())
	return out
}
