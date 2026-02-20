package main

import (
	"bytes"
	"compress/zlib"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"examples_tui_todo_gopher/hxrt"
	"io"
	"math"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

func Harness_assertContract(runtime profile__TodoRuntime) *string {
	app := New_app__TodoApp(runtime)
	Harness_runBaseline(app)
	baseline := app.baselineSignature()
	if !hxrt.StringEqualStringPtr(baseline, hxrt.StringFromLiteral("open=1,done=1,total=2")) {
		hxrt.Throw(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("baseline drift: "), baseline))
		var hx_throw_zero_1 *string
		return hx_throw_zero_1
	}
	if runtime.supportsBatchAdd() {
		added := app.__hx_this.addMany(Harness_batchTitles(), 3)
		if (added != 2) || (app.__hx_this.totalCount() != 4) {
			hxrt.Throw(hxrt.StringFromLiteral("batch add drift"))
			var hx_throw_zero_2 *string
			return hx_throw_zero_2
		}
	} else {
		if app.__hx_this.totalCount() != 2 {
			hxrt.Throw(hxrt.StringFromLiteral("portable total drift"))
			var hx_throw_zero_3 *string
			return hx_throw_zero_3
		}
	}
	if runtime.supportsDiagnostics() {
		diag := app.__hx_this.diagnostics()
		if !hxrt.StringEqualStringPtr(diag, hxrt.StringFromLiteral("p1=1,completed=1")) {
			hxrt.Throw(hxrt.StringFromLiteral("missing diagnostics"))
			var hx_throw_zero_4 *string
			return hx_throw_zero_4
		}
	}
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("OK "), runtime.profileId())
}

func Harness_batchTitles() *haxe__ds__List {
	out := New_haxe__ds__List()
	out.add(hxrt.StringFromLiteral("Ship generated-go sync"))
	out.add(hxrt.StringFromLiteral("Add binary matrix"))
	return out
}

func Harness_run(runtime profile__TodoRuntime) *string {
	app := New_app__TodoApp(runtime)
	baselineView := Harness_runBaseline(app)
	_ = baselineView
	baseline := app.baselineSignature()
	_ = baseline
	extras := hxrt.StringFromLiteral("batch_add=0")
	if runtime.supportsBatchAdd() {
		added := app.addMany(Harness_batchTitles(), 3)
		extras = hxrt.StringConcatAny(hxrt.StringFromLiteral("batch_add="), added)
	}
	extras = hxrt.StringConcatStringPtr(extras, hxrt.StringConcatStringPtr(hxrt.StringFromLiteral(",diag="), app.__hx_this.diagnostics()))
	return hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("profile="), runtime.profileId()), hxrt.StringFromLiteral("\nbaseline=")), baseline), hxrt.StringFromLiteral("\nbaseline_view:\n")), baselineView), hxrt.StringFromLiteral("\nfinal_view:\n")), app.__hx_this.render()), hxrt.StringFromLiteral("\nextras=")), extras)
}

func Harness_runBaseline(app *app__TodoApp) *string {
	app.__hx_this.add(hxrt.StringFromLiteral("Write profile docs"), 2)
	app.__hx_this.add(hxrt.StringFromLiteral("Backfill regression snapshots"), 1)
	app.__hx_this.toggle(2)
	app.__hx_this.tag(1, hxrt.StringFromLiteral("docs"))
	app.__hx_this.tag(2, hxrt.StringFromLiteral("tests"))
	return app.__hx_this.render()
}

var InteractiveCli_STATE_FILE *string = hxrt.StringFromLiteral(".tui_todo_state.txt")

func InteractiveCli_clearState() {
	hxrt.TryCatch(func() {
		sys__io__File_saveContent(hxrt.StringFromLiteral(".tui_todo_state.txt"), hxrt.StringFromLiteral(""))
	}, func(hx_caught_5 any) {
		hx_tmp := hx_caught_5
		_ = hx_tmp
	})
}

func InteractiveCli_decodeTags(raw *string) *haxe__ds__List {
	out := New_haxe__ds__List()
	if hxrt.StringEqualStringPtr(raw, hxrt.StringFromLiteral("")) {
		return out
	}
	values := InteractiveCli_splitEscaped(raw, 44)
	count := values.length
	_ = count
	i := 0
	for i < count {
		value := func(hx_value_7 any) *string {
			if hx_value_7 == nil {
				var hx_zero_8 *string
				return hx_zero_8
			}
			return hx_value_7.(*string)
		}(values.pop())
		if hxrt.StringEqualStringPtr(value, nil) {
			break
		}
		tag := value
		if !hxrt.StringEqualStringPtr(tag, hxrt.StringFromLiteral("")) {
			out.add(tag)
		}
		values.add(tag)
		i = int(int32((i + 1)))
	}
	return out
}

func InteractiveCli_decodeToken(raw *string) *string {
	return StringTools_replace(raw, hxrt.StringFromLiteral("_"), hxrt.StringFromLiteral(" "))
}

func InteractiveCli_encodeField(raw *string) *string {
	out := New_haxe__io__BytesBuffer()
	_ = out
	bytes := haxe__io__Bytes_ofString(raw)
	_ = bytes
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

func InteractiveCli_encodeTags(tags *haxe__ds__List) *string {
	out := hxrt.StringFromLiteral("")
	_ = out
	first := true
	_ = first
	count := tags.length
	_ = count
	i := 0
	for i < count {
		value := func(hx_value_9 any) *string {
			if hx_value_9 == nil {
				var hx_zero_10 *string
				return hx_zero_10
			}
			return hx_value_9.(*string)
		}(tags.pop())
		if hxrt.StringEqualStringPtr(value, nil) {
			break
		}
		tag := value
		_ = tag
		if !first {
			out = hxrt.StringConcatStringPtr(out, hxrt.StringFromLiteral(","))
		}
		out = hxrt.StringConcatStringPtr(out, InteractiveCli_encodeField(tag))
		tags.add(tag)
		first = false
		i = int(int32((i + 1)))
	}
	return out
}

func InteractiveCli_failUsage(message *string) {
	hxrt.Println(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("error: "), message))
	hxrt.Println(hxrt.StringFromLiteral("run `help` for command syntax"))
}

func InteractiveCli_listIndex(values *haxe__ds__List, index int) *string {
	count := values.length
	_ = count
	i := 0
	_ = i
	out := hxrt.StringFromLiteral("")
	for i < count {
		value := func(hx_value_11 any) *string {
			if hx_value_11 == nil {
				var hx_zero_12 *string
				return hx_zero_12
			}
			return hx_value_11.(*string)
		}(values.pop())
		if hxrt.StringEqualStringPtr(value, nil) {
			break
		}
		entry := value
		if i == index {
			out = entry
		}
		values.add(entry)
		i = int(int32((i + 1)))
	}
	return out
}

func InteractiveCli_loadState(app *app__TodoApp) {
	hxrt.TryCatch(func() {
		raw := sys__io__File_getContent(hxrt.StringFromLiteral(".tui_todo_state.txt"))
		if hxrt.StringEqualStringPtr(raw, hxrt.StringFromLiteral("")) {
			return
		}
		lines := InteractiveCli_splitRaw(raw, 10)
		count := lines.length
		_ = count
		i := 0
		for i < count {
			lineValue := func(hx_value_15 any) *string {
				if hx_value_15 == nil {
					var hx_zero_16 *string
					return hx_zero_16
				}
				return hx_value_15.(*string)
			}(lines.pop())
			if hxrt.StringEqualStringPtr(lineValue, nil) {
				break
			}
			line := lineValue
			if hxrt.StringEqualStringPtr(line, hxrt.StringFromLiteral("")) {
				lines.add(line)
				i = int(int32((i + 1)))
				continue
			}
			fields := InteractiveCli_splitEscaped(line, 9)
			title := InteractiveCli_listIndex(fields, 0)
			_ = title
			priority := InteractiveCli_parsePositiveInt(InteractiveCli_listIndex(fields, 1))
			if priority < 0 {
				priority = 0
			}
			done := hxrt.StringEqualStringPtr(InteractiveCli_listIndex(fields, 2), hxrt.StringFromLiteral("1"))
			_ = done
			id := app.__hx_this.add(title, priority)
			if done {
				app.__hx_this.toggle(id)
			}
			tags := InteractiveCli_decodeTags(InteractiveCli_listIndex(fields, 3))
			tagCount := tags.length
			_ = tagCount
			j := 0
			for j < tagCount {
				tagValue := func(hx_value_17 any) *string {
					if hx_value_17 == nil {
						var hx_zero_18 *string
						return hx_zero_18
					}
					return hx_value_17.(*string)
				}(tags.pop())
				if hxrt.StringEqualStringPtr(tagValue, nil) {
					break
				}
				tag := tagValue
				app.__hx_this.tag(id, tag)
				tags.add(tag)
				j = int(int32((j + 1)))
			}
			lines.add(line)
			i = int(int32((i + 1)))
		}
	}, func(hx_caught_13 any) {
		hx_tmp := hx_caught_13
		_ = hx_tmp
		return
	})
}

func InteractiveCli_parsePositiveInt(raw *string) int {
	if hxrt.StringEqualStringPtr(raw, hxrt.StringFromLiteral("")) {
		return -1
	}
	bytes := haxe__io__Bytes_ofString(raw)
	_ = bytes
	value := 0
	_ = value
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
			hxrt.Println(app.__hx_this.render())
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("summary")) {
			hxrt.Println(app.__hx_this.baselineSignature())
			i = int(int32((i + 1)))
			continue
		}
		if hxrt.StringEqualStringPtr(cmd, hxrt.StringFromLiteral("diag")) {
			hxrt.Println(app.__hx_this.diagnostics())
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
			app.__hx_this.add(title, priority)
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
			if app.__hx_this.toggle(id) {
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
			if app.__hx_this.tag(id_1, tag) {
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
			titles := New_haxe__ds__List()
			titles.add(InteractiveCli_decodeToken(args[int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(2))))]))
			titles.add(InteractiveCli_decodeToken(args[int(int32((hxrt.Int32Wrap(i) + hxrt.Int32Wrap(3))))]))
			added := app.__hx_this.addMany(titles, priority_1)
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
	items := app.__hx_this.items()
	_ = items
	out := hxrt.StringFromLiteral("")
	_ = out
	count := items.length
	_ = count
	i := 0
	for i < count {
		raw := func(hx_value_19 any) *model__TodoItem {
			if hx_value_19 == nil {
				var hx_zero_20 *model__TodoItem
				return hx_zero_20
			}
			return hx_value_19.(*model__TodoItem)
		}(items.pop())
		if hxrt.StringEqualAny(raw, nil) {
			break
		}
		item := raw
		out = hxrt.StringConcatStringPtr(out, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(InteractiveCli_encodeField(item.title), hxrt.StringFromLiteral("\t")), item.priority), hxrt.StringFromLiteral("\t")), func() *string {
			var hx_if_21 *string
			if item.done {
				hx_if_21 = hxrt.StringFromLiteral("1")
			} else {
				hx_if_21 = hxrt.StringFromLiteral("0")
			}
			return hx_if_21
		}()), hxrt.StringFromLiteral("\t")), InteractiveCli_encodeTags(item.tags)), hxrt.StringFromLiteral("\n")))
		items.add(item)
		i = int(int32((i + 1)))
	}
	sys__io__File_saveContent(hxrt.StringFromLiteral(".tui_todo_state.txt"), out)
}

func InteractiveCli_splitEscaped(raw *string, separatorCode int) *haxe__ds__List {
	out := New_haxe__ds__List()
	_ = out
	current := New_haxe__io__BytesBuffer()
	_ = current
	bytes := haxe__io__Bytes_ofString(raw)
	_ = bytes
	escaped := false
	_ = escaped
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
			out.add(current.getBytes().toString())
			current = New_haxe__io__BytesBuffer()
			i = int(int32((i + 1)))
			continue
		}
		current.b = append(current.b, (code & 255))
		i = int(int32((i + 1)))
	}
	out.add(current.getBytes().toString())
	return out
}

func InteractiveCli_splitRaw(raw *string, separatorCode int) *haxe__ds__List {
	out := New_haxe__ds__List()
	_ = out
	current := New_haxe__io__BytesBuffer()
	_ = current
	bytes := haxe__io__Bytes_ofString(raw)
	_ = bytes
	i := 0
	for i < bytes.length {
		code := bytes.b[i]
		if code == separatorCode {
			out.add(current.getBytes().toString())
			current = New_haxe__io__BytesBuffer()
		} else {
			if code != 13 {
				current.b = append(current.b, (code & 255))
			}
		}
		i = int(int32((i + 1)))
	}
	out.add(current.getBytes().toString())
	return out
}

func hasArg(flag *string) bool {
	_g := 0
	_ = _g
	_g1 := Sys_args()
	for _g < len(_g1) {
		arg := _g1[_g]
		_ = arg
		_g = int(int32((_g + 1)))
		if hxrt.StringEqualStringPtr(arg, flag) {
			return true
		}
	}
	return false
}

func main() {
	var runtime profile__TodoRuntime = profile__RuntimeFactory_create()
	if hasArg(hxrt.StringFromLiteral("--scripted")) {
		hxrt.Println(Harness_run(runtime))
	} else {
		InteractiveCli_run(runtime)
	}
}

type I_app__TodoApp interface {
	add(title *string, priority int) int
	addMany(titles *haxe__ds__List, priority int) int
	toggle(id int) bool
	tag(id int, tag *string) bool
	baselineSignature() *string
	totalCount() int
	openCount() int
	doneCount() int
	diagnostics() *string
	render() *string
	items() *haxe__ds__List
}

type app__TodoApp struct {
	__hx_this I_app__TodoApp
	runtime   profile__TodoRuntime
	store     *model__TodoStore
}

func New_app__TodoApp(runtime profile__TodoRuntime) *app__TodoApp {
	self := &app__TodoApp{}
	self.__hx_this = self
	self.runtime = runtime
	self.store = New_model__TodoStore()
	return self
}

func (self *app__TodoApp) add(title *string, priority int) int {
	item := self.store.__hx_this.add(self.runtime.normalizeTitle(title), priority)
	return item.id
}

func (self *app__TodoApp) addMany(titles *haxe__ds__List, priority int) int {
	if !self.runtime.supportsBatchAdd() {
		return 0
	}
	added := 0
	_ = added
	count := titles.length
	_ = count
	i := 0
	for i < count {
		raw := func(hx_value_22 any) *string {
			if hx_value_22 == nil {
				var hx_zero_23 *string
				return hx_zero_23
			}
			return hx_value_22.(*string)
		}(titles.pop())
		if hxrt.StringEqualStringPtr(raw, nil) {
			break
		}
		title := raw
		self.add(title, priority)
		titles.add(title)
		added = int(int32((added + 1)))
		i = int(int32((i + 1)))
	}
	return added
}

func (self *app__TodoApp) toggle(id int) bool {
	return self.store.__hx_this.toggle(id)
}

func (self *app__TodoApp) tag(id int, tag *string) bool {
	return self.store.__hx_this.addTag(id, self.runtime.normalizeTag(tag))
}

func (self *app__TodoApp) baselineSignature() *string {
	return hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringFromLiteral("open="), self.openCount()), hxrt.StringFromLiteral(",done=")), self.doneCount()), hxrt.StringFromLiteral(",total=")), self.totalCount())
}

func (self *app__TodoApp) totalCount() int {
	return self.store.__hx_this.totalCount()
}

func (self *app__TodoApp) openCount() int {
	return self.store.__hx_this.openCount()
}

func (self *app__TodoApp) doneCount() int {
	return self.store.__hx_this.doneCount()
}

func (self *app__TodoApp) diagnostics() *string {
	if !self.runtime.supportsDiagnostics() {
		return hxrt.StringFromLiteral("off")
	}
	return self.runtime.diagnostics(self.store.__hx_this.list())
}

func (self *app__TodoApp) render() *string {
	out := hxrt.StringFromLiteral("== TODO ==")
	_ = out
	items := self.store.__hx_this.list()
	count := items.length
	_ = count
	i := 0
	for i < count {
		raw := func(hx_value_24 any) *model__TodoItem {
			if hx_value_24 == nil {
				var hx_zero_25 *model__TodoItem
				return hx_zero_25
			}
			return hx_value_24.(*model__TodoItem)
		}(items.pop())
		if hxrt.StringEqualAny(raw, nil) {
			break
		}
		item := raw
		_ = item
		state := hxrt.StringFromLiteral("[ ]")
		if item.done {
			state = hxrt.StringFromLiteral("[x]")
		}
		tags := hxrt.StringFromLiteral("-")
		if item.tags.length != 0 {
			tags = app__TodoApp_joinStringList(item.tags, hxrt.StringFromLiteral(","))
		}
		out = hxrt.StringConcatStringPtr(out, hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatAny(hxrt.StringConcatStringPtr(hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("\n"), state), hxrt.StringFromLiteral(" #")), item.id), hxrt.StringFromLiteral(" p")), item.priority), hxrt.StringFromLiteral(" ")), item.title), hxrt.StringFromLiteral(" tags:")), tags))
		items.add(item)
		i = int(int32((i + 1)))
	}
	out = hxrt.StringConcatStringPtr(out, hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("\nsummary "), self.baselineSignature()))
	return out
}

func (self *app__TodoApp) items() *haxe__ds__List {
	return self.store.__hx_this.list()
}

func app__TodoApp_joinStringList(values *haxe__ds__List, separator *string) *string {
	out := hxrt.StringFromLiteral("")
	_ = out
	first := true
	_ = first
	count := values.length
	_ = count
	i := 0
	for i < count {
		raw := func(hx_value_26 any) *string {
			if hx_value_26 == nil {
				var hx_zero_27 *string
				return hx_zero_27
			}
			return hx_value_26.(*string)
		}(values.pop())
		if hxrt.StringEqualStringPtr(raw, nil) {
			break
		}
		value := raw
		_ = value
		if !first {
			out = hxrt.StringConcatStringPtr(out, separator)
		}
		out = hxrt.StringConcatStringPtr(out, value)
		values.add(value)
		first = false
		i = int(int32((i + 1)))
	}
	return out
}

type I_model__TodoItem interface {
	set_title(value *string) *string
	set_done(value bool) bool
	set_priority(value int) int
}

type model__TodoItem struct {
	__hx_this I_model__TodoItem
	id        int
	title     *string
	done      bool
	priority  int
	tags      *haxe__ds__List
}

func New_model__TodoItem(id int, title *string, priority int) *model__TodoItem {
	self := &model__TodoItem{}
	self.__hx_this = self
	self.id = id
	self.__hx_this.set_title(title)
	self.__hx_this.set_done(false)
	self.__hx_this.set_priority(priority)
	self.tags = New_haxe__ds__List()
	return self
}

func (self *model__TodoItem) set_title(value *string) *string {
	self.title = value
	return value
}

func (self *model__TodoItem) set_done(value bool) bool {
	self.done = value
	return value
}

func (self *model__TodoItem) set_priority(value int) int {
	self.priority = value
	return value
}

type I_model__TodoStore interface {
	add(title *string, priority int) *model__TodoItem
	toggle(id int) bool
	addTag(id int, tag *string) bool
	list() *haxe__ds__List
	totalCount() int
	openCount() int
	doneCount() int
	findById(id int) *model__TodoItem
}

type model__TodoStore struct {
	__hx_this I_model__TodoStore
	nextId    int
	entries   *haxe__ds__List
}

func New_model__TodoStore() *model__TodoStore {
	self := &model__TodoStore{}
	self.__hx_this = self
	self.nextId = 1
	self.entries = New_haxe__ds__List()
	return self
}

func (self *model__TodoStore) add(title *string, priority int) *model__TodoItem {
	item := New_model__TodoItem(self.nextId, title, priority)
	_ = item
	self.nextId = int(int32((self.nextId + 1)))
	self.entries.add(item)
	return item
}

func (self *model__TodoStore) toggle(id int) bool {
	item := self.findById(id)
	if hxrt.StringEqualAny(item, nil) {
		return false
	}
	item.__hx_this.set_done(!item.done)
	return true
}

func (self *model__TodoStore) addTag(id int, tag *string) bool {
	item := self.findById(id)
	if hxrt.StringEqualAny(item, nil) {
		return false
	}
	item.tags.add(tag)
	return true
}

func (self *model__TodoStore) list() *haxe__ds__List {
	return self.entries
}

func (self *model__TodoStore) totalCount() int {
	return self.entries.length
}

func (self *model__TodoStore) openCount() int {
	total := 0
	_ = total
	count := self.entries.length
	_ = count
	i := 0
	for i < count {
		value := func(hx_value_28 any) *model__TodoItem {
			if hx_value_28 == nil {
				var hx_zero_29 *model__TodoItem
				return hx_zero_29
			}
			return hx_value_28.(*model__TodoItem)
		}(self.entries.pop())
		if hxrt.StringEqualAny(value, nil) {
			break
		}
		item := value
		if !item.done {
			total = int(int32((total + 1)))
		}
		self.entries.add(item)
		i = int(int32((i + 1)))
	}
	return total
}

func (self *model__TodoStore) doneCount() int {
	total := 0
	_ = total
	count := self.entries.length
	_ = count
	i := 0
	for i < count {
		value := func(hx_value_30 any) *model__TodoItem {
			if hx_value_30 == nil {
				var hx_zero_31 *model__TodoItem
				return hx_zero_31
			}
			return hx_value_30.(*model__TodoItem)
		}(self.entries.pop())
		if hxrt.StringEqualAny(value, nil) {
			break
		}
		item := value
		if item.done {
			total = int(int32((total + 1)))
		}
		self.entries.add(item)
		i = int(int32((i + 1)))
	}
	return total
}

func (self *model__TodoStore) findById(id int) *model__TodoItem {
	var found *model__TodoItem = nil
	_ = found
	count := self.entries.length
	_ = count
	i := 0
	for i < count {
		value := func(hx_value_32 any) *model__TodoItem {
			if hx_value_32 == nil {
				var hx_zero_33 *model__TodoItem
				return hx_zero_33
			}
			return hx_value_32.(*model__TodoItem)
		}(self.entries.pop())
		if hxrt.StringEqualAny(value, nil) {
			break
		}
		item := value
		if item.id == id {
			found = item
		}
		self.entries.add(item)
		i = int(int32((i + 1)))
	}
	return found
}

type I_profile__GopherRuntime interface {
	profileId() *string
	normalizeTitle(title *string) *string
	normalizeTag(tag *string) *string
	supportsBatchAdd() bool
	supportsDiagnostics() bool
	diagnostics(items *haxe__ds__List) *string
}

type profile__GopherRuntime struct {
	__hx_this I_profile__GopherRuntime
}

func New_profile__GopherRuntime() *profile__GopherRuntime {
	self := &profile__GopherRuntime{}
	self.__hx_this = self
	return self
}

func (self *profile__GopherRuntime) profileId() *string {
	return hxrt.StringFromLiteral("gopher")
}

func (self *profile__GopherRuntime) normalizeTitle(title *string) *string {
	return title
}

func (self *profile__GopherRuntime) normalizeTag(tag *string) *string {
	return hxrt.StringConcatStringPtr(hxrt.StringFromLiteral("go-"), tag)
}

func (self *profile__GopherRuntime) supportsBatchAdd() bool {
	return true
}

func (self *profile__GopherRuntime) supportsDiagnostics() bool {
	return false
}

func (self *profile__GopherRuntime) diagnostics(items *haxe__ds__List) *string {
	return hxrt.StringFromLiteral("off")
}

func profile__RuntimeFactory_create() profile__TodoRuntime {
	return New_profile__GopherRuntime()
}

type profile__TodoRuntime interface {
	profileId() *string
	normalizeTitle(title *string) *string
	normalizeTag(tag *string) *string
	supportsBatchAdd() bool
	supportsDiagnostics() bool
	diagnostics(items *haxe__ds__List) *string
}

type haxe__io__Encoding struct {
}

type haxe__io__Input interface {
	get_bigEndian() bool
	set_bigEndian(e bool) bool
	readByte() int
	readBytes(buf *haxe__io__Bytes, pos int, len int) int
	close()
}

type haxe__io__Output interface {
	get_bigEndian() bool
	set_bigEndian(e bool) bool
	writeByte(c int)
	writeBytes(s *haxe__io__Bytes, pos int, len int) int
	flush()
	close()
}

type haxe__io__Eof struct {
}

type haxe__io__Error struct {
	tag    int
	params []any
}

type haxe__io__Bytes struct {
	b             []int
	length        int
	__hx_raw      []byte
	__hx_rawValid bool
}

type haxe__io__BytesBuffer struct {
	b []int
}

type haxe__io__BytesInput struct {
	bigEndian bool
	b         []int
	pos       int
	len       int
	totlen    int
}

type haxe__io__BytesOutput struct {
	bigEndian bool
	b         *haxe__io__BytesBuffer
}

func New_haxe__io__Input() haxe__io__Input {
	return New_haxe__io__BytesInput(&haxe__io__Bytes{b: []int{}, length: 0})
}

func New_haxe__io__Output() haxe__io__Output {
	return New_haxe__io__BytesOutput()
}

func New_haxe__io__Eof() *haxe__io__Eof {
	return &haxe__io__Eof{}
}

func (self *haxe__io__Eof) toString() *string {
	return hxrt.StringFromLiteral("Eof")
}

var haxe__io__Error_Blocked *haxe__io__Error = &haxe__io__Error{tag: 0}

var haxe__io__Error_Overflow *haxe__io__Error = &haxe__io__Error{tag: 1}

var haxe__io__Error_OutsideBounds *haxe__io__Error = &haxe__io__Error{tag: 2}

func haxe__io__Error_Custom(e any) *haxe__io__Error {
	return &haxe__io__Error{tag: 3, params: []any{e}}
}

func (self *haxe__io__Error) String() string {
	if self == nil {
		return "null"
	}
	switch self.tag {
	case 0:
		return "Blocked"
	case 1:
		return "Overflow"
	case 2:
		return "OutsideBounds"
	case 3:
		if len(self.params) == 0 {
			return "Custom(null)"
		}
		return "Custom(" + *hxrt.StdString(self.params[0]) + ")"
	default:
		return "Error"
	}
}

func (self *haxe__io__Error) toString() *string {
	return hxrt.StringFromLiteral(self.String())
}

func New_haxe__io__Bytes(length int, b []int) *haxe__io__Bytes {
	if b == nil {
		b = make([]int, length)
	}
	return &haxe__io__Bytes{b: b, length: len(b)}
}

func haxe__io__Bytes_alloc(length int) *haxe__io__Bytes {
	return &haxe__io__Bytes{b: make([]int, length), length: length}
}

func haxe__io__Bytes_ofString(value *string, encoding ...*haxe__io__Encoding) *haxe__io__Bytes {
	raw := []byte(*hxrt.StdString(value))
	converted := make([]int, len(raw))
	for i := 0; i < len(raw); i++ {
		converted[i] = int(raw[i])
	}
	return &haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: raw, __hx_rawValid: true}
}

func (self *haxe__io__Bytes) toString() *string {
	if self == nil {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.BytesToString(self.b)
}

func (self *haxe__io__Bytes) get(pos int) int {
	return self.b[pos]
}

func (self *haxe__io__Bytes) set(pos int, value int) {
	self.b[pos] = value
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) blit(pos int, src *haxe__io__Bytes, srcpos int, len int) {
	if self == nil || src == nil || pos < 0 || srcpos < 0 || len < 0 || pos+len > self.length || srcpos+len > src.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return
	}
	if len == 0 {
		return
	}
	if self == src && pos > srcpos {
		for i := len - 1; i >= 0; i-- {
			self.b[pos+i] = src.b[srcpos+i]
		}
	} else {
		for i := 0; i < len; i++ {
			self.b[pos+i] = src.b[srcpos+i]
		}
	}
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) fill(pos int, len int, value int) {
	if self == nil || pos < 0 || len < 0 || pos+len > self.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return
	}
	masked := value & 255
	for i := 0; i < len; i++ {
		self.b[pos+i] = masked
	}
	self.__hx_rawValid = false
}

func (self *haxe__io__Bytes) sub(pos int, len int) *haxe__io__Bytes {
	if self == nil || pos < 0 || len < 0 || pos+len > self.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	if len == 0 {
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	copied := make([]int, len)
	copy(copied, self.b[pos:pos+len])
	return &haxe__io__Bytes{b: copied, length: len}
}

func (self *haxe__io__Bytes) compare(other *haxe__io__Bytes) int {
	if self == nil && other == nil {
		return 0
	}
	if self == nil {
		return -1
	}
	if other == nil {
		return 1
	}
	limit := self.length
	if other.length < limit {
		limit = other.length
	}
	for i := 0; i < limit; i++ {
		if self.b[i] < other.b[i] {
			return -1
		}
		if self.b[i] > other.b[i] {
			return 1
		}
	}
	if self.length < other.length {
		return -1
	}
	if self.length > other.length {
		return 1
	}
	return 0
}

func New_haxe__io__BytesBuffer() *haxe__io__BytesBuffer {
	return &haxe__io__BytesBuffer{b: []int{}}
}

func (self *haxe__io__BytesBuffer) addByte(value int) {
	self.b = append(self.b, (value & 255))
}

func (self *haxe__io__BytesBuffer) add(src *haxe__io__Bytes) {
	if src == nil {
		return
	}
	self.b = append(self.b, src.b...)
}

func (self *haxe__io__BytesBuffer) addBytes(src *haxe__io__Bytes, pos int, len int) {
	if src == nil || pos < 0 || len < 0 || pos+len > src.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return
	}
	if len == 0 {
		return
	}
	self.b = append(self.b, src.b[pos:pos+len]...)
}

func (self *haxe__io__BytesBuffer) addString(value *string, encoding ...*haxe__io__Encoding) {
	self.add(haxe__io__Bytes_ofString(value))
}

func (self *haxe__io__BytesBuffer) getBytes() *haxe__io__Bytes {
	copied := hxrt.BytesClone(self.b)
	return &haxe__io__Bytes{b: copied, length: len(copied)}
}

func (self *haxe__io__BytesBuffer) get_length() int {
	return len(self.b)
}

func New_haxe__io__BytesInput(b *haxe__io__Bytes, opts ...int) *haxe__io__BytesInput {
	if b == nil {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return &haxe__io__BytesInput{}
	}
	start := 0
	if len(opts) > 0 {
		start = opts[0]
	}
	sliceLen := (b.length - start)
	if len(opts) > 1 {
		sliceLen = opts[1]
	}
	if start < 0 || sliceLen < 0 || start+sliceLen > b.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return &haxe__io__BytesInput{}
	}
	return &haxe__io__BytesInput{b: b.b, pos: start, len: sliceLen, totlen: sliceLen}
}

func (self *haxe__io__BytesInput) get_position() int {
	return self.pos
}

func (self *haxe__io__BytesInput) set_position(p int) int {
	if p < 0 {
		p = 0
	} else {
		if p > self.totlen {
			p = self.totlen
		}
	}
	self.len = (self.totlen - p)
	self.pos = p
	return p
}

func (self *haxe__io__BytesInput) get_length() int {
	return self.totlen
}

func (self *haxe__io__BytesInput) readByte() int {
	if self == nil || self.len == 0 {
		hxrt.Throw(&haxe__io__Eof{})
		return 0
	}
	self.len = (self.len - 1)
	value := self.b[self.pos]
	self.pos = (self.pos + 1)
	return value
}

func (self *haxe__io__BytesInput) readBytes(buf *haxe__io__Bytes, pos int, len int) int {
	if buf == nil || pos < 0 || len < 0 || pos+len > buf.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return 0
	}
	if len > 0 && (self == nil || self.len == 0) {
		hxrt.Throw(&haxe__io__Eof{})
		return 0
	}
	if self == nil {
		return 0
	}
	if self.len < len {
		len = self.len
	}
	for i := 0; i < len; i++ {
		buf.b[pos+i] = self.b[self.pos+i]
	}
	self.pos += len
	self.len -= len
	return len
}

func (self *haxe__io__BytesInput) get_bigEndian() bool {
	if self == nil {
		return false
	}
	return self.bigEndian
}

func (self *haxe__io__BytesInput) set_bigEndian(e bool) bool {
	if self != nil {
		self.bigEndian = e
	}
	return e
}

func (self *haxe__io__BytesInput) close() {
	_ = self
}

func New_haxe__io__BytesOutput() *haxe__io__BytesOutput {
	return &haxe__io__BytesOutput{b: &haxe__io__BytesBuffer{b: []int{}}}
}

func (self *haxe__io__BytesOutput) get_length() int {
	if self == nil || self.b == nil {
		return 0
	}
	return self.b.get_length()
}

func (self *haxe__io__BytesOutput) writeByte(c int) {
	if self == nil || self.b == nil {
		return
	}
	self.b.addByte(c)
}

func (self *haxe__io__BytesOutput) writeBytes(buf *haxe__io__Bytes, pos int, len int) int {
	if buf == nil || pos < 0 || len < 0 || pos+len > buf.length {
		hxrt.Throw(haxe__io__Error_OutsideBounds)
		return 0
	}
	if self == nil || self.b == nil {
		return 0
	}
	self.b.addBytes(buf, pos, len)
	return len
}

func (self *haxe__io__BytesOutput) get_bigEndian() bool {
	if self == nil {
		return false
	}
	return self.bigEndian
}

func (self *haxe__io__BytesOutput) set_bigEndian(e bool) bool {
	if self != nil {
		self.bigEndian = e
	}
	return e
}

func (self *haxe__io__BytesOutput) flush() {
	_ = self
}

func (self *haxe__io__BytesOutput) close() {
	_ = self
}

func (self *haxe__io__BytesOutput) getBytes() *haxe__io__Bytes {
	if self == nil || self.b == nil {
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	return self.b.getBytes()
}

type haxe__ds__IntMap struct {
	h map[int]any
}

type haxe__ds__StringMap struct {
	h map[string]any
}

type haxe__ds__ObjectMap struct {
	h map[any]any
}

type haxe__ds__EnumValueMap struct {
	h map[any]any
}

type haxe__ds__List struct {
	items  []any
	length int
}

func New_haxe__ds__IntMap() *haxe__ds__IntMap {
	return &haxe__ds__IntMap{h: map[int]any{}}
}

func (self *haxe__ds__IntMap) set(key int, value any) {
	self.h[key] = value
}

func (self *haxe__ds__IntMap) get(key int) any {
	value := self.h[key]
	return value
}

func (self *haxe__ds__IntMap) exists(key int) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__IntMap) remove(key int) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__StringMap() *haxe__ds__StringMap {
	return &haxe__ds__StringMap{h: map[string]any{}}
}

func (self *haxe__ds__StringMap) set(key *string, value any) {
	self.h[*hxrt.StdString(key)] = value
}

func (self *haxe__ds__StringMap) get(key *string) any {
	value := self.h[*hxrt.StdString(key)]
	return value
}

func (self *haxe__ds__StringMap) exists(key *string) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	return ok
}

func (self *haxe__ds__StringMap) remove(key *string) bool {
	_, ok := self.h[*hxrt.StdString(key)]
	delete(self.h, *hxrt.StdString(key))
	return ok
}

func New_haxe__ds__ObjectMap() *haxe__ds__ObjectMap {
	return &haxe__ds__ObjectMap{h: map[any]any{}}
}

func (self *haxe__ds__ObjectMap) set(key any, value any) {
	self.h[key] = value
}

func (self *haxe__ds__ObjectMap) get(key any) any {
	return self.h[key]
}

func (self *haxe__ds__ObjectMap) exists(key any) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__ObjectMap) remove(key any) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__EnumValueMap() *haxe__ds__EnumValueMap {
	return &haxe__ds__EnumValueMap{h: map[any]any{}}
}

func (self *haxe__ds__EnumValueMap) set(key any, value any) {
	self.h[key] = value
}

func (self *haxe__ds__EnumValueMap) get(key any) any {
	return self.h[key]
}

func (self *haxe__ds__EnumValueMap) exists(key any) bool {
	_, ok := self.h[key]
	return ok
}

func (self *haxe__ds__EnumValueMap) remove(key any) bool {
	_, ok := self.h[key]
	delete(self.h, key)
	return ok
}

func New_haxe__ds__List() *haxe__ds__List {
	return &haxe__ds__List{items: []any{}, length: 0}
}

func (self *haxe__ds__List) add(item any) {
	self.items = append(self.items, item)
	self.length = len(self.items)
}

func (self *haxe__ds__List) push(item any) {
	self.items = append([]any{item}, self.items...)
	self.length = len(self.items)
}

func (self *haxe__ds__List) pop() any {
	if len(self.items) == 0 {
		return nil
	}
	head := self.items[0]
	self.items = self.items[1:]
	self.length = len(self.items)
	return head
}

func (self *haxe__ds__List) first() any {
	if len(self.items) == 0 {
		return nil
	}
	return self.items[0]
}

func (self *haxe__ds__List) last() any {
	size := len(self.items)
	if size == 0 {
		return nil
	}
	return self.items[(size - 1)]
}

type Sys struct {
}

type sys__io__File struct {
}

type sys__io__ProcessOutput struct {
	impl *hxrt.ProcessOutput
}

type sys__io__Process struct {
	impl   *hxrt.Process
	stdout *sys__io__ProcessOutput
}

func Sys_getCwd() *string {
	return hxrt.SysGetCwd()
}

func Sys_args() []*string {
	return hxrt.SysArgs()
}

func sys__io__File_saveContent(path *string, content *string) {
	hxrt.FileSaveContent(path, content)
}

func sys__io__File_getContent(path *string) *string {
	return hxrt.FileGetContent(path)
}

func New_sys__io__Process(command *string, args []*string) *sys__io__Process {
	impl := hxrt.NewProcess(command, args)
	stdout := &sys__io__ProcessOutput{}
	if impl != nil {
		stdout.impl = impl.Stdout()
	}
	return &sys__io__Process{impl: impl, stdout: stdout}
}

func (self *sys__io__ProcessOutput) readLine() *string {
	if self == nil || self.impl == nil {
		return hxrt.StringFromLiteral("")
	}
	return self.impl.ReadLine()
}

func (self *sys__io__Process) close() {
	if self == nil || self.impl == nil {
		return
	}
	self.impl.Close()
}

type Std struct {
}

type StringTools struct {
}

func StringTools_trim(value *string) *string {
	return hxrt.StringFromLiteral(strings.TrimSpace(*hxrt.StdString(value)))
}

func StringTools_startsWith(value *string, prefix *string) bool {
	return strings.HasPrefix(*hxrt.StdString(value), *hxrt.StdString(prefix))
}

func StringTools_replace(value *string, sub *string, by *string) *string {
	return hxrt.StringFromLiteral(strings.ReplaceAll(*hxrt.StdString(value), *hxrt.StdString(sub), *hxrt.StdString(by)))
}

type Date struct {
	value time.Time
}

func Date_fromString(source *string) *Date {
	raw := *hxrt.StdString(source)
	parsed, err := time.ParseInLocation("2006-01-02 15:04:05", raw, time.Local)
	if err != nil {
		parsedDateOnly, errDateOnly := time.ParseInLocation("2006-01-02", raw, time.Local)
		if errDateOnly == nil {
			parsed = parsedDateOnly
		} else {
			parsed = time.Unix(0, 0)
		}
	}
	return &Date{value: parsed}
}

func Date_now() *Date {
	return &Date{value: time.Now()}
}

func (self *Date) getFullYear() int {
	return self.value.Year()
}

func (self *Date) getMonth() int {
	return int(self.value.Month()) - 1
}

func (self *Date) getDate() int {
	return self.value.Day()
}

func (self *Date) getHours() int {
	return self.value.Hour()
}

type Math struct {
}

func Math_floor(value float64) int {
	return int(math.Floor(value))
}

func Math_ceil(value float64) int {
	return int(math.Ceil(value))
}

func Math_round(value float64) int {
	return int(math.Floor(value + 0.5))
}

func Math_abs(value float64) float64 {
	return math.Abs(value)
}

func Math_isNaN(value float64) bool {
	return math.IsNaN(value)
}

func Math_isFinite(value float64) bool {
	return !math.IsInf(value, 0)
}

func Math_min(a float64, b float64) float64 {
	return math.Min(a, b)
}

func Math_max(a float64, b float64) float64 {
	return math.Max(a, b)
}

type Type struct {
}

type Reflect struct {
}

func Reflect_compare(a any, b any) int {
	toFloat := func(value any) (float64, bool) {
		switch v := value.(type) {
		case int:
			return float64(v), true
		case int8:
			return float64(v), true
		case int16:
			return float64(v), true
		case int32:
			return float64(v), true
		case int64:
			return float64(v), true
		case uint:
			return float64(v), true
		case uint8:
			return float64(v), true
		case uint16:
			return float64(v), true
		case uint32:
			return float64(v), true
		case uint64:
			return float64(v), true
		case float32:
			return float64(v), true
		case float64:
			return v, true
		default:
			return 0, false
		}
	}
	if af, ok := toFloat(a); ok {
		if bf, okB := toFloat(b); okB {
			if af < bf {
				return -1
			}
			if af > bf {
				return 1
			}
			return 0
		}
	}
	aStr := *hxrt.StdString(a)
	bStr := *hxrt.StdString(b)
	if aStr < bStr {
		return -1
	}
	if aStr > bStr {
		return 1
	}
	return 0
}

func Reflect_field(obj any, field *string) any {
	if obj == nil {
		return nil
	}
	key := *hxrt.StdString(field)
	switch value := obj.(type) {
	case map[string]any:
		return value[key]
	case map[any]any:
		return value[key]
	case *map[string]any:
		if value == nil {
			return nil
		}
		return (*value)[key]
	case *map[any]any:
		if value == nil {
			return nil
		}
		return (*value)[key]
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return nil
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return nil
		}
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		if fieldValue := rv.FieldByName(key); fieldValue.IsValid() && fieldValue.CanInterface() {
			return fieldValue.Interface()
		}
	}
	method := reflect.ValueOf(obj).MethodByName(key)
	if method.IsValid() {
		return method.Interface()
	}
	return nil
}

func Reflect_hasField(obj any, field *string) bool {
	if obj == nil {
		return false
	}
	key := *hxrt.StdString(field)
	switch value := obj.(type) {
	case map[string]any:
		_, ok := value[key]
		return ok
	case map[any]any:
		_, ok := value[key]
		return ok
	case *map[string]any:
		if value == nil {
			return false
		}
		_, ok := (*value)[key]
		return ok
	case *map[any]any:
		if value == nil {
			return false
		}
		_, ok := (*value)[key]
		return ok
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() {
		return false
	}
	if rv.Kind() == reflect.Pointer {
		if rv.IsNil() {
			return false
		}
		rv = rv.Elem()
	}
	if rv.Kind() == reflect.Struct {
		if rv.FieldByName(key).IsValid() {
			return true
		}
	}
	return reflect.ValueOf(obj).MethodByName(key).IsValid()
}

func Reflect_setField(obj any, field *string, value any) {
	if obj == nil {
		hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		return
	}
	key := *hxrt.StdString(field)
	switch target := obj.(type) {
	case map[string]any:
		target[key] = value
		return
	case map[any]any:
		target[key] = value
		return
	case *map[string]any:
		if target == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			return
		}
		(*target)[key] = value
		return
	case *map[any]any:
		if target == nil {
			hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
			return
		}
		(*target)[key] = value
		return
	}
	rv := reflect.ValueOf(obj)
	if !rv.IsValid() || rv.Kind() != reflect.Pointer {
		return
	}
	if rv.IsNil() {
		hxrt.Throw(hxrt.StringFromLiteral("Null Access"))
		return
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return
	}
	fieldValue := rv.FieldByName(key)
	if !fieldValue.IsValid() || !fieldValue.CanSet() {
		return
	}
	if value == nil {
		fieldValue.Set(reflect.Zero(fieldValue.Type()))
		return
	}
	incoming := reflect.ValueOf(value)
	if incoming.Type().AssignableTo(fieldValue.Type()) {
		fieldValue.Set(incoming)
		return
	}
	if incoming.Type().ConvertibleTo(fieldValue.Type()) {
		fieldValue.Set(incoming.Convert(fieldValue.Type()))
		return
	}
	if fieldValue.Kind() == reflect.Interface {
		fieldValue.Set(incoming)
	}
}

type Xml struct {
	raw *string
}

func Xml_parse(source *string) *Xml {
	return haxe__xml__Parser_parse(source)
}

func (self *Xml) toString() *string {
	if self == nil || self.raw == nil {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.StringFromLiteral(*self.raw)
}

type haxe__crypto__Base64 struct {
}

type haxe__crypto__Md5 struct {
}

type haxe__crypto__Sha1 struct {
}

type haxe__crypto__Sha224 struct {
}

type haxe__crypto__Sha256 struct {
}

func hxrt_haxeBytesToRaw(value *haxe__io__Bytes) []byte {
	if value == nil {
		return []byte{}
	}
	if value.__hx_rawValid && len(value.__hx_raw) == len(value.b) {
		return value.__hx_raw
	}
	raw := make([]byte, len(value.b))
	for i := 0; i < len(value.b); i++ {
		raw[i] = byte(value.b[i])
	}
	value.__hx_raw = raw
	value.__hx_rawValid = true
	return raw
}

func hxrt_rawToHaxeBytes(value []byte) *haxe__io__Bytes {
	converted := make([]int, len(value))
	for i := 0; i < len(value); i++ {
		converted[i] = int(value[i])
	}
	return &haxe__io__Bytes{b: converted, length: len(converted), __hx_raw: value, __hx_rawValid: true}
}

func haxe__crypto__Base64_encode(bytes *haxe__io__Bytes, complement ...bool) *string {
	useComplement := true
	if len(complement) > 0 {
		useComplement = complement[0]
	}
	encoded := base64.StdEncoding.EncodeToString(hxrt_haxeBytesToRaw(bytes))
	if !useComplement {
		encoded = strings.TrimRight(encoded, "=")
	}
	return hxrt.StringFromLiteral(encoded)
}

func haxe__crypto__Base64_decode(value *string, complement ...bool) *haxe__io__Bytes {
	useComplement := true
	if len(complement) > 0 {
		useComplement = complement[0]
	}
	rawValue := *hxrt.StdString(value)
	if useComplement {
		rawValue = strings.TrimRight(rawValue, "=")
	}
	decoded, err := base64.RawStdEncoding.DecodeString(rawValue)
	if err != nil {
		decoded, err = base64.StdEncoding.DecodeString(*hxrt.StdString(value))
		if err != nil {
			hxrt.Throw(err)
			return &haxe__io__Bytes{b: []int{}, length: 0}
		}
	}
	return hxrt_rawToHaxeBytes(decoded)
}

func haxe__crypto__Base64_urlEncode(bytes *haxe__io__Bytes, complement ...bool) *string {
	useComplement := false
	if len(complement) > 0 {
		useComplement = complement[0]
	}
	encoded := base64.RawURLEncoding.EncodeToString(hxrt_haxeBytesToRaw(bytes))
	if useComplement {
		missing := len(encoded) % 4
		if missing != 0 {
			encoded = (encoded + strings.Repeat("=", (4-missing)))
		}
	}
	return hxrt.StringFromLiteral(encoded)
}

func haxe__crypto__Base64_urlDecode(value *string, complement ...bool) *haxe__io__Bytes {
	rawValue := *hxrt.StdString(value)
	decoded, err := base64.RawURLEncoding.DecodeString(strings.TrimRight(rawValue, "="))
	if err != nil {
		hxrt.Throw(err)
		return &haxe__io__Bytes{b: []int{}, length: 0}
	}
	return hxrt_rawToHaxeBytes(decoded)
}

func haxe__crypto__Md5_encode(value *string) *string {
	sum := md5.Sum([]byte(*hxrt.StdString(value)))
	return hxrt.StringFromLiteral(hex.EncodeToString(sum[:]))
}

func haxe__crypto__Md5_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	sum := md5.Sum(hxrt_haxeBytesToRaw(value))
	return hxrt_rawToHaxeBytes(sum[:])
}

func haxe__crypto__Sha1_encode(value *string) *string {
	sum := sha1.Sum([]byte(*hxrt.StdString(value)))
	return hxrt.StringFromLiteral(hex.EncodeToString(sum[:]))
}

func haxe__crypto__Sha1_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	sum := sha1.Sum(hxrt_haxeBytesToRaw(value))
	return hxrt_rawToHaxeBytes(sum[:])
}

func haxe__crypto__Sha224_encode(value *string) *string {
	sum := sha256.Sum224([]byte(*hxrt.StdString(value)))
	return hxrt.StringFromLiteral(hex.EncodeToString(sum[:]))
}

func haxe__crypto__Sha224_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	sum := sha256.Sum224(hxrt_haxeBytesToRaw(value))
	return hxrt_rawToHaxeBytes(sum[:])
}

func haxe__crypto__Sha256_encode(value *string) *string {
	sum := sha256.Sum256([]byte(*hxrt.StdString(value)))
	return hxrt.StringFromLiteral(hex.EncodeToString(sum[:]))
}

func haxe__crypto__Sha256_make(value *haxe__io__Bytes) *haxe__io__Bytes {
	sum := sha256.Sum256(hxrt_haxeBytesToRaw(value))
	return hxrt_rawToHaxeBytes(sum[:])
}

type haxe__ds__BalancedTree struct {
}

type haxe__ds__Option struct {
	tag    int
	params []any
}

var haxe__ds__Option_None *haxe__ds__Option = &haxe__ds__Option{tag: 1, params: []any{}}

func haxe__ds__Option_Some(value any) *haxe__ds__Option {
	return &haxe__ds__Option{tag: 0, params: []any{value}}
}

type haxe__io__Path struct {
	dir       *string
	file      *string
	ext       *string
	backslash bool
}

func New_haxe__io__Path(path *string) *haxe__io__Path {
	raw := *hxrt.StdString(path)
	dir := filepath.Dir(raw)
	if dir == "." {
		dir = ""
	}
	base := filepath.Base(raw)
	dotExt := filepath.Ext(base)
	file := base
	if dotExt != "" {
		file = strings.TrimSuffix(base, dotExt)
	}
	ext := strings.TrimPrefix(dotExt, ".")
	return &haxe__io__Path{dir: hxrt.StringFromLiteral(dir), file: hxrt.StringFromLiteral(file), ext: hxrt.StringFromLiteral(ext), backslash: strings.Contains(raw, "\\")}
}

func haxe__io__Path_join(parts []*string) *string {
	if len(parts) == 0 {
		return hxrt.StringFromLiteral("")
	}
	joined := filepath.ToSlash(filepath.Join(hxrt.StringSlice(parts)...))
	return hxrt.StringFromLiteral(joined)
}

type haxe__io__StringInput struct {
}

type haxe__xml__Parser struct {
}

type haxe__xml__Printer struct {
}

func haxe__xml__Parser_parse(source *string, strict ...bool) *Xml {
	raw := *hxrt.StdString(source)
	decoder := xml.NewDecoder(strings.NewReader(raw))
	for {
		_, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			hxrt.Throw(err)
			return &Xml{raw: hxrt.StringFromLiteral("")}
		}
	}
	return &Xml{raw: hxrt.StringFromLiteral(raw)}
}

func haxe__xml__Printer_print(value *Xml, pretty ...bool) *string {
	if value == nil || value.raw == nil {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.StringFromLiteral(*value.raw)
}

type haxe__zip__Compress struct {
}

type haxe__zip__Uncompress struct {
}

func haxe__zip__Compress_run(src *haxe__io__Bytes, level int) *haxe__io__Bytes {
	raw := hxrt_haxeBytesToRaw(src)
	var buffer bytes.Buffer
	writer, err := zlib.NewWriterLevel(&buffer, level)
	if err != nil {
		hxrt.Throw(err)
		return nil
	}
	if _, err := writer.Write(raw); err != nil {
		_ = writer.Close()
		hxrt.Throw(err)
		return nil
	}
	if err := writer.Close(); err != nil {
		hxrt.Throw(err)
		return nil
	}
	return hxrt_rawToHaxeBytes(buffer.Bytes())
}

func haxe__zip__Uncompress_run(src *haxe__io__Bytes, bufsize ...int) *haxe__io__Bytes {
	raw := hxrt_haxeBytesToRaw(src)
	reader, err := zlib.NewReader(bytes.NewReader(raw))
	if err != nil {
		hxrt.Throw(err)
		return nil
	}
	defer reader.Close()
	decoded, err := io.ReadAll(reader)
	if err != nil {
		hxrt.Throw(err)
		return nil
	}
	return hxrt_rawToHaxeBytes(decoded)
}

type sys__FileSystem struct {
}
