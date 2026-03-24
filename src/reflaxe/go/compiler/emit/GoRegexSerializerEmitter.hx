package reflaxe.go.compiler.emit;

#if macro
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.compiler.GoStdlibOwnership;

/**
	What:
	Builds the compiler-owned declarations behind `EReg`, serializer, and
	unserializer support.

	Why:
	This surface is representation-sensitive and metadata-driven, so it remains
	compiler-owned. Extracting it from `GoCompiler` keeps ownership explicit
	without growing the monolith further.

	How:
	Consumes precomputed serializer metadata plus the existing string-literal
	helper from `GoCompiler`, then emits the same declaration set that previously
	lived inline inside `lowerRegexSerializerShimDecls()`.
**/
class GoRegexSerializerEmitter {
	public static function emit(classMetadata:Array<{goTypeName:String, haxeTypeName:String}>,
			enumMetadata:Array<{goTypeName:String, haxeTypeName:String, constructors:Array<String>}>, goRawQuotedString:String->String):Array<GoDecl> {
		var classLookupBody = [GoStmt.GoRaw("switch typeName {")];
		for (entry in classMetadata) {
			classLookupBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.goTypeName) + ":"));
			classLookupBody.push(GoStmt.GoRaw("\treturn " + goRawQuotedString(entry.haxeTypeName) + ", true"));
		}
		classLookupBody.push(GoStmt.GoRaw("default:"));
		classLookupBody.push(GoStmt.GoRaw("\treturn \"\", false"));
		classLookupBody.push(GoStmt.GoRaw("}"));

		var enumLookupBody = [GoStmt.GoRaw("switch typeName {")];
		for (entry in enumMetadata) {
			var constructorLiterals = [for (constructor in entry.constructors) goRawQuotedString(constructor)].join(", ");
			enumLookupBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.goTypeName) + ":"));
			enumLookupBody.push(GoStmt.GoRaw("\tconstructors := []string{" + constructorLiterals + "}"));
			enumLookupBody.push(GoStmt.GoRaw("\tif tag < 0 || tag >= len(constructors) {"));
			enumLookupBody.push(GoStmt.GoRaw("\t\treturn \"\", \"\", false"));
			enumLookupBody.push(GoStmt.GoRaw("\t}"));
			enumLookupBody.push(GoStmt.GoRaw("\treturn " + goRawQuotedString(entry.haxeTypeName) + ", constructors[tag], true"));
		}
		enumLookupBody.push(GoStmt.GoRaw("default:"));
		enumLookupBody.push(GoStmt.GoRaw("\treturn \"\", \"\", false"));
		enumLookupBody.push(GoStmt.GoRaw("}"));

		var enumLookupByNameBody = [GoStmt.GoRaw("switch enumName {")];
		for (entry in enumMetadata) {
			var constructorLiterals = [for (constructor in entry.constructors) goRawQuotedString(constructor)].join(", ");
			enumLookupByNameBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			enumLookupByNameBody.push(GoStmt.GoRaw("\tconstructors := []string{" + constructorLiterals + "}"));
			enumLookupByNameBody.push(GoStmt.GoRaw("\tif index < 0 || index >= len(constructors) {"));
			enumLookupByNameBody.push(GoStmt.GoRaw("\t\treturn \"\", false"));
			enumLookupByNameBody.push(GoStmt.GoRaw("\t}"));
			enumLookupByNameBody.push(GoStmt.GoRaw("\treturn constructors[index], true"));
		}
		enumLookupByNameBody.push(GoStmt.GoRaw("default:"));
		enumLookupByNameBody.push(GoStmt.GoRaw("\treturn \"\", false"));
		enumLookupByNameBody.push(GoStmt.GoRaw("}"));

		var enumLookupIndexBody = [GoStmt.GoRaw("switch enumName {")];
		for (entry in enumMetadata) {
			enumLookupIndexBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			for (index in 0...entry.constructors.length) {
				enumLookupIndexBody.push(GoStmt.GoRaw("\tif constructorName == " + goRawQuotedString(entry.constructors[index]) + " {"));
				enumLookupIndexBody.push(GoStmt.GoRaw("\t\treturn " + index + ", true"));
				enumLookupIndexBody.push(GoStmt.GoRaw("\t}"));
			}
			enumLookupIndexBody.push(GoStmt.GoRaw("\treturn 0, false"));
		}
		enumLookupIndexBody.push(GoStmt.GoRaw("default:"));
		enumLookupIndexBody.push(GoStmt.GoRaw("\treturn 0, false"));
		enumLookupIndexBody.push(GoStmt.GoRaw("}"));

		var classExistsBody = [GoStmt.GoRaw("switch className {")];
		for (entry in classMetadata) {
			classExistsBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			classExistsBody.push(GoStmt.GoRaw("\treturn true"));
		}
		classExistsBody.push(GoStmt.GoRaw("default:"));
		classExistsBody.push(GoStmt.GoRaw("\treturn false"));
		classExistsBody.push(GoStmt.GoRaw("}"));

		var classCreateBody = [GoStmt.GoRaw("switch className {")];
		for (entry in classMetadata) {
			if (!GoStdlibOwnership.canConstructEmptyTypeValue(entry.goTypeName)) {
				continue;
			}
			classCreateBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			classCreateBody.push(GoStmt.GoRaw("\tinstance := &" + entry.goTypeName + "{}"));
			classCreateBody.push(GoStmt.GoRaw("\thxrt_unserializerBindSelf(instance)"));
			classCreateBody.push(GoStmt.GoRaw("\treturn instance, true"));
		}
		classCreateBody.push(GoStmt.GoRaw("default:"));
		classCreateBody.push(GoStmt.GoRaw("\treturn nil, false"));
		classCreateBody.push(GoStmt.GoRaw("}"));

		var enumExistsBody = [GoStmt.GoRaw("switch enumName {")];
		for (entry in enumMetadata) {
			enumExistsBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			enumExistsBody.push(GoStmt.GoRaw("\treturn true"));
		}
		enumExistsBody.push(GoStmt.GoRaw("default:"));
		enumExistsBody.push(GoStmt.GoRaw("\treturn false"));
		enumExistsBody.push(GoStmt.GoRaw("}"));

		var enumCreateBody = [GoStmt.GoRaw("switch enumName {")];
		for (entry in enumMetadata) {
			enumCreateBody.push(GoStmt.GoRaw("case " + goRawQuotedString(entry.haxeTypeName) + ":"));
			enumCreateBody.push(GoStmt.GoRaw("\tenumValue := &" + entry.goTypeName + "{tag: constructorIndex}"));
			enumCreateBody.push(GoStmt.GoRaw("\tif len(args) > 0 {"));
			enumCreateBody.push(GoStmt.GoRaw("\t\tcopied := make([]any, len(args))"));
			enumCreateBody.push(GoStmt.GoRaw("\t\tcopy(copied, args)"));
			enumCreateBody.push(GoStmt.GoRaw("\t\tenumValue.params = copied"));
			enumCreateBody.push(GoStmt.GoRaw("\t}"));
			enumCreateBody.push(GoStmt.GoRaw("\treturn enumValue, true"));
		}
		enumCreateBody.push(GoStmt.GoRaw("default:"));
		enumCreateBody.push(GoStmt.GoRaw("\treturn nil, false"));
		enumCreateBody.push(GoStmt.GoRaw("}"));

		return [
			GoDecl.GoStructDecl("EReg", [
				{name: "regex", typeName: "*regexp.Regexp"},
				{name: "global", typeName: "bool"},
				{name: "lastSource", typeName: "*string"},
				{name: "lastIndices", typeName: "[]int"}
			]),
			GoDecl.GoFuncDecl("New_EReg", null, [
				{
					name: "pattern",
					typeName: "*string"
				},
				{name: "options", typeName: "*string"}
			], ["*EReg"], [
				GoStmt.GoRaw("rawPattern := *hxrt.StdString(pattern)"),
				GoStmt.GoRaw("rawOptions := *hxrt.StdString(options)"),
				GoStmt.GoRaw("global := false"),
				GoStmt.GoRaw("flagI := false"),
				GoStmt.GoRaw("flagM := false"),
				GoStmt.GoRaw("flagS := false"),
				GoStmt.GoRaw("for _, option := range rawOptions {"),
				GoStmt.GoRaw("\tswitch option {"),
				GoStmt.GoRaw("\tcase 'g':"),
				GoStmt.GoRaw("\t\tglobal = true"),
				GoStmt.GoRaw("\tcase 'i':"),
				GoStmt.GoRaw("\t\tflagI = true"),
				GoStmt.GoRaw("\tcase 'm':"),
				GoStmt.GoRaw("\t\tflagM = true"),
				GoStmt.GoRaw("\tcase 's':"),
				GoStmt.GoRaw("\t\tflagS = true"),
				GoStmt.GoRaw("\tcase 'u':"),
				GoStmt.GoRaw("\t\t// RE2 is UTF-8 aware by default; keep parity by accepting and ignoring this option."),
				GoStmt.GoRaw("\tdefault:"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Unsupported regexp option '\" + string(option) + \"'\"))"),
				GoStmt.GoRaw("\t\treturn &EReg{regex: regexp.MustCompile(\"a^\"), global: false, lastSource: nil, lastIndices: nil}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("inlineFlags := \"\""),
				GoStmt.GoRaw("if flagI {"),
				GoStmt.GoRaw("\tinlineFlags += \"i\""),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if flagM {"),
				GoStmt.GoRaw("\tinlineFlags += \"m\""),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if flagS {"),
				GoStmt.GoRaw("\tinlineFlags += \"s\""),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if inlineFlags != \"\" {"),
				GoStmt.GoRaw("\trawPattern = \"(?\" + inlineFlags + \")\" + rawPattern"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("compiled, err := regexp.Compile(rawPattern)"),
				GoStmt.GoIf(GoExpr.GoBinary("!=", GoExpr.GoIdent("err"), GoExpr.GoNil), [
					GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [GoExpr.GoIdent("err")])),
					GoStmt.GoAssign(GoExpr.GoIdent("compiled"),
						GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("regexp"), "MustCompile"), [GoExpr.GoStringLiteral("a^")]))
				],
					null),
				GoStmt.GoReturn(GoExpr.GoRaw("&EReg{regex: compiled, global: global, lastSource: nil, lastIndices: nil}"))
			]),
			GoDecl.GoFuncDecl("hxrt_eregHasMatch", null, [
				{
					name: "self",
					typeName: "*EReg"
				}
			], ["bool"], [
				GoStmt.GoRaw("return self != nil && self.lastSource != nil && len(self.lastIndices) >= 2 && self.lastIndices[0] >= 0 && self.lastIndices[1] >= self.lastIndices[0]")
			]),
			GoDecl.GoFuncDecl("hxrt_eregThrowNoMatch", null, [], [], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [
					GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("Invalid regex operation because no match was made")])
				]))
			]),
			GoDecl.GoFuncDecl("hxrt_eregThrowInvalidGroup", null, [], [], [
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt.Throw"), [
					GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("Invalid group")])
				]))
			]),
			GoDecl.GoFuncDecl("match", {
				name: "self",
				typeName: "*EReg"
			}, [{name: "source", typeName: "*string"}], ["bool"], [
				GoStmt.GoRaw("if self == nil || self.regex == nil {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("raw := *hxrt.StdString(source)"),
				GoStmt.GoRaw("found := self.regex.FindStringSubmatchIndex(raw)"),
				GoStmt.GoRaw("if found == nil {"),
				GoStmt.GoRaw("\tself.lastSource = nil"),
				GoStmt.GoRaw("\tself.lastIndices = nil"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("indices := make([]int, len(found))"),
				GoStmt.GoRaw("copy(indices, found)"),
				GoStmt.GoRaw("self.lastSource = hxrt.StringFromLiteral(raw)"),
				GoStmt.GoRaw("self.lastIndices = indices"),
				GoStmt.GoReturn(GoExpr.GoBoolLiteral(true))
			]),
			GoDecl.GoFuncDecl("matchSub", {
				name: "self",
				typeName: "*EReg"
			}, [
				{
					name: "source",
					typeName: "*string"
				},
				{name: "pos", typeName: "int"}
			], ["bool"], [
				GoStmt.GoRaw("if self == nil || self.regex == nil {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("raw := *hxrt.StdString(source)"),
				GoStmt.GoRaw("if pos < 0 {"),
				GoStmt.GoRaw("\tpos = 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if pos > len(raw) {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("found := self.regex.FindStringSubmatchIndex(raw[pos:])"),
				GoStmt.GoRaw("if found == nil {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("shifted := make([]int, len(found))"),
				GoStmt.GoRaw("for i := 0; i < len(found); i++ {"),
				GoStmt.GoRaw("\tif found[i] >= 0 {"),
				GoStmt.GoRaw("\t\tshifted[i] = found[i] + pos"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\tshifted[i] = -1"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.lastSource = hxrt.StringFromLiteral(raw)"),
				GoStmt.GoRaw("self.lastIndices = shifted"),
				GoStmt.GoReturn(GoExpr.GoBoolLiteral(true))
			]),
			GoDecl.GoFuncDecl("matched", {
				name: "self",
				typeName: "*EReg"
			}, [{name: "index", typeName: "int"}], ["*string"], [
				GoStmt.GoRaw("if !hxrt_eregHasMatch(self) {"),
				GoStmt.GoRaw("\thxrt_eregThrowNoMatch()"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if index < 0 {"),
				GoStmt.GoRaw("\thxrt_eregThrowInvalidGroup()"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("offset := index * 2"),
				GoStmt.GoRaw("if offset+1 >= len(self.lastIndices) {"),
				GoStmt.GoRaw("\thxrt_eregThrowInvalidGroup()"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("start := self.lastIndices[offset]"),
				GoStmt.GoRaw("end := self.lastIndices[offset+1]"),
				GoStmt.GoRaw("if start < 0 || end < start {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("raw := *hxrt.StdString(self.lastSource)"),
				GoStmt.GoRaw("if end > len(raw) {"),
				GoStmt.GoRaw("\tend = len(raw)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("raw[start:end]")]))
			]),
			GoDecl.GoFuncDecl("matchedPos", {
				name: "self",
				typeName: "*EReg"
			}, [], ["map[string]any"], [
				GoStmt.GoRaw("if !hxrt_eregHasMatch(self) {"),
				GoStmt.GoRaw("\thxrt_eregThrowNoMatch()"),
				GoStmt.GoRaw("\treturn map[string]any{\"pos\": 0, \"len\": 0}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("start := self.lastIndices[0]"),
				GoStmt.GoRaw("end := self.lastIndices[1]"),
				GoStmt.GoReturn(GoExpr.GoRaw("map[string]any{\"pos\": start, \"len\": end - start}"))
			]),
			GoDecl.GoFuncDecl("matchedLeft", {
				name: "self",
				typeName: "*EReg"
			}, [], ["*string"], [
				GoStmt.GoRaw("if !hxrt_eregHasMatch(self) {"),
				GoStmt.GoRaw("\thxrt_eregThrowNoMatch()"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("raw := *hxrt.StdString(self.lastSource)"),
				GoStmt.GoRaw("start := self.lastIndices[0]"),
				GoStmt.GoRaw("if start > len(raw) {"),
				GoStmt.GoRaw("\tstart = len(raw)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("raw[:start]")]))
			]),
			GoDecl.GoFuncDecl("matchedRight", {
				name: "self",
				typeName: "*EReg"
			}, [], ["*string"], [
				GoStmt.GoRaw("if !hxrt_eregHasMatch(self) {"),
				GoStmt.GoRaw("\thxrt_eregThrowNoMatch()"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("raw := *hxrt.StdString(self.lastSource)"),
				GoStmt.GoRaw("end := self.lastIndices[1]"),
				GoStmt.GoRaw("if end > len(raw) {"),
				GoStmt.GoRaw("\tend = len(raw)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("raw[end:]")]))
			]),
			GoDecl.GoFuncDecl("split", {
				name: "self",
				typeName: "*EReg"
			}, [{name: "source", typeName: "*string"}], ["[]*string"], [
				GoStmt.GoRaw("raw := *hxrt.StdString(source)"),
				GoStmt.GoRaw("if self == nil || self.regex == nil {"),
				GoStmt.GoRaw("\treturn []*string{hxrt.StringFromLiteral(raw)}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("parts := self.regex.Split(raw, -1)"),
				GoStmt.GoRaw("out := make([]*string, 0, len(parts))"),
				GoStmt.GoRaw("for _, part := range parts {"),
				GoStmt.GoRaw("\tout = append(out, hxrt.StringFromLiteral(part))"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("out"))
			]),
			GoDecl.GoFuncDecl("replace", {
				name: "self",
				typeName: "*EReg"
			}, [
				{
					name: "source",
					typeName: "*string"
				},
				{name: "by", typeName: "*string"}
			], ["*string"], [
				GoStmt.GoRaw("if self == nil || self.regex == nil {"),
				GoStmt.GoRaw("\treturn source"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("rawSource := *hxrt.StdString(source)"),
				GoStmt.GoRaw("rawBy := *hxrt.StdString(by)"),
				GoStmt.GoRaw("if self.global {"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"),
					[
						GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoSelector(GoExpr.GoIdent("self"), "regex"), "ReplaceAllString"),
							[GoExpr.GoIdent("rawSource"), GoExpr.GoIdent("rawBy")])
					])),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("first := self.regex.FindStringSubmatchIndex(rawSource)"),
				GoStmt.GoRaw("if first == nil {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(rawSource)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("replacement := self.regex.ExpandString(nil, rawBy, rawSource, first)"),
				GoStmt.GoRaw("out := rawSource[:first[0]] + string(replacement) + rawSource[first[1]:]"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoIdent("out")]))
			]),
			GoDecl.GoFuncDecl("map_", {
				name: "self",
				typeName: "*EReg"
			}, [
				{
					name: "source",
					typeName: "*string"
				},
				{name: "callback", typeName: "func(*EReg) *string"}
			], ["*string"], [
				GoStmt.GoRaw("if self == nil || self.regex == nil {"),
				GoStmt.GoRaw("\treturn source"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("raw := *hxrt.StdString(source)"),
				GoStmt.GoRaw("if callback == nil {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(raw)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("var matches [][]int"),
				GoStmt.GoRaw("if self.global {"),
				GoStmt.GoRaw("\tmatches = self.regex.FindAllStringSubmatchIndex(raw, -1)"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\tif first := self.regex.FindStringSubmatchIndex(raw); first != nil {"),
				GoStmt.GoRaw("\t\tmatches = [][]int{first}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if len(matches) == 0 {"),
				GoStmt.GoRaw("\treturn hxrt.StringFromLiteral(raw)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("var builder strings.Builder"),
				GoStmt.GoRaw("cursor := 0"),
				GoStmt.GoRaw("for _, match := range matches {"),
				GoStmt.GoRaw("\tif len(match) < 2 {"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tstart := match[0]"),
				GoStmt.GoRaw("\tend := match[1]"),
				GoStmt.GoRaw("\tif start < cursor {"),
				GoStmt.GoRaw("\t\tstart = cursor"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif end < start {"),
				GoStmt.GoRaw("\t\tend = start"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif start > len(raw) {"),
				GoStmt.GoRaw("\t\tstart = len(raw)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif end > len(raw) {"),
				GoStmt.GoRaw("\t\tend = len(raw)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tbuilder.WriteString(raw[cursor:start])"),
				GoStmt.GoRaw("\tindices := make([]int, len(match))"),
				GoStmt.GoRaw("\tcopy(indices, match)"),
				GoStmt.GoRaw("\tself.lastSource = hxrt.StringFromLiteral(raw)"),
				GoStmt.GoRaw("\tself.lastIndices = indices"),
				GoStmt.GoRaw("\treplacement := callback(self)"),
				GoStmt.GoRaw("\tbuilder.WriteString(*hxrt.StdString(replacement))"),
				GoStmt.GoRaw("\tcursor = end"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("builder.WriteString(raw[cursor:])"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"),
					[GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("builder"), "String"), [])]))
			]),
			GoDecl.GoStructDecl("haxe__SerializedDate",
				[
					{
						name: "ms",
						typeName: "float64"
					}
				]),
			GoDecl.GoStructDecl("haxe__SerializedBytes", [{name: "data", typeName: "[]byte"}]),
			GoDecl.GoStructDecl("haxe__SerializedClassRef", [{name: "name", typeName: "string"}]),
			GoDecl.GoStructDecl("haxe__SerializedEnumRef", [{name: "name", typeName: "string"}]),
			GoDecl.GoStructDecl("haxe__SerializedClass", [
				{
					name: "name",
					typeName: "string"
				},
				{name: "fieldNames", typeName: "[]string"},
				{name: "fieldValues", typeName: "[]any"}
			]),
			GoDecl.GoStructDecl("haxe__SerializedEnum",
				[
					{
						name: "name",
						typeName: "string"
					},
					{name: "constructor", typeName: "string"},
					{name: "constructorIndex", typeName: "int"},
					{name: "hasConstructorIndex", typeName: "bool"},
					{name: "args", typeName: "[]any"}
				]),
			GoDecl.GoStructDecl("haxe__Unserializer__DefaultResolver", []),
			GoDecl.GoStructDecl("haxe__Unserializer__NullResolver", []),
			GoDecl.GoGlobalVarDecl("haxe__Serializer_USE_CACHE", "bool", GoExpr.GoBoolLiteral(false)),
			GoDecl.GoGlobalVarDecl("haxe__Serializer_USE_ENUM_INDEX", "bool", GoExpr.GoBoolLiteral(false)),
			GoDecl.GoGlobalVarDecl("haxe__Unserializer_DEFAULT_RESOLVER", "any", GoExpr.GoRaw("&haxe__Unserializer__DefaultResolver{}")),
			GoDecl.GoGlobalVarDecl("haxe__Unserializer_NULL_RESOLVER", "any", GoExpr.GoRaw("&haxe__Unserializer__NullResolver{}")),
			GoDecl.GoFuncDecl("hxrt_serializerLookupClassName", null, [
				{
					name: "typeName",
					typeName: "string"
				}
			], ["string", "bool"], classLookupBody),
			GoDecl.GoFuncDecl("hxrt_serializerLookupEnumConstructor", null, [
				{
					name: "typeName",
					typeName: "string"
				},
				{name: "tag", typeName: "int"}
			], ["string", "string", "bool"],
				enumLookupBody),
			GoDecl.GoFuncDecl("hxrt_serializerLookupEnumConstructorByName", null, [
				{
					name: "enumName",
					typeName: "string"
				},
				{name: "index", typeName: "int"}
			], ["string", "bool"], enumLookupByNameBody),
			GoDecl.GoFuncDecl("hxrt_serializerLookupEnumIndexByName", null, [
				{
					name: "enumName",
					typeName: "string"
				},
				{name: "constructorName", typeName: "string"}
			], ["int", "bool"], enumLookupIndexBody),
			GoDecl.GoFuncDecl("resolveClass", {
				name: "self",
				typeName: "*haxe__Unserializer__DefaultResolver"
			}, [{name: "name", typeName: "*string"}], ["any"], [
				GoStmt.GoRaw("className := *hxrt.StdString(name)"),
				GoStmt.GoRaw("if !hxrt_unserializerHasClass(className) {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return &haxe__SerializedClassRef{name: className}")
			]),
			GoDecl.GoFuncDecl("resolveEnum", {
				name: "self",
				typeName: "*haxe__Unserializer__DefaultResolver"
			}, [{name: "name", typeName: "*string"}], ["any"], [
				GoStmt.GoRaw("enumName := *hxrt.StdString(name)"),
				GoStmt.GoRaw("if !hxrt_unserializerHasEnum(enumName) {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return &haxe__SerializedEnumRef{name: enumName}")
			]),
			GoDecl.GoFuncDecl("resolveClass", {
				name: "self",
				typeName: "*haxe__Unserializer__NullResolver"
			}, [{name: "name", typeName: "*string"}],
				["any"], [GoStmt.GoReturn(GoExpr.GoNil)]),
			GoDecl.GoFuncDecl("resolveEnum", {
				name: "self",
				typeName: "*haxe__Unserializer__NullResolver"
			},
				[{name: "name", typeName: "*string"}], ["any"], [GoStmt.GoReturn(GoExpr.GoNil)]),
			GoDecl.GoFuncDecl("hxrt_unserializerHasClass", null, [{name: "className", typeName: "string"}], ["bool"], classExistsBody),
			GoDecl.GoFuncDecl("hxrt_unserializerHasEnum", null, [{name: "enumName", typeName: "string"}], ["bool"], enumExistsBody),
			GoDecl.GoFuncDecl("hxrt_unserializerBindSelf", null, [{name: "instance", typeName: "any"}], [], [
				GoStmt.GoRaw("if instance == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("rv := reflect.ValueOf(instance)"),
				GoStmt.GoRaw("if !rv.IsValid() || rv.Kind() != reflect.Pointer || rv.IsNil() {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("elem := rv.Elem()"),
				GoStmt.GoRaw("if !elem.IsValid() || elem.Kind() != reflect.Struct {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("field := elem.FieldByName(\"__hx_this\")"),
				GoStmt.GoRaw("if !field.IsValid() {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !rv.Type().AssignableTo(field.Type()) {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if field.CanSet() {"),
				GoStmt.GoRaw("\tfield.Set(rv)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !field.CanAddr() {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("lifted := reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()"),
				GoStmt.GoRaw("lifted.Set(rv)")
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerInvokeResolver", null, [
				{
					name: "resolver",
					typeName: "any"
				},
				{name: "methodName", typeName: "string"},
				{name: "name", typeName: "string"}
			], ["any", "bool"], [
				GoStmt.GoRaw("if resolver == nil {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("result := any(nil)"),
				GoStmt.GoRaw("ok := false"),
				GoStmt.GoRaw("defer func() {"),
				GoStmt.GoRaw("\tif recover() != nil {"),
				GoStmt.GoRaw("\t\tresult = nil"),
				GoStmt.GoRaw("\t\tok = false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}()"),
				GoStmt.GoRaw("rv := reflect.ValueOf(resolver)"),
				GoStmt.GoRaw("if !rv.IsValid() {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for rv.IsValid() && rv.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\tif rv.IsNil() {"),
				GoStmt.GoRaw("\t\treturn nil, false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\trv = rv.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("method := reflect.Value{}"),
				GoStmt.GoRaw("if rv.IsValid() {"),
				GoStmt.GoRaw("\tmethod = rv.MethodByName(methodName)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !method.IsValid() && rv.IsValid() && rv.Kind() != reflect.Pointer && rv.CanAddr() {"),
				GoStmt.GoRaw("\tmethod = rv.Addr().MethodByName(methodName)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !method.IsValid() && rv.IsValid() && rv.Kind() == reflect.Pointer && !rv.IsNil() {"),
				GoStmt.GoRaw("\tmethod = rv.Elem().MethodByName(methodName)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !method.IsValid() && rv.IsValid() {"),
				GoStmt.GoRaw("\tswitch rv.Kind() {"),
				GoStmt.GoRaw("\tcase reflect.Struct:"),
				GoStmt.GoRaw("\t\tfield := rv.FieldByName(methodName)"),
				GoStmt.GoRaw("\t\tif field.IsValid() {"),
				GoStmt.GoRaw("\t\t\tfor field.IsValid() && field.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\t\t\t\tif field.IsNil() {"),
				GoStmt.GoRaw("\t\t\t\t\tbreak"),
				GoStmt.GoRaw("\t\t\t\t}"),
				GoStmt.GoRaw("\t\t\t\tfield = field.Elem()"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t\tif field.IsValid() && field.Kind() == reflect.Func {"),
				GoStmt.GoRaw("\t\t\t\tmethod = field"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\tcase reflect.Pointer:"),
				GoStmt.GoRaw("\t\tif !rv.IsNil() && rv.Elem().Kind() == reflect.Struct {"),
				GoStmt.GoRaw("\t\t\tfield := rv.Elem().FieldByName(methodName)"),
				GoStmt.GoRaw("\t\t\tif field.IsValid() {"),
				GoStmt.GoRaw("\t\t\t\tfor field.IsValid() && field.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\t\t\t\t\tif field.IsNil() {"),
				GoStmt.GoRaw("\t\t\t\t\t\tbreak"),
				GoStmt.GoRaw("\t\t\t\t\t}"),
				GoStmt.GoRaw("\t\t\t\t\tfield = field.Elem()"),
				GoStmt.GoRaw("\t\t\t\t}"),
				GoStmt.GoRaw("\t\t\t\tif field.IsValid() && field.Kind() == reflect.Func {"),
				GoStmt.GoRaw("\t\t\t\t\tmethod = field"),
				GoStmt.GoRaw("\t\t\t\t}"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\tcase reflect.Map:"),
				GoStmt.GoRaw("\t\tif rv.Type().Key().Kind() == reflect.String {"),
				GoStmt.GoRaw("\t\t\tfield := rv.MapIndex(reflect.ValueOf(methodName))"),
				GoStmt.GoRaw("\t\t\tif field.IsValid() {"),
				GoStmt.GoRaw("\t\t\t\tfor field.IsValid() && field.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\t\t\t\t\tif field.IsNil() {"),
				GoStmt.GoRaw("\t\t\t\t\t\tbreak"),
				GoStmt.GoRaw("\t\t\t\t\t}"),
				GoStmt.GoRaw("\t\t\t\t\tfield = field.Elem()"),
				GoStmt.GoRaw("\t\t\t\t}"),
				GoStmt.GoRaw("\t\t\t\tif field.IsValid() && field.Kind() == reflect.Func {"),
				GoStmt.GoRaw("\t\t\t\t\tmethod = field"),
				GoStmt.GoRaw("\t\t\t\t}"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !method.IsValid() || method.Kind() != reflect.Func {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("methodType := method.Type()"),
				GoStmt.GoRaw("if methodType.NumIn() != 1 || methodType.NumOut() < 1 {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("argType := methodType.In(0)"),
				GoStmt.GoRaw("nameValue := reflect.ValueOf(name)"),
				GoStmt.GoRaw("var arg reflect.Value"),
				GoStmt.GoRaw("if nameValue.Type().AssignableTo(argType) {"),
				GoStmt.GoRaw("\targ = nameValue"),
				GoStmt.GoRaw("} else if nameValue.Type().ConvertibleTo(argType) {"),
				GoStmt.GoRaw("\targ = nameValue.Convert(argType)"),
				GoStmt.GoRaw("} else if argType.Kind() == reflect.Pointer && argType.Elem().Kind() == reflect.String {"),
				GoStmt.GoRaw("\tnameCopy := name"),
				GoStmt.GoRaw("\targ = reflect.ValueOf(&nameCopy)"),
				GoStmt.GoRaw("\tif !arg.Type().AssignableTo(argType) {"),
				GoStmt.GoRaw("\t\tif arg.Type().ConvertibleTo(argType) {"),
				GoStmt.GoRaw("\t\t\targ = arg.Convert(argType)"),
				GoStmt.GoRaw("\t\t} else {"),
				GoStmt.GoRaw("\t\t\treturn nil, false"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("out := method.Call([]reflect.Value{arg})"),
				GoStmt.GoRaw("if len(out) == 0 {"),
				GoStmt.GoRaw("\treturn nil, true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("value := out[0]"),
				GoStmt.GoRaw("if !value.IsValid() {"),
				GoStmt.GoRaw("\treturn nil, true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for value.IsValid() && value.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\tif value.IsNil() {"),
				GoStmt.GoRaw("\t\treturn nil, true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tvalue = value.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !value.IsValid() {"),
				GoStmt.GoRaw("\treturn nil, true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch value.Kind() {"),
				GoStmt.GoRaw("case reflect.Pointer, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan:"),
				GoStmt.GoRaw("\tif value.IsNil() {"),
				GoStmt.GoRaw("\t\treturn nil, true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("result = value.Interface()"),
				GoStmt.GoRaw("ok = true"),
				GoStmt.GoRaw("return result, ok")
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerResolveClass", null, [
				{
					name: "self",
					typeName: "*haxe__Unserializer"
				},
				{name: "name", typeName: "string"}
			], ["any"], [
				GoStmt.GoRaw("var resolver any"),
				GoStmt.GoRaw("if self != nil {"),
				GoStmt.GoRaw("\tresolver = self.resolver"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if resolver == nil {"),
				GoStmt.GoRaw("\tresolver = haxe__Unserializer_DEFAULT_RESOLVER"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if resolver == nil {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch current := resolver.(type) {"),
				GoStmt.GoRaw("case interface{ resolveClass(*string) any }:"),
				GoStmt.GoRaw("\treturn current.resolveClass(hxrt.StringFromLiteral(name))"),
				GoStmt.GoRaw("case interface{ resolveClass(string) any }:"),
				GoStmt.GoRaw("\treturn current.resolveClass(name)"),
				GoStmt.GoRaw("case interface{ resolveClass(any) any }:"),
				GoStmt.GoRaw("\treturn current.resolveClass(name)"),
				GoStmt.GoRaw("case interface{ ResolveClass(*string) any }:"),
				GoStmt.GoRaw("\treturn current.ResolveClass(hxrt.StringFromLiteral(name))"),
				GoStmt.GoRaw("case interface{ ResolveClass(string) any }:"),
				GoStmt.GoRaw("\treturn current.ResolveClass(name)"),
				GoStmt.GoRaw("case interface{ ResolveClass(any) any }:"),
				GoStmt.GoRaw("\treturn current.ResolveClass(name)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("resolved, ok := hxrt_unserializerInvokeResolver(resolver, \"resolveClass\", name)"),
				GoStmt.GoRaw("if ok {"),
				GoStmt.GoRaw("\treturn resolved"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("resolved, ok = hxrt_unserializerInvokeResolver(resolver, \"ResolveClass\", name)"),
				GoStmt.GoRaw("if ok {"),
				GoStmt.GoRaw("\treturn resolved"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return nil")
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerResolveEnum", null, [
				{
					name: "self",
					typeName: "*haxe__Unserializer"
				},
				{name: "name", typeName: "string"}
			], ["any"], [
				GoStmt.GoRaw("var resolver any"),
				GoStmt.GoRaw("if self != nil {"),
				GoStmt.GoRaw("\tresolver = self.resolver"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if resolver == nil {"),
				GoStmt.GoRaw("\tresolver = haxe__Unserializer_DEFAULT_RESOLVER"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if resolver == nil {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch current := resolver.(type) {"),
				GoStmt.GoRaw("case interface{ resolveEnum(*string) any }:"),
				GoStmt.GoRaw("\treturn current.resolveEnum(hxrt.StringFromLiteral(name))"),
				GoStmt.GoRaw("case interface{ resolveEnum(string) any }:"),
				GoStmt.GoRaw("\treturn current.resolveEnum(name)"),
				GoStmt.GoRaw("case interface{ resolveEnum(any) any }:"),
				GoStmt.GoRaw("\treturn current.resolveEnum(name)"),
				GoStmt.GoRaw("case interface{ ResolveEnum(*string) any }:"),
				GoStmt.GoRaw("\treturn current.ResolveEnum(hxrt.StringFromLiteral(name))"),
				GoStmt.GoRaw("case interface{ ResolveEnum(string) any }:"),
				GoStmt.GoRaw("\treturn current.ResolveEnum(name)"),
				GoStmt.GoRaw("case interface{ ResolveEnum(any) any }:"),
				GoStmt.GoRaw("\treturn current.ResolveEnum(name)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("resolved, ok := hxrt_unserializerInvokeResolver(resolver, \"resolveEnum\", name)"),
				GoStmt.GoRaw("if ok {"),
				GoStmt.GoRaw("\treturn resolved"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("resolved, ok = hxrt_unserializerInvokeResolver(resolver, \"ResolveEnum\", name)"),
				GoStmt.GoRaw("if ok {"),
				GoStmt.GoRaw("\treturn resolved"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return nil")
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerExtractNameField", null, [
				{
					name: "resolved",
					typeName: "any"
				}
			], ["string", "bool"], [
				GoStmt.GoRaw("rv := reflect.ValueOf(resolved)"),
				GoStmt.GoRaw("if !rv.IsValid() {"),
				GoStmt.GoRaw("\treturn \"\", false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for rv.IsValid() && (rv.Kind() == reflect.Interface || rv.Kind() == reflect.Pointer) {"),
				GoStmt.GoRaw("\tif rv.IsNil() {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\trv = rv.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !rv.IsValid() || rv.Kind() != reflect.Struct {"),
				GoStmt.GoRaw("\treturn \"\", false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("field := rv.FieldByName(\"name\")"),
				GoStmt.GoRaw("if !field.IsValid() {"),
				GoStmt.GoRaw("\treturn \"\", false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for field.IsValid() && field.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\tif field.IsNil() {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tfield = field.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !field.IsValid() {"),
				GoStmt.GoRaw("\treturn \"\", false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if field.Kind() == reflect.String {"),
				GoStmt.GoRaw("\treturn field.String(), true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if field.Kind() == reflect.Pointer {"),
				GoStmt.GoRaw("\tif field.IsNil() {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif field.Elem().Kind() == reflect.String {"),
				GoStmt.GoRaw("\t\treturn field.Elem().String(), true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return \"\", false")
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerResolvedClassName", null, [
				{
					name: "resolved",
					typeName: "any"
				}
			], ["string", "bool"], [
				GoStmt.GoRaw("switch current := resolved.(type) {"),
				GoStmt.GoRaw("case haxe__SerializedClassRef:"),
				GoStmt.GoRaw("\treturn current.name, true"),
				GoStmt.GoRaw("case *haxe__SerializedClassRef:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn current.name, true"),
				GoStmt.GoRaw("case string:"),
				GoStmt.GoRaw("\treturn current, true"),
				GoStmt.GoRaw("case *string:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn *current, true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return hxrt_unserializerExtractNameField(resolved)")
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerResolvedEnumName", null, [
				{
					name: "resolved",
					typeName: "any"
				}
			], ["string", "bool"], [
				GoStmt.GoRaw("switch current := resolved.(type) {"),
				GoStmt.GoRaw("case haxe__SerializedEnumRef:"),
				GoStmt.GoRaw("\treturn current.name, true"),
				GoStmt.GoRaw("case *haxe__SerializedEnumRef:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn current.name, true"),
				GoStmt.GoRaw("case string:"),
				GoStmt.GoRaw("\treturn current, true"),
				GoStmt.GoRaw("case *string:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\treturn \"\", false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn *current, true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return hxrt_unserializerExtractNameField(resolved)")
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerCreateClassInstance", null, [
				{
					name: "className",
					typeName: "string"
				}
			], ["any", "bool"], classCreateBody),
			GoDecl.GoFuncDecl("hxrt_unserializerCreateEnumInstance", null, [
				{
					name: "enumName",
					typeName: "string"
				},
				{name: "constructorName", typeName: "string"},
				{name: "constructorIndex", typeName: "int"},
				{name: "hasConstructorIndex", typeName: "bool"},
				{name: "args", typeName: "[]any"}
			], ["any", "bool"], [
				GoStmt.GoRaw("if hasConstructorIndex {"),
				GoStmt.GoRaw("\tif _, ok := hxrt_serializerLookupEnumConstructorByName(enumName, constructorIndex); !ok {"),
				GoStmt.GoRaw("\t\treturn nil, false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\tresolvedIndex, ok := hxrt_serializerLookupEnumIndexByName(enumName, constructorName)"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\treturn nil, false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tconstructorIndex = resolvedIndex"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch enumName {")
			].concat(enumCreateBody.slice(1))),
			GoDecl.GoStructDecl("haxe__Serializer", [
				{
					name: "buf",
					typeName: "*string"
				},
				{name: "useCache", typeName: "bool"},
				{name: "useEnumIndex", typeName: "bool"},
				{name: "stringCache", typeName: "map[string]int"},
				{name: "cacheRefs", typeName: "map[uintptr]int"}
			]),
			GoDecl.GoFuncDecl("New_haxe__Serializer", null, [], ["*haxe__Serializer"], [
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__Serializer{buf: hxrt.StringFromLiteral(\"\"), useCache: haxe__Serializer_USE_CACHE, useEnumIndex: haxe__Serializer_USE_ENUM_INDEX, stringCache: map[string]int{}, cacheRefs: map[uintptr]int{}}"))
			]),
			GoDecl.GoFuncDecl("serialize", {
				name: "self",
				typeName: "*haxe__Serializer"
			}, [{name: "value", typeName: "any"}], [], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(null)], null),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt_serializerWriteValue"), [GoExpr.GoIdent("self"), GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("serializeException", {
				name: "self",
				typeName: "*haxe__Serializer"
			}, [{name: "value", typeName: "any"}], [], [
				GoStmt.GoIf(GoExpr.GoBinary("==", GoExpr.GoIdent("self"), GoExpr.GoNil), [GoStmt.GoReturn(null)], null),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"x\")"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoIdent("hxrt_serializerWriteValue"), [GoExpr.GoIdent("self"), GoExpr.GoIdent("value")]))
			]),
			GoDecl.GoFuncDecl("toString", {
				name: "self",
				typeName: "*haxe__Serializer"
			}, [], ["*string"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.buf == nil"), [
					GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoStringLiteral("")]))
				], null),
				GoStmt.GoReturn(GoExpr.GoSelector(GoExpr.GoIdent("self"), "buf"))
			]),
			GoDecl.GoFuncDecl("hxrt_serializerAppend", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "chunk", typeName: "string"}
			], [], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.buf == nil {"),
				GoStmt.GoRaw("\tself.buf = hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.buf = hxrt.StringFromLiteral(*hxrt.StdString(self.buf) + chunk)")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerEscape", null, [
				{
					name: "value",
					typeName: "*string"
				}
			], ["string"], [
				GoStmt.GoRaw("raw := *hxrt.StdString(value)"),
				GoStmt.GoRaw("var builder strings.Builder"),
				GoStmt.GoRaw("hex := \"0123456789ABCDEF\""),
				GoStmt.GoRaw("for i := 0; i < len(raw); i++ {"),
				GoStmt.GoRaw("\tb := raw[i]"),
				GoStmt.GoRaw("\tif (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') || b == '_' || b == '.' || b == '-' {"),
				GoStmt.GoRaw("\t\tbuilder.WriteByte(b)"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tbuilder.WriteByte('%')"),
				GoStmt.GoRaw("\tbuilder.WriteByte(hex[b>>4])"),
				GoStmt.GoRaw("\tbuilder.WriteByte(hex[b&0x0F])"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("builder"), "String"), []))
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteStringToken", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "value", typeName: "string"}
			], [], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.stringCache == nil {"),
				GoStmt.GoRaw("\tself.stringCache = map[string]int{}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if index, ok := self.stringCache[value]; ok {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"R\" + strconv.Itoa(index))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("escaped := hxrt_serializerEscape(hxrt.StringFromLiteral(value))"),
				GoStmt.GoRaw("self.stringCache[value] = len(self.stringCache)"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"y\" + strconv.Itoa(len(escaped)) + \":\" + escaped)")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteIntToken", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "value", typeName: "int64"}
			], [], [
				GoStmt.GoRaw("if value == 0 {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"z\")"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"i\" + strconv.FormatInt(value, 10))")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteBytesToken", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "raw", typeName: "[]byte"}
			], [], [
				GoStmt.GoRaw("encoded := base64.StdEncoding.EncodeToString(raw)"),
				GoStmt.GoRaw("encoded = strings.TrimRight(encoded, \"=\")"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"s\" + strconv.Itoa(len(encoded)) + \":\" + encoded)")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteEnumToken", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "enumName", typeName: "string"},
				{name: "constructorName", typeName: "string"},
				{name: "constructorIndex", typeName: "int"},
				{name: "hasConstructorIndex", typeName: "bool"},
				{name: "args", typeName: "[]any"}
			], [], [
				GoStmt.GoRaw("if self != nil && self.useEnumIndex && hasConstructorIndex {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"j\")"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, enumName)"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \":\" + strconv.Itoa(constructorIndex))"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"w\")"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, enumName)"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, constructorName)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \":\" + strconv.Itoa(len(args)))"),
				GoStmt.GoRaw("for _, arg := range args {"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, arg)"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteListToken", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "items", typeName: "[]any"}
			], [], [
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"l\")"),
				GoStmt.GoRaw("for _, item := range items {"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, item)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"h\")")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteStringMapToken", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "entries", typeName: "map[string]any"}
			], [], [
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"b\")"),
				GoStmt.GoRaw("keys := make([]string, 0, len(entries))"),
				GoStmt.GoRaw("for key := range entries {"),
				GoStmt.GoRaw("\tkeys = append(keys, key)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("sort.Strings(keys)"),
				GoStmt.GoRaw("for _, key := range keys {"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, key)"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, entries[key])"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"h\")")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteIntMapToken", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "entries", typeName: "map[int]any"}
			], [], [
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"q\")"),
				GoStmt.GoRaw("keys := make([]int, 0, len(entries))"),
				GoStmt.GoRaw("for key := range entries {"),
				GoStmt.GoRaw("\tkeys = append(keys, key)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("sort.Ints(keys)"),
				GoStmt.GoRaw("for _, key := range keys {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \":\" + strconv.Itoa(key))"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, entries[key])"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"h\")")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteObjectMapToken", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "entries", typeName: "map[any]any"}
			], [], [
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"M\")"),
				GoStmt.GoRaw("for key, value := range entries {"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, key)"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, value)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"h\")")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerReflectAny", null, [
				{
					name: "value",
					typeName: "reflect.Value"
				}
			], ["any", "bool"], [
				GoStmt.GoRaw("defer func() {"),
				GoStmt.GoRaw("\tif recover() != nil {"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}()"),
				GoStmt.GoRaw("if !value.IsValid() {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if value.CanInterface() {"),
				GoStmt.GoRaw("\treturn value.Interface(), true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !value.CanAddr() {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("lifted := reflect.NewAt(value.Type(), unsafe.Pointer(value.UnsafeAddr())).Elem()"),
				GoStmt.GoRaw("if !lifted.IsValid() || !lifted.CanInterface() {"),
				GoStmt.GoRaw("\treturn nil, false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return lifted.Interface(), true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteSerializedClass", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "serialized", typeName: "*haxe__SerializedClass"}
			], [], [
				GoStmt.GoRaw("if serialized == nil {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"c\")"),
				GoStmt.GoRaw("hxrt_serializerWriteStringToken(self, serialized.name)"),
				GoStmt.GoRaw("limit := len(serialized.fieldNames)"),
				GoStmt.GoRaw("if len(serialized.fieldValues) < limit {"),
				GoStmt.GoRaw("\tlimit = len(serialized.fieldValues)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for i := 0; i < limit; i++ {"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, serialized.fieldNames[i])"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, serialized.fieldValues[i])"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"g\")")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteSerializedEnum", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "serialized", typeName: "*haxe__SerializedEnum"}
			], [], [
				GoStmt.GoRaw("if serialized == nil {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("constructorIndex := serialized.constructorIndex"),
				GoStmt.GoRaw("hasConstructorIndex := serialized.hasConstructorIndex"),
				GoStmt.GoRaw("if !hasConstructorIndex {"),
				GoStmt.GoRaw("\tif resolvedIndex, ok := hxrt_serializerLookupEnumIndexByName(serialized.name, serialized.constructor); ok {"),
				GoStmt.GoRaw("\t\tconstructorIndex = resolvedIndex"),
				GoStmt.GoRaw("\t\thasConstructorIndex = true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerWriteEnumToken(self, serialized.name, serialized.constructor, constructorIndex, hasConstructorIndex, serialized.args)")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTryDsListStruct", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "ref", typeName: "reflect.Value"}
			], ["bool"], [
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct || ref.Type().Name() != \"haxe__ds__List\" {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("itemsField := ref.FieldByName(\"items\")"),
				GoStmt.GoRaw("if !itemsField.IsValid() || itemsField.Kind() != reflect.Slice {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("items := make([]any, 0, itemsField.Len())"),
				GoStmt.GoRaw("for i := 0; i < itemsField.Len(); i++ {"),
				GoStmt.GoRaw("\titem, ok := hxrt_serializerReflectAny(itemsField.Index(i))"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\titems = append(items, item)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerWriteListToken(self, items)"),
				GoStmt.GoRaw("return true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTryDsStringMapStruct", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "ref", typeName: "reflect.Value"}
			], ["bool"], [
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct || ref.Type().Name() != \"haxe__ds__StringMap\" {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("mapField := ref.FieldByName(\"h\")"),
				GoStmt.GoRaw("if !mapField.IsValid() || mapField.Kind() != reflect.Map {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("entries := map[string]any{}"),
				GoStmt.GoRaw("for _, key := range mapField.MapKeys() {"),
				GoStmt.GoRaw("\tif key.Kind() != reflect.String {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tvalue, ok := hxrt_serializerReflectAny(mapField.MapIndex(key))"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tentries[key.String()] = value"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerWriteStringMapToken(self, entries)"),
				GoStmt.GoRaw("return true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTryDsIntMapStruct", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "ref", typeName: "reflect.Value"}
			], ["bool"], [
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct || ref.Type().Name() != \"haxe__ds__IntMap\" {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("mapField := ref.FieldByName(\"h\")"),
				GoStmt.GoRaw("if !mapField.IsValid() || mapField.Kind() != reflect.Map {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("entries := map[int]any{}"),
				GoStmt.GoRaw("for _, key := range mapField.MapKeys() {"),
				GoStmt.GoRaw("\tvar intKey int"),
				GoStmt.GoRaw("\tswitch key.Kind() {"),
				GoStmt.GoRaw("\tcase reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:"),
				GoStmt.GoRaw("\t\tintKey = int(key.Int())"),
				GoStmt.GoRaw("\tdefault:"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tvalue, ok := hxrt_serializerReflectAny(mapField.MapIndex(key))"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tentries[intKey] = value"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerWriteIntMapToken(self, entries)"),
				GoStmt.GoRaw("return true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTryDsObjectMapStruct", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "ref", typeName: "reflect.Value"}
			], ["bool"], [
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct || ref.Type().Name() != \"haxe__ds__ObjectMap\" {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("mapField := ref.FieldByName(\"h\")"),
				GoStmt.GoRaw("if !mapField.IsValid() || mapField.Kind() != reflect.Map {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("entries := map[any]any{}"),
				GoStmt.GoRaw("for _, key := range mapField.MapKeys() {"),
				GoStmt.GoRaw("\tkeyAny, ok := hxrt_serializerReflectAny(key)"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tvalueAny, ok := hxrt_serializerReflectAny(mapField.MapIndex(key))"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tentries[keyAny] = valueAny"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerWriteObjectMapToken(self, entries)"),
				GoStmt.GoRaw("return true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTryClassStruct", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "value", typeName: "any"},
				{name: "ref", typeName: "reflect.Value"}
			], ["bool"], [
				GoStmt.GoRaw("defer func() {"),
				GoStmt.GoRaw("\tif recover() != nil {"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}()"),
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("className, ok := hxrt_serializerLookupClassName(ref.Type().Name())"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTrackRef(self, value) {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if custom, ok := value.(interface{ hxSerialize(*haxe__Serializer) }); ok {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"C\")"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, className)"),
				GoStmt.GoRaw("\tcustom.hxSerialize(self)"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"g\")"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if custom, ok := value.(interface{ HxSerialize(*haxe__Serializer) }); ok {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"C\")"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, className)"),
				GoStmt.GoRaw("\tcustom.HxSerialize(self)"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"g\")"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if ref.CanAddr() {"),
				GoStmt.GoRaw("\taddr := ref.Addr().Interface()"),
				GoStmt.GoRaw("\tif custom, ok := addr.(interface{ hxSerialize(*haxe__Serializer) }); ok {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"C\")"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteStringToken(self, className)"),
				GoStmt.GoRaw("\t\tcustom.hxSerialize(self)"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"g\")"),
				GoStmt.GoRaw("\t\treturn true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif custom, ok := addr.(interface{ HxSerialize(*haxe__Serializer) }); ok {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"C\")"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteStringToken(self, className)"),
				GoStmt.GoRaw("\t\tcustom.HxSerialize(self)"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"g\")"),
				GoStmt.GoRaw("\t\treturn true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"c\")"),
				GoStmt.GoRaw("hxrt_serializerWriteStringToken(self, className)"),
				GoStmt.GoRaw("refType := ref.Type()"),
				GoStmt.GoRaw("for i := 0; i < ref.NumField(); i++ {"),
				GoStmt.GoRaw("\tfieldInfo := refType.Field(i)"),
				GoStmt.GoRaw("\tfieldName := fieldInfo.Name"),
				GoStmt.GoRaw("\tif strings.HasPrefix(fieldName, \"__hx_\") {"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tfieldValue, ok := hxrt_serializerReflectAny(ref.Field(i))"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, fieldName)"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, fieldValue)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"g\")"),
				GoStmt.GoRaw("return true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTryEnumStruct", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "value", typeName: "any"},
				{name: "ref", typeName: "reflect.Value"}
			], ["bool"], [
				GoStmt.GoRaw("defer func() {"),
				GoStmt.GoRaw("\tif recover() != nil {"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}()"),
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("tagField := ref.FieldByName(\"tag\")"),
				GoStmt.GoRaw("paramsField := ref.FieldByName(\"params\")"),
				GoStmt.GoRaw("if !tagField.IsValid() || !paramsField.IsValid() || paramsField.Kind() != reflect.Slice {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("var tag int"),
				GoStmt.GoRaw("switch tagField.Kind() {"),
				GoStmt.GoRaw("case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:"),
				GoStmt.GoRaw("\ttag = int(tagField.Int())"),
				GoStmt.GoRaw("default:"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("enumName, constructorName, ok := hxrt_serializerLookupEnumConstructor(ref.Type().Name(), tag)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTrackRef(self, value) {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("args := make([]any, 0, paramsField.Len())"),
				GoStmt.GoRaw("for i := 0; i < paramsField.Len(); i++ {"),
				GoStmt.GoRaw("\tvalue, ok := hxrt_serializerReflectAny(paramsField.Index(i))"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\targs = append(args, value)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerWriteEnumToken(self, enumName, constructorName, tag, true, args)"),
				GoStmt.GoRaw("return true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTryDateStruct", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "ref", typeName: "reflect.Value"}
			], ["bool"], [
				GoStmt.GoRaw("defer func() {"),
				GoStmt.GoRaw("\tif recover() != nil {"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}()"),
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("valueField := ref.FieldByName(\"value\")"),
				GoStmt.GoRaw("if !valueField.IsValid() || !valueField.CanAddr() {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("fieldType := valueField.Type()"),
				GoStmt.GoRaw("if fieldType.PkgPath() != \"time\" || fieldType.Name() != \"Time\" {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("timeAny := reflect.NewAt(fieldType, unsafe.Pointer(valueField.UnsafeAddr())).Elem().Interface()"),
				GoStmt.GoRaw("timeValue, ok := timeAny.(time.Time)"),
				GoStmt.GoRaw("if !ok {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("ms := float64(timeValue.UnixNano()) / 1000000.0"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"v\" + strconv.FormatFloat(ms, 'g', -1, 64))"),
				GoStmt.GoRaw("return true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTrySpecialReflect", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "value", typeName: "any"}
			], ["bool"], [
				GoStmt.GoRaw("ref := reflect.ValueOf(value)"),
				GoStmt.GoRaw("if !ref.IsValid() {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for ref.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\tif ref.IsNil() {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tref = ref.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for ref.Kind() == reflect.Pointer {"),
				GoStmt.GoRaw("\tif ref.IsNil() {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tref = ref.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("typeName := ref.Type().Name()"),
				GoStmt.GoRaw("if typeName == \"Date\" {"),
				GoStmt.GoRaw("\tif hxrt_serializerTryDateStruct(self, ref) {"),
				GoStmt.GoRaw("\t\treturn true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if typeName == \"haxe__io__Bytes\" {"),
				GoStmt.GoRaw("\tbytesField := ref.FieldByName(\"b\")"),
				GoStmt.GoRaw("\tif !bytesField.IsValid() || bytesField.Kind() != reflect.Slice {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\traw := make([]byte, bytesField.Len())"),
				GoStmt.GoRaw("\tfor i := 0; i < bytesField.Len(); i++ {"),
				GoStmt.GoRaw("\t\tentry := bytesField.Index(i)"),
				GoStmt.GoRaw("\t\tif !entry.IsValid() {"),
				GoStmt.GoRaw("\t\t\treturn false"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tswitch entry.Kind() {"),
				GoStmt.GoRaw("\t\tcase reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:"),
				GoStmt.GoRaw("\t\t\traw[i] = byte(entry.Int())"),
				GoStmt.GoRaw("\t\tcase reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:"),
				GoStmt.GoRaw("\t\t\traw[i] = byte(entry.Uint())"),
				GoStmt.GoRaw("\t\tdefault:"),
				GoStmt.GoRaw("\t\t\treturn false"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerWriteBytesToken(self, raw)"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTryDsListStruct(self, ref) {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTryDsStringMapStruct(self, ref) {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTryDsIntMapStruct(self, ref) {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTryDsObjectMapStruct(self, ref) {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTryEnumStruct(self, value, ref) {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTryClassStruct(self, value, ref) {"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return false")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTryTypeValueRef", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "value", typeName: "any"}
			], ["bool"], [
				GoStmt.GoRaw("ref := reflect.ValueOf(value)"),
				GoStmt.GoRaw("if !ref.IsValid() {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for ref.Kind() == reflect.Interface || ref.Kind() == reflect.Pointer {"),
				GoStmt.GoRaw("\tif ref.IsNil() {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tref = ref.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !ref.IsValid() || ref.Kind() != reflect.Struct {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("typeName := ref.Type().Name()"),
				GoStmt.GoRaw("if typeName != \"hxrt__TypeClassValue\" && typeName != \"hxrt__TypeEnumValue\" {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("nameField := ref.FieldByName(\"name\")"),
				GoStmt.GoRaw("if !nameField.IsValid() {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for nameField.IsValid() && nameField.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\tif nameField.IsNil() {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tnameField = nameField.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if !nameField.IsValid() {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("resolvedName := \"\""),
				GoStmt.GoRaw("if nameField.Kind() == reflect.String {"),
				GoStmt.GoRaw("\tresolvedName = nameField.String()"),
				GoStmt.GoRaw("} else if nameField.Kind() == reflect.Pointer {"),
				GoStmt.GoRaw("\tif nameField.IsNil() || nameField.Elem().Kind() != reflect.String {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tresolvedName = nameField.Elem().String()"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if resolvedName == \"\" {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if typeName == \"hxrt__TypeClassValue\" {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"A\")"),
				GoStmt.GoRaw("} else {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"B\")"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt_serializerWriteStringToken(self, resolvedName)"),
				GoStmt.GoRaw("return true")
			]),
			GoDecl.GoFuncDecl("hxrt_serializerTrackRef", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "value", typeName: "any"}
			], ["bool"], [
				GoStmt.GoRaw("if self == nil || !self.useCache {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.cacheRefs == nil {"),
				GoStmt.GoRaw("\tself.cacheRefs = map[uintptr]int{}"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("ref := reflect.ValueOf(value)"),
				GoStmt.GoRaw("if !ref.IsValid() {"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("for ref.Kind() == reflect.Interface {"),
				GoStmt.GoRaw("\tif ref.IsNil() {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tref = ref.Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch ref.Kind() {"),
				GoStmt.GoRaw("case reflect.Map, reflect.Slice, reflect.Pointer:"),
				GoStmt.GoRaw("\tif ref.IsNil() {"),
				GoStmt.GoRaw("\t\treturn false"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tkey := ref.Pointer()"),
				GoStmt.GoRaw("\tif index, ok := self.cacheRefs[key]; ok {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"r\" + strconv.Itoa(index))"),
				GoStmt.GoRaw("\t\treturn true"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.cacheRefs[key] = len(self.cacheRefs)"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return false")
			]),
			GoDecl.GoFuncDecl("haxe__Serializer_run", null, [
				{
					name: "value",
					typeName: "any"
				}
			], ["*string"], [
				GoStmt.GoRaw("serializer := New_haxe__Serializer()"),
				GoStmt.GoExprStmt(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("serializer"), "serialize"), [GoExpr.GoIdent("value")])),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("serializer"), "toString"), []))
			]),
			GoDecl.GoFuncDecl("hxrt_serializerWriteValue", null, [
				{
					name: "self",
					typeName: "*haxe__Serializer"
				},
				{name: "value", typeName: "any"}
			], [], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if value == nil {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTryTypeValueRef(self, value) {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch current := value.(type) {"),
				GoStmt.GoRaw("case bool:"),
				GoStmt.GoRaw("\tif current {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"t\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"f\")"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case string:"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, current)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *string:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteStringToken(self, *current)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case haxe__SerializedDate:"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"v\" + strconv.FormatFloat(current.ms, 'g', -1, 64))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__SerializedDate:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"v\" + strconv.FormatFloat(current.ms, 'g', -1, 64))"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case haxe__SerializedBytes:"),
				GoStmt.GoRaw("\thxrt_serializerWriteBytesToken(self, current.data)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__SerializedBytes:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteBytesToken(self, current.data)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case haxe__SerializedClass:"),
				GoStmt.GoRaw("\thxrt_serializerWriteSerializedClass(self, &current)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__SerializedClass:"),
				GoStmt.GoRaw("\thxrt_serializerWriteSerializedClass(self, current)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case haxe__SerializedEnum:"),
				GoStmt.GoRaw("\thxrt_serializerWriteSerializedEnum(self, &current)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__SerializedEnum:"),
				GoStmt.GoRaw("\thxrt_serializerWriteSerializedEnum(self, current)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case haxe__SerializedClassRef:"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"A\")"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, current.name)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__SerializedClassRef:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"A\")"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteStringToken(self, current.name)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case haxe__SerializedEnumRef:"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"B\")"),
				GoStmt.GoRaw("\thxrt_serializerWriteStringToken(self, current.name)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__SerializedEnumRef:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"B\")"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteStringToken(self, current.name)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__ds__List:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\tif hxrt_serializerTrackRef(self, current) {"),
				GoStmt.GoRaw("\t\t\treturn"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteListToken(self, current.items)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__ds__StringMap:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\tif hxrt_serializerTrackRef(self, current) {"),
				GoStmt.GoRaw("\t\t\treturn"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteStringMapToken(self, current.h)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__ds__IntMap:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\tif hxrt_serializerTrackRef(self, current) {"),
				GoStmt.GoRaw("\t\t\treturn"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteIntMapToken(self, current.h)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *haxe__ds__ObjectMap:"),
				GoStmt.GoRaw("\tif current == nil {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\tif hxrt_serializerTrackRef(self, current) {"),
				GoStmt.GoRaw("\t\t\treturn"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteObjectMapToken(self, current.h)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case int:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case int8:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case int16:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case int32:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case int64:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, current)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case uint:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case uint8:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case uint16:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case uint32:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case uint64:"),
				GoStmt.GoRaw("\thxrt_serializerWriteIntToken(self, int64(current))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case float32:"),
				GoStmt.GoRaw("\tvalue64 := float64(current)"),
				GoStmt.GoRaw("\tif math.IsNaN(value64) {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"k\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif math.IsInf(value64, 1) {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"p\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif math.IsInf(value64, -1) {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"m\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"d\" + strconv.FormatFloat(value64, 'g', -1, 64))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case float64:"),
				GoStmt.GoRaw("\tif math.IsNaN(current) {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"k\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif math.IsInf(current, 1) {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"p\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif math.IsInf(current, -1) {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"m\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"d\" + strconv.FormatFloat(current, 'g', -1, 64))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("ref := reflect.ValueOf(value)"),
				GoStmt.GoRaw("if !ref.IsValid() {"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if hxrt_serializerTrySpecialReflect(self, value) {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch ref.Kind() {"),
				GoStmt.GoRaw("case reflect.Slice, reflect.Array:"),
				GoStmt.GoRaw("\tif ref.Kind() == reflect.Slice && ref.IsNil() {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif hxrt_serializerTrackRef(self, value) {"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"a\")"),
				GoStmt.GoRaw("\tnullRun := 0"),
				GoStmt.GoRaw("\tfor i := 0; i < ref.Len(); i++ {"),
				GoStmt.GoRaw("\t\titem := ref.Index(i).Interface()"),
				GoStmt.GoRaw("\t\tif item == nil {"),
				GoStmt.GoRaw("\t\t\tnullRun++"),
				GoStmt.GoRaw("\t\t\tcontinue"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif nullRun > 1 {"),
				GoStmt.GoRaw("\t\t\thxrt_serializerAppend(self, \"u\" + strconv.Itoa(nullRun))"),
				GoStmt.GoRaw("\t\t} else if nullRun == 1 {"),
				GoStmt.GoRaw("\t\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tnullRun = 0"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteValue(self, item)"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif nullRun > 1 {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"u\" + strconv.Itoa(nullRun))"),
				GoStmt.GoRaw("\t} else if nullRun == 1 {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"h\")"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case reflect.Map:"),
				GoStmt.GoRaw("\tif ref.IsNil() {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif ref.Type().Key().Kind() != reflect.String {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Serializer map keys must be strings\"))"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif hxrt_serializerTrackRef(self, value) {"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"o\")"),
				GoStmt.GoRaw("\tkeys := ref.MapKeys()"),
				GoStmt.GoRaw("\tsortedKeys := make([]string, 0, len(keys))"),
				GoStmt.GoRaw("\tfor _, key := range keys {"),
				GoStmt.GoRaw("\t\tsortedKeys = append(sortedKeys, key.String())"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tsort.Strings(sortedKeys)"),
				GoStmt.GoRaw("\tfor _, key := range sortedKeys {"),
				GoStmt.GoRaw("\t\thxrt_serializerWriteStringToken(self, key)"),
				GoStmt.GoRaw("\t\tvalueRef := ref.MapIndex(reflect.ValueOf(key))"),
				GoStmt.GoRaw("\t\tif valueRef.IsValid() {"),
				GoStmt.GoRaw("\t\t\thxrt_serializerWriteValue(self, valueRef.Interface())"),
				GoStmt.GoRaw("\t\t} else {"),
				GoStmt.GoRaw("\t\t\thxrt_serializerWriteValue(self, nil)"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerAppend(self, \"g\")"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case reflect.Pointer:"),
				GoStmt.GoRaw("\tif ref.IsNil() {"),
				GoStmt.GoRaw("\t\thxrt_serializerAppend(self, \"n\")"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif hxrt_serializerTrackRef(self, value) {"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thxrt_serializerWriteValue(self, ref.Elem().Interface())"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("hxrt.Throw(hxrt.StringFromLiteral(\"Unsupported serializer value type\"))"),
				GoStmt.GoRaw("hxrt_serializerAppend(self, \"n\")")
			]),
			GoDecl.GoStructDecl("haxe__Unserializer", [
				{
					name: "buf",
					typeName: "*string"
				},
				{name: "pos", typeName: "int"},
				{name: "stringCache", typeName: "[]*string"},
				{name: "cache", typeName: "[]any"},
				{name: "resolver", typeName: "any"}
			]),
			GoDecl.GoFuncDecl("New_haxe__Unserializer", null, [{name: "buf", typeName: "*string"}], ["*haxe__Unserializer"], [
				GoStmt.GoRaw("resolver := haxe__Unserializer_DEFAULT_RESOLVER"),
				GoStmt.GoRaw("if resolver == nil {"),
				GoStmt.GoRaw("\tresolver = &haxe__Unserializer__DefaultResolver{}"),
				GoStmt.GoRaw("\thaxe__Unserializer_DEFAULT_RESOLVER = resolver"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoRaw("&haxe__Unserializer{buf: buf, pos: 0, stringCache: []*string{}, cache: []any{}, resolver: resolver}"))
			]),
			GoDecl.GoFuncDecl("setResolver", {
				name: "self",
				typeName: "*haxe__Unserializer"
			}, [{name: "resolver", typeName: "any"}], [], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if resolver == nil {"),
				GoStmt.GoRaw("\tself.resolver = haxe__Unserializer_NULL_RESOLVER"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("self.resolver = resolver")
			]),
			GoDecl.GoFuncDecl("getResolver", {
				name: "self",
				typeName: "*haxe__Unserializer"
			}, [], ["any"], [
				GoStmt.GoRaw("if self == nil {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("return self.resolver")
			]),
			GoDecl.GoFuncDecl("unserialize", {
				name: "self",
				typeName: "*haxe__Unserializer"
			}, [], ["any"], [
				GoStmt.GoIf(GoExpr.GoRaw("self == nil || self.buf == nil"), [GoStmt.GoReturn(GoExpr.GoNil)], null),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("haxe__Unserializer_readValue"), [GoExpr.GoIdent("self")]))
			]),
			GoDecl.GoFuncDecl("haxe__Unserializer_readUInt", null, [
				{
					name: "self",
					typeName: "*haxe__Unserializer"
				}
			], ["int"], [
				GoStmt.GoRaw("raw := *hxrt.StdString(self.buf)"),
				GoStmt.GoRaw("start := self.pos"),
				GoStmt.GoRaw("for self.pos < len(raw) && raw[self.pos] >= '0' && raw[self.pos] <= '9' {"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.pos == start {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized integer\"))"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("parsed, err := strconv.Atoi(raw[start:self.pos])"),
				GoStmt.GoRaw("if err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("parsed"))
			]),
			GoDecl.GoFuncDecl("haxe__Unserializer_readDigits", null, [
				{
					name: "self",
					typeName: "*haxe__Unserializer"
				}
			], ["int"], [
				GoStmt.GoRaw("raw := *hxrt.StdString(self.buf)"),
				GoStmt.GoRaw("start := self.pos"),
				GoStmt.GoRaw("if self.pos < len(raw) && (raw[self.pos] == '-' || raw[self.pos] == '+') {"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("digitStart := self.pos"),
				GoStmt.GoRaw("for self.pos < len(raw) && raw[self.pos] >= '0' && raw[self.pos] <= '9' {"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if self.pos == digitStart {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized integer\"))"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("parsed, err := strconv.Atoi(raw[start:self.pos])"),
				GoStmt.GoRaw("if err != nil {"),
				GoStmt.GoRaw("\thxrt.Throw(err)"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoIdent("parsed"))
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerSetField", null, [
				{
					name: "target",
					typeName: "any"
				},
				{name: "fieldName", typeName: "string"},
				{name: "value", typeName: "any"}
			], [], [
				GoStmt.GoRaw("if target == nil {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("switch obj := target.(type) {"),
				GoStmt.GoRaw("case map[string]any:"),
				GoStmt.GoRaw("\tobj[fieldName] = value"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case map[any]any:"),
				GoStmt.GoRaw("\tobj[fieldName] = value"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *map[string]any:"),
				GoStmt.GoRaw("\tif obj != nil {"),
				GoStmt.GoRaw("\t\t(*obj)[fieldName] = value"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("case *map[any]any:"),
				GoStmt.GoRaw("\tif obj != nil {"),
				GoStmt.GoRaw("\t\t(*obj)[fieldName] = value"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("rv := reflect.ValueOf(target)"),
				GoStmt.GoRaw("if !rv.IsValid() || rv.Kind() != reflect.Pointer || rv.IsNil() {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("elem := rv.Elem()"),
				GoStmt.GoRaw("if !elem.IsValid() || elem.Kind() != reflect.Struct {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("field := elem.FieldByName(fieldName)"),
				GoStmt.GoRaw("if !field.IsValid() {"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("targetField := field"),
				GoStmt.GoRaw("if !targetField.CanSet() {"),
				GoStmt.GoRaw("\tif !targetField.CanAddr() {"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\ttargetField = reflect.NewAt(targetField.Type(), unsafe.Pointer(targetField.UnsafeAddr())).Elem()"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if value == nil {"),
				GoStmt.GoRaw("\ttargetField.Set(reflect.Zero(targetField.Type()))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("incoming := reflect.ValueOf(value)"),
				GoStmt.GoRaw("if incoming.IsValid() && incoming.Type().AssignableTo(targetField.Type()) {"),
				GoStmt.GoRaw("\ttargetField.Set(incoming)"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if incoming.IsValid() && incoming.Type().ConvertibleTo(targetField.Type()) {"),
				GoStmt.GoRaw("\ttargetField.Set(incoming.Convert(targetField.Type()))"),
				GoStmt.GoRaw("\treturn"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("if targetField.Kind() == reflect.Interface && incoming.IsValid() {"),
				GoStmt.GoRaw("\ttargetField.Set(incoming)"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("hxrt_unserializerReadObjectFields", null, [
				{
					name: "self",
					typeName: "*haxe__Unserializer"
				},
				{name: "target", typeName: "any"},
				{name: "invalidMessage", typeName: "string"}
			], [], [
				GoStmt.GoRaw("raw := *hxrt.StdString(self.buf)"),
				GoStmt.GoRaw("for {"),
				GoStmt.GoRaw("\tif self.pos >= len(raw) {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(invalidMessage))"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif raw[self.pos] == 'g' {"),
				GoStmt.GoRaw("\t\tself.pos++"),
				GoStmt.GoRaw("\t\treturn"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tfieldNameAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\tfieldName := *hxrt.StdString(fieldNameAny)"),
				GoStmt.GoRaw("\tfieldValue := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\thxrt_unserializerSetField(target, fieldName, fieldValue)"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("haxe__Unserializer_readHexNibble", null, [
				{
					name: "ch",
					typeName: "byte"
				}
			], ["int"], [
				GoStmt.GoRaw("switch {"),
				GoStmt.GoRaw("case ch >= '0' && ch <= '9':"),
				GoStmt.GoRaw("\treturn int(ch - '0')"),
				GoStmt.GoRaw("case ch >= 'A' && ch <= 'F':"),
				GoStmt.GoRaw("\treturn int(ch-'A') + 10"),
				GoStmt.GoRaw("case ch >= 'a' && ch <= 'f':"),
				GoStmt.GoRaw("\treturn int(ch-'a') + 10"),
				GoStmt.GoRaw("default:"),
				GoStmt.GoRaw("\treturn -1"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("haxe__Unserializer_unescape", null, [
				{
					name: "value",
					typeName: "string"
				}
			], ["*string"], [
				GoStmt.GoRaw("out := make([]byte, 0, len(value))"),
				GoStmt.GoRaw("for i := 0; i < len(value); i++ {"),
				GoStmt.GoRaw("\tif value[i] != '%' {"),
				GoStmt.GoRaw("\t\tout = append(out, value[i])"),
				GoStmt.GoRaw("\t\tcontinue"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif i+2 >= len(value) {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized string escape\"))"),
				GoStmt.GoRaw("\t\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\thigh := haxe__Unserializer_readHexNibble(value[i+1])"),
				GoStmt.GoRaw("\tlow := haxe__Unserializer_readHexNibble(value[i+2])"),
				GoStmt.GoRaw("\tif high < 0 || low < 0 {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized string escape\"))"),
				GoStmt.GoRaw("\t\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tout = append(out, byte((high<<4)|low))"),
				GoStmt.GoRaw("\ti += 2"),
				GoStmt.GoRaw("}"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoIdent("hxrt.StringFromLiteral"), [GoExpr.GoRaw("string(out)")]))
			]),
			GoDecl.GoFuncDecl("haxe__Unserializer_readValue", null, [
				{
					name: "self",
					typeName: "*haxe__Unserializer"
				}
			], ["any"], [
				GoStmt.GoRaw("if self == nil || self.buf == nil {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("raw := *hxrt.StdString(self.buf)"),
				GoStmt.GoRaw("if self.pos >= len(raw) {"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized string\"))"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("token := raw[self.pos]"),
				GoStmt.GoRaw("self.pos++"),
				GoStmt.GoRaw("switch token {"),
				GoStmt.GoRaw("case 'n':"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("case 't':"),
				GoStmt.GoRaw("\treturn true"),
				GoStmt.GoRaw("case 'f':"),
				GoStmt.GoRaw("\treturn false"),
				GoStmt.GoRaw("case 'z':"),
				GoStmt.GoRaw("\treturn 0"),
				GoStmt.GoRaw("case 'k':"),
				GoStmt.GoRaw("\treturn math.NaN()"),
				GoStmt.GoRaw("case 'p':"),
				GoStmt.GoRaw("\treturn math.Inf(1)"),
				GoStmt.GoRaw("case 'm':"),
				GoStmt.GoRaw("\treturn math.Inf(-1)"),
				GoStmt.GoRaw("case 'i':"),
				GoStmt.GoRaw("\tstart := self.pos"),
				GoStmt.GoRaw("\tif self.pos < len(raw) && (raw[self.pos] == '-' || raw[self.pos] == '+') {"),
				GoStmt.GoRaw("\t\tself.pos++"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tfor self.pos < len(raw) && raw[self.pos] >= '0' && raw[self.pos] <= '9' {"),
				GoStmt.GoRaw("\t\tself.pos++"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif self.pos == start || (self.pos == start+1 && (raw[start] == '-' || raw[start] == '+')) {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized integer\"))"),
				GoStmt.GoRaw("\t\treturn 0"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tparsed, err := strconv.Atoi(raw[start:self.pos])"),
				GoStmt.GoRaw("\tif err != nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(err)"),
				GoStmt.GoRaw("\t\treturn 0"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn parsed"),
				GoStmt.GoRaw("case 'd':"),
				GoStmt.GoRaw("\tstart := self.pos"),
				GoStmt.GoRaw("\thasDigit := false"),
				GoStmt.GoRaw("\tfor self.pos < len(raw) {"),
				GoStmt.GoRaw("\t\tch := raw[self.pos]"),
				GoStmt.GoRaw("\t\tif ch >= '0' && ch <= '9' {"),
				GoStmt.GoRaw("\t\t\thasDigit = true"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tcontinue"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif ch == '+' || ch == '-' || ch == '.' || ch == 'e' || ch == 'E' {"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tcontinue"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tbreak"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif !hasDigit {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized float\"))"),
				GoStmt.GoRaw("\t\treturn 0.0"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tparsed, err := strconv.ParseFloat(raw[start:self.pos], 64)"),
				GoStmt.GoRaw("\tif err != nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(err)"),
				GoStmt.GoRaw("\t\treturn 0.0"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn parsed"),
				GoStmt.GoRaw("case 'v':"),
				GoStmt.GoRaw("\tstart := self.pos"),
				GoStmt.GoRaw("\thasDigit := false"),
				GoStmt.GoRaw("\tfor self.pos < len(raw) {"),
				GoStmt.GoRaw("\t\tch := raw[self.pos]"),
				GoStmt.GoRaw("\t\tif ch >= '0' && ch <= '9' {"),
				GoStmt.GoRaw("\t\t\thasDigit = true"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tcontinue"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif ch == '+' || ch == '-' || ch == '.' || ch == 'e' || ch == 'E' {"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tcontinue"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tbreak"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif !hasDigit {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized date\"))"),
				GoStmt.GoRaw("\t\treturn &haxe__SerializedDate{ms: 0}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tparsed, err := strconv.ParseFloat(raw[start:self.pos], 64)"),
				GoStmt.GoRaw("\tif err != nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(err)"),
				GoStmt.GoRaw("\t\treturn &haxe__SerializedDate{ms: 0}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn &haxe__SerializedDate{ms: parsed}"),
				GoStmt.GoRaw("case 's':"),
				GoStmt.GoRaw("\tlength := haxe__Unserializer_readUInt(self)"),
				GoStmt.GoRaw("\tif self.pos >= len(raw) || raw[self.pos] != ':' {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized bytes\"))"),
				GoStmt.GoRaw("\t\treturn &haxe__SerializedBytes{data: []byte{}}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("\tif length < 0 || self.pos+length > len(raw) {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized bytes length\"))"),
				GoStmt.GoRaw("\t\treturn &haxe__SerializedBytes{data: []byte{}}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tencoded := raw[self.pos : self.pos+length]"),
				GoStmt.GoRaw("\tself.pos += length"),
				GoStmt.GoRaw("\tdecoded, err := base64.RawStdEncoding.DecodeString(encoded)"),
				GoStmt.GoRaw("\tif err != nil {"),
				GoStmt.GoRaw("\t\tdecoded, err = base64.StdEncoding.DecodeString(encoded)"),
				GoStmt.GoRaw("\t\tif err != nil {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(err)"),
				GoStmt.GoRaw("\t\t\treturn &haxe__SerializedBytes{data: []byte{}}"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tout := make([]byte, len(decoded))"),
				GoStmt.GoRaw("\tcopy(out, decoded)"),
				GoStmt.GoRaw("\treturn &haxe__SerializedBytes{data: out}"),
				GoStmt.GoRaw("case 'y':"),
				GoStmt.GoRaw("\tlength := haxe__Unserializer_readUInt(self)"),
				GoStmt.GoRaw("\tif self.pos >= len(raw) || raw[self.pos] != ':' {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized string\"))"),
				GoStmt.GoRaw("\t\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("\tif length < 0 || self.pos+length > len(raw) {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized string length\"))"),
				GoStmt.GoRaw("\t\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tdecoded := haxe__Unserializer_unescape(raw[self.pos : self.pos+length])"),
				GoStmt.GoRaw("\tself.pos += length"),
				GoStmt.GoRaw("\tself.stringCache = append(self.stringCache, decoded)"),
				GoStmt.GoRaw("\treturn decoded"),
				GoStmt.GoRaw("case 'R':"),
				GoStmt.GoRaw("\tindex := haxe__Unserializer_readUInt(self)"),
				GoStmt.GoRaw("\tif index < 0 || index >= len(self.stringCache) {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid string reference\"))"),
				GoStmt.GoRaw("\t\treturn hxrt.StringFromLiteral(\"\")"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn self.stringCache[index]"),
				GoStmt.GoRaw("case 'x':"),
				GoStmt.GoRaw("\thxrt.Throw(haxe__Unserializer_readValue(self))"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("case 'l':"),
				GoStmt.GoRaw("\tlist := New_haxe__ds__List()"),
				GoStmt.GoRaw("\tcacheIndex := len(self.cache)"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, list)"),
				GoStmt.GoRaw("\tfor {"),
				GoStmt.GoRaw("\t\tif self.pos >= len(raw) {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized list\"))"),
				GoStmt.GoRaw("\t\t\treturn list"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif raw[self.pos] == 'h' {"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tbreak"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tlist.items = append(list.items, haxe__Unserializer_readValue(self))"),
				GoStmt.GoRaw("\t\tlist.length = len(list.items)"),
				GoStmt.GoRaw("\t\tself.cache[cacheIndex] = list"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn list"),
				GoStmt.GoRaw("case 'b':"),
				GoStmt.GoRaw("\tstringMap := New_haxe__ds__StringMap()"),
				GoStmt.GoRaw("\tcacheIndex := len(self.cache)"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, stringMap)"),
				GoStmt.GoRaw("\tfor {"),
				GoStmt.GoRaw("\t\tif self.pos >= len(raw) {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized StringMap\"))"),
				GoStmt.GoRaw("\t\t\treturn stringMap"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif raw[self.pos] == 'h' {"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tbreak"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tkeyAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\t\tkey := *hxrt.StdString(keyAny)"),
				GoStmt.GoRaw("\t\tstringMap.h[key] = haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\t\tself.cache[cacheIndex] = stringMap"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn stringMap"),
				GoStmt.GoRaw("case 'q':"),
				GoStmt.GoRaw("\tintMap := New_haxe__ds__IntMap()"),
				GoStmt.GoRaw("\tcacheIndex := len(self.cache)"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, intMap)"),
				GoStmt.GoRaw("\tfor {"),
				GoStmt.GoRaw("\t\tif self.pos >= len(raw) {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized IntMap\"))"),
				GoStmt.GoRaw("\t\t\treturn intMap"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif raw[self.pos] == 'h' {"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tbreak"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif raw[self.pos] != ':' {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized IntMap format\"))"),
				GoStmt.GoRaw("\t\t\treturn intMap"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tself.pos++"),
				GoStmt.GoRaw("\t\tkey := haxe__Unserializer_readDigits(self)"),
				GoStmt.GoRaw("\t\tintMap.h[key] = haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\t\tself.cache[cacheIndex] = intMap"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn intMap"),
				GoStmt.GoRaw("case 'M':"),
				GoStmt.GoRaw("\tobjectMap := New_haxe__ds__ObjectMap()"),
				GoStmt.GoRaw("\tcacheIndex := len(self.cache)"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, objectMap)"),
				GoStmt.GoRaw("\tfor {"),
				GoStmt.GoRaw("\t\tif self.pos >= len(raw) {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized ObjectMap\"))"),
				GoStmt.GoRaw("\t\t\treturn objectMap"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif raw[self.pos] == 'h' {"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tbreak"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tkey := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\t\tobjectMap.h[key] = haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\t\tself.cache[cacheIndex] = objectMap"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn objectMap"),
				GoStmt.GoRaw("case 'a':"),
				GoStmt.GoRaw("\tarr := make([]any, 0)"),
				GoStmt.GoRaw("\tcacheIndex := len(self.cache)"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, arr)"),
				GoStmt.GoRaw("\tfor {"),
				GoStmt.GoRaw("\t\tif self.pos >= len(raw) {"),
				GoStmt.GoRaw("\t\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized array\"))"),
				GoStmt.GoRaw("\t\t\treturn arr"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif raw[self.pos] == 'h' {"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tbreak"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tif raw[self.pos] == 'u' {"),
				GoStmt.GoRaw("\t\t\tself.pos++"),
				GoStmt.GoRaw("\t\t\tskip := haxe__Unserializer_readUInt(self)"),
				GoStmt.GoRaw("\t\t\tfor i := 0; i < skip; i++ {"),
				GoStmt.GoRaw("\t\t\t\tarr = append(arr, nil)"),
				GoStmt.GoRaw("\t\t\t}"),
				GoStmt.GoRaw("\t\t\tself.cache[cacheIndex] = arr"),
				GoStmt.GoRaw("\t\t\tcontinue"),
				GoStmt.GoRaw("\t\t}"),
				GoStmt.GoRaw("\t\tarr = append(arr, haxe__Unserializer_readValue(self))"),
				GoStmt.GoRaw("\t\tself.cache[cacheIndex] = arr"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn arr"),
				GoStmt.GoRaw("case 'o':"),
				GoStmt.GoRaw("\tobj := map[string]any{}"),
				GoStmt.GoRaw("\tcacheIndex := len(self.cache)"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, obj)"),
				GoStmt.GoRaw("\thxrt_unserializerReadObjectFields(self, obj, \"Invalid serialized object\")"),
				GoStmt.GoRaw("\tself.cache[cacheIndex] = obj"),
				GoStmt.GoRaw("\treturn obj"),
				GoStmt.GoRaw("case 'C':"),
				GoStmt.GoRaw("\tclassNameAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\trequestedName := *hxrt.StdString(classNameAny)"),
				GoStmt.GoRaw("\tresolvedClass := hxrt_unserializerResolveClass(self, requestedName)"),
				GoStmt.GoRaw("\tif resolvedClass == nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Class not found \" + requestedName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tclassName, ok := hxrt_unserializerResolvedClassName(resolvedClass)"),
				GoStmt.GoRaw("\tif !ok || className == \"\" {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Class not found \" + requestedName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tinstance, ok := hxrt_unserializerCreateClassInstance(className)"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Class not found \" + className))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tcacheIndex := len(self.cache)"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, instance)"),
				GoStmt.GoRaw("\tif custom, ok := instance.(interface{ hxUnserialize(*haxe__Unserializer) }); ok {"),
				GoStmt.GoRaw("\t\tcustom.hxUnserialize(self)"),
				GoStmt.GoRaw("\t} else if custom, ok := instance.(interface{ HxUnserialize(*haxe__Unserializer) }); ok {"),
				GoStmt.GoRaw("\t\tcustom.HxUnserialize(self)"),
				GoStmt.GoRaw("\t} else {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid custom data\"))"),
				GoStmt.GoRaw("\t\treturn instance"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif self.pos >= len(raw) || raw[self.pos] != 'g' {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid custom data\"))"),
				GoStmt.GoRaw("\t\treturn instance"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("\tself.cache[cacheIndex] = instance"),
				GoStmt.GoRaw("\treturn instance"),
				GoStmt.GoRaw("case 'A':"),
				GoStmt.GoRaw("\tnameAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\trequestedName := *hxrt.StdString(nameAny)"),
				GoStmt.GoRaw("\tresolvedClass := hxrt_unserializerResolveClass(self, requestedName)"),
				GoStmt.GoRaw("\tif resolvedClass == nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Class not found \" + requestedName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn resolvedClass"),
				GoStmt.GoRaw("case 'B':"),
				GoStmt.GoRaw("\tnameAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\trequestedName := *hxrt.StdString(nameAny)"),
				GoStmt.GoRaw("\tresolvedEnum := hxrt_unserializerResolveEnum(self, requestedName)"),
				GoStmt.GoRaw("\tif resolvedEnum == nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Enum not found \" + requestedName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn resolvedEnum"),
				GoStmt.GoRaw("case 'c':"),
				GoStmt.GoRaw("\tclassNameAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\trequestedName := *hxrt.StdString(classNameAny)"),
				GoStmt.GoRaw("\tresolvedClass := hxrt_unserializerResolveClass(self, requestedName)"),
				GoStmt.GoRaw("\tif resolvedClass == nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Class not found \" + requestedName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tclassName, ok := hxrt_unserializerResolvedClassName(resolvedClass)"),
				GoStmt.GoRaw("\tif !ok || className == \"\" {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Class not found \" + requestedName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tinstance, ok := hxrt_unserializerCreateClassInstance(className)"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Class not found \" + className))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tcacheIndex := len(self.cache)"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, instance)"),
				GoStmt.GoRaw("\thxrt_unserializerReadObjectFields(self, instance, \"Invalid serialized class\")"),
				GoStmt.GoRaw("\tself.cache[cacheIndex] = instance"),
				GoStmt.GoRaw("\treturn instance"),
				GoStmt.GoRaw("case 'j':"),
				GoStmt.GoRaw("\tenumNameAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\trequestedEnumName := *hxrt.StdString(enumNameAny)"),
				GoStmt.GoRaw("\tresolvedEnum := hxrt_unserializerResolveEnum(self, requestedEnumName)"),
				GoStmt.GoRaw("\tif resolvedEnum == nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Enum not found \" + requestedEnumName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tenumName, ok := hxrt_unserializerResolvedEnumName(resolvedEnum)"),
				GoStmt.GoRaw("\tif !ok || enumName == \"\" {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Enum not found \" + requestedEnumName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tif self.pos >= len(raw) || raw[self.pos] != ':' {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized enum index\"))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("\tenumIndex := haxe__Unserializer_readDigits(self)"),
				GoStmt.GoRaw("\tif self.pos >= len(raw) || raw[self.pos] != ':' {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized enum\"))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("\targCount := haxe__Unserializer_readUInt(self)"),
				GoStmt.GoRaw("\tif argCount < 0 {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized enum arity\"))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\targs := make([]any, 0, argCount)"),
				GoStmt.GoRaw("\tfor i := 0; i < argCount; i++ {"),
				GoStmt.GoRaw("\t\targs = append(args, haxe__Unserializer_readValue(self))"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tenumValue, ok := hxrt_unserializerCreateEnumInstance(enumName, \"\", enumIndex, true, args)"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Unknown enum index \" + enumName + \"@\" + strconv.Itoa(enumIndex)))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, enumValue)"),
				GoStmt.GoRaw("\treturn enumValue"),
				GoStmt.GoRaw("case 'w':"),
				GoStmt.GoRaw("\tenumNameAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\tconstructorAny := haxe__Unserializer_readValue(self)"),
				GoStmt.GoRaw("\trequestedEnumName := *hxrt.StdString(enumNameAny)"),
				GoStmt.GoRaw("\tresolvedEnum := hxrt_unserializerResolveEnum(self, requestedEnumName)"),
				GoStmt.GoRaw("\tif resolvedEnum == nil {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Enum not found \" + requestedEnumName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tenumName, ok := hxrt_unserializerResolvedEnumName(resolvedEnum)"),
				GoStmt.GoRaw("\tif !ok || enumName == \"\" {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Enum not found \" + requestedEnumName))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tconstructorName := *hxrt.StdString(constructorAny)"),
				GoStmt.GoRaw("\tif self.pos >= len(raw) || raw[self.pos] != ':' {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized enum\"))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.pos++"),
				GoStmt.GoRaw("\targCount := haxe__Unserializer_readUInt(self)"),
				GoStmt.GoRaw("\tif argCount < 0 {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized enum arity\"))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\targs := make([]any, 0, argCount)"),
				GoStmt.GoRaw("\tfor i := 0; i < argCount; i++ {"),
				GoStmt.GoRaw("\t\targs = append(args, haxe__Unserializer_readValue(self))"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tenumValue, ok := hxrt_unserializerCreateEnumInstance(enumName, constructorName, 0, false, args)"),
				GoStmt.GoRaw("\tif !ok {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized enum\"))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\tself.cache = append(self.cache, enumValue)"),
				GoStmt.GoRaw("\treturn enumValue"),
				GoStmt.GoRaw("case 'r':"),
				GoStmt.GoRaw("\tindex := haxe__Unserializer_readUInt(self)"),
				GoStmt.GoRaw("\tif index < 0 || index >= len(self.cache) {"),
				GoStmt.GoRaw("\t\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid object reference\"))"),
				GoStmt.GoRaw("\t\treturn nil"),
				GoStmt.GoRaw("\t}"),
				GoStmt.GoRaw("\treturn self.cache[index]"),
				GoStmt.GoRaw("default:"),
				GoStmt.GoRaw("\thxrt.Throw(hxrt.StringFromLiteral(\"Invalid serialized token\"))"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}")
			]),
			GoDecl.GoFuncDecl("haxe__Unserializer_run", null, [
				{
					name: "source",
					typeName: "*string"
				}
			], ["any"], [
				GoStmt.GoRaw("if source == nil {"),
				GoStmt.GoRaw("\treturn nil"),
				GoStmt.GoRaw("}"),
				GoStmt.GoRaw("decoder := New_haxe__Unserializer(source)"),
				GoStmt.GoReturn(GoExpr.GoCall(GoExpr.GoSelector(GoExpr.GoIdent("decoder"), "unserialize"), []))
			])
		];
	}
}
#end
