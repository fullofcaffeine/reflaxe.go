package reflaxe.go.ast;

import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoFile;
import reflaxe.go.ast.GoAST.GoInterfaceMethod;
import reflaxe.go.ast.GoAST.GoParam;
import reflaxe.go.ast.GoAST.GoSelectCase;
import reflaxe.go.ast.GoAST.GoSelectClause;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoAST.GoSwitchCase;
import reflaxe.go.ast.GoAST.GoTypeSwitchCase;

class GoASTPrinter {
	public static function printFile(file:GoFile):String {
		var out = new StringBuf();
		out.add("package ");
		out.add(file.packageName.value());
		out.add("\n\n");

		if (file.imports.length == 1) {
			out.add('import "');
			out.add(file.imports[0].value());
			out.add('"\n\n');
		} else if (file.imports.length > 1) {
			out.add("import (\n");
			for (path in file.imports) {
				out.add('\t"');
				out.add(path.value());
				out.add('"\n');
			}
			out.add(")\n\n");
		}

		var isFirst = true;
		for (decl in file.decls) {
			if (!isFirst) {
				out.add("\n");
			}
			isFirst = false;
			out.add(printDecl(decl));
		}

		return out.toString();
	}

	/**
		Print a single Go expression for target-code injection placeholder expansion.

		Why:
		`untyped __go__("...{0}...", expr)` needs a stable way to render already-lowered
		Go expressions back into source text without re-printing an entire file.

		What:
		This is the minimal expression-only printer used when expanding `__go__`
		placeholder arguments.

		How:
		Delegate to the normal expression printer so injection expansion and normal output
		share one rendering implementation.
	**/
	public static function printExprForInjection(expr:GoExpr):String {
		return printExpr(expr);
	}

	static function printDecl(decl:GoDecl):String {
		return switch (decl) {
			case GoInterfaceDecl(name, methods):
				var out = new StringBuf();
				out.add("type ");
				out.add(name);
				out.add(" interface {\n");
				for (method in methods) {
					out.add("\t");
					out.add(printInterfaceMethod(method));
					out.add("\n");
				}
				out.add("}\n");
				out.toString();
			case GoStructDecl(name, fields):
				var out = new StringBuf();
				out.add("type ");
				out.add(name);
				out.add(" struct {\n");
				for (field in fields) {
					out.add("\t");
					if (field.name != "") {
						out.add(field.name);
						out.add(" ");
					}
					out.add(field.typeName.render());
					out.add("\n");
				}
				out.add("}\n");
				out.toString();
			case GoGlobalVarDecl(name, typeName, value):
				if (value == null) {
					"var " + name + " " + typeName.render() + "\n";
				} else {
					"var " + name + " " + typeName.render() + " = " + printExpr(value) + "\n";
				}
			case GoFuncDecl(name, receiver, params, results, body):
				var out = new StringBuf();
				out.add("func ");
				if (receiver != null) {
					out.add("(");
					out.add(receiver.name);
					out.add(" ");
					out.add(receiver.typeName.render());
					out.add(") ");
				}
				out.add(name);
				out.add("(");
				out.add(printParams(params));
				out.add(")");

				if (results.length == 1) {
					out.add(" ");
					out.add(results[0].render());
				} else if (results.length > 1) {
					out.add(" (");
					out.add(renderTypes(results));
					out.add(")");
				}

				out.add(" {\n");
				for (bodyStmt in body) {
					out.add("\t");
					out.add(printStmt(bodyStmt));
					out.add("\n");
				}
				out.add("}\n");
				out.toString();
		}
	}

	static function printParams(params:Array<GoParam>):String {
		var rendered = new Array<String>();
		for (param in params) {
			rendered.push(param.name + " " + param.typeName.render());
		}
		return rendered.join(", ");
	}

	static function printInterfaceMethod(method:GoInterfaceMethod):String {
		var out = new StringBuf();
		out.add(method.name);
		out.add("(");
		out.add(printParams(method.params));
		out.add(")");
		if (method.results.length == 1) {
			out.add(" ");
			out.add(method.results[0].render());
		} else if (method.results.length > 1) {
			out.add(" (");
			out.add(renderTypes(method.results));
			out.add(")");
		}
		return out.toString();
	}

	static function printStmt(stmt:GoStmt):String {
		return switch (stmt) {
			case GoVarDecl(name, typeName, value, useShort):
				if (useShort && value != null) {
					name + " := " + printExpr(value);
				} else if (value == null) {
					typeName == null ? "var " + name : "var " + name + " " + typeName.render();
				} else if (typeName == null) {
					"var " + name + " = " + printExpr(value);
				} else {
					"var " + name + " " + typeName.render() + " = " + printExpr(value);
				}
			case GoMultiAssign(names, value, useShort):
				names.join(", ") + (useShort ? " := " : " = ") + printExpr(value);
			case GoAssign(left, right, op):
				printExpr(left) + " " + assignmentToken(op) + " " + printExpr(right);
			case GoIncDec(target, op):
				printExpr(target) + op.token();
			case GoExprStmt(expr): printExpr(expr);
			case GoGoStmt(call):
				"go " + printExpr(call);
			case GoDeferStmt(call):
				"defer " + printExpr(call);
			case GoSendStmt(channel, value):
				printExpr(channel) + " <- " + printExpr(value);
			case GoRaw(code): code;
			case GoWhile(cond, body):
				var out = new StringBuf();
				out.add("for ");
				out.add(printExpr(cond));
				out.add(" {\n");
				for (bodyStmt in body) {
					out.add("\t");
					out.add(printStmt(bodyStmt));
					out.add("\n");
				}
				out.add("}");
				out.toString();
			case GoFor(initializer, condition, post, body):
				var out = new StringBuf();
				out.add("for");
				if (initializer == null && post == null) {
					if (condition != null) {
						out.add(" ");
						out.add(printExpr(condition));
					}
				} else {
					out.add(" ");
					if (initializer != null) {
						out.add(printSimpleStmt(initializer, false));
					}
					out.add("; ");
					if (condition != null) {
						out.add(printExpr(condition));
					}
					out.add("; ");
					if (post != null) {
						out.add(printSimpleStmt(post, true));
					}
				}
				out.add(" {\n");
				for (inner in body) {
					out.add("\t");
					out.add(printStmt(inner));
					out.add("\n");
				}
				out.add("}");
				out.toString();
			case GoLabeled(label, stmt):
				label + ":\n" + printStmt(stmt);
			case GoRangeStmt(keyName, valueName, source, useShort, body):
				var out = new StringBuf();
				out.add("for ");
				var hasBindings = keyName != null || valueName != null;
				if (hasBindings) {
					if (keyName != null) {
						out.add(keyName);
					}
					if (valueName != null) {
						if (keyName == null) {
							out.add("_");
						}
						out.add(", ");
						out.add(valueName);
					}
					out.add(useShort ? " := range " : " = range ");
				} else {
					out.add("range ");
				}
				out.add(printExpr(source));
				out.add(" {\n");
				for (bodyStmt in body) {
					out.add("\t");
					out.add(printStmt(bodyStmt));
					out.add("\n");
				}
				out.add("}");
				out.toString();
			case GoIf(cond, thenBody, elseBody):
				var out = new StringBuf();
				out.add("if ");
				out.add(printExpr(cond));
				out.add(" {\n");
				for (bodyStmt in thenBody) {
					out.add("\t");
					out.add(printStmt(bodyStmt));
					out.add("\n");
				}
				out.add("}");
				if (elseBody != null) {
					out.add(" else {\n");
					for (bodyStmt in elseBody) {
						out.add("\t");
						out.add(printStmt(bodyStmt));
						out.add("\n");
					}
					out.add("}");
				}
				out.toString();
			case GoSwitch(value, cases, defaultBody):
				var out = new StringBuf();
				out.add("switch ");
				out.add(printExpr(value));
				out.add(" {\n");
				for (caseEntry in cases) {
					out.add("\t");
					out.add(printSwitchCase(caseEntry));
					out.add("\n");
				}
				if (defaultBody != null) {
					out.add("\tdefault:\n");
					for (bodyStmt in defaultBody) {
						out.add("\t\t");
						out.add(printStmt(bodyStmt));
						out.add("\n");
					}
				}
				out.add("}");
				out.toString();
			case GoTypeSwitch(value, bindingName, cases, defaultBody):
				var out = new StringBuf();
				out.add("switch ");
				if (bindingName != null) {
					out.add(bindingName);
					out.add(" := ");
				}
				out.add(printExpr(value));
				out.add(".(type) {\n");
				for (caseEntry in cases) {
					out.add("\t");
					out.add(printTypeSwitchCase(caseEntry));
					out.add("\n");
				}
				if (defaultBody != null) {
					out.add("\tdefault:\n");
					for (bodyStmt in defaultBody) {
						out.add("\t\t");
						out.add(printStmt(bodyStmt));
						out.add("\n");
					}
				}
				out.add("}");
				out.toString();
			case GoSelect(cases):
				var out = new StringBuf();
				out.add("select {\n");
				for (caseEntry in cases) {
					out.add("\t");
					out.add(printSelectCase(caseEntry));
					out.add("\n");
				}
				out.add("}");
				out.toString();
			case GoBreak(label):
				label == null ? "break" : "break " + label;
			case GoContinue:
				"continue";
			case GoReturn(expr): expr == null ? "return" : "return " + printExpr(expr);
		}
	}

	static function printSwitchCase(caseEntry:GoSwitchCase):String {
		var out = new StringBuf();
		out.add("case ");
		out.add([for (value in caseEntry.values) printExpr(value)].join(", "));
		out.add(":\n");
		for (stmt in caseEntry.body) {
			out.add("\t\t");
			out.add(printStmt(stmt));
			out.add("\n");
		}
		return StringTools.rtrim(out.toString());
	}

	static function printTypeSwitchCase(caseEntry:GoTypeSwitchCase):String {
		var out = new StringBuf();
		out.add("case ");
		out.add(caseEntry.typeName.render());
		out.add(":\n");
		for (stmt in caseEntry.body) {
			out.add("\t\t");
			out.add(printStmt(stmt));
			out.add("\n");
		}
		return StringTools.rtrim(out.toString());
	}

	static function printSelectCase(caseEntry:GoSelectCase):String {
		var out = new StringBuf();
		switch (caseEntry.clause) {
			case GoSelectClause.GoSelectSend(channel, value):
				out.add("case ");
				out.add(printExpr(channel));
				out.add(" <- ");
				out.add(printExpr(value));
				out.add(":\n");
			case GoSelectClause.GoSelectRecv(recv):
				out.add("case ");
				out.add(printExpr(recv));
				out.add(":\n");
			case GoSelectClause.GoSelectRecvAssign(target, recv, useShort):
				out.add("case ");
				out.add(printExpr(target));
				out.add(useShort ? " := " : " = ");
				out.add(printExpr(recv));
				out.add(":\n");
			case GoSelectClause.GoSelectRecvAssignOk(target, okTarget, recv, useShort):
				out.add("case ");
				out.add(printExpr(target));
				out.add(", ");
				out.add(printExpr(okTarget));
				out.add(useShort ? " := " : " = ");
				out.add(printExpr(recv));
				out.add(":\n");
			case GoSelectClause.GoSelectDefault:
				out.add("default:\n");
		}
		for (stmt in caseEntry.body) {
			out.add("\t\t");
			out.add(printStmt(stmt));
			out.add("\n");
		}
		return StringTools.rtrim(out.toString());
	}

	static function printExpr(expr:GoExpr):String {
		return switch (expr) {
			case GoIdent(name): name;
			case GoIntLiteral(value): Std.string(value);
			case GoFloatLiteral(value): value;
			case GoBoolLiteral(value): value ? "true" : "false";
			case GoStringLiteral(value): '"' + escapeString(value) + '"';
			case GoNil: "nil";
			case GoSelector(target, field): printExpr(target) + "." + field;
			case GoIndex(target, index): printExpr(target) + "[" + printExpr(index) + "]";
			case GoSlice(target, start, end):
				var startCode = start == null ? "" : printExpr(start);
				var endCode = end == null ? "" : printExpr(end);
				printExpr(target) + "[" + startCode + ":" + endCode + "]";
			case GoCompositeLiteral(typeName, elements):
				if (!typeName.supportsCompositeLiteral()) {
					throw 'Invalid Go composite literal type "' + typeName.render() + '"';
				}
				typeName.render() + "{" + [for (element in elements) printCompositeElement(element)].join(", ") + "}";
			case GoMakeSlice(elementType, length, capacity):
				"make([]"
				+ elementType.render()
				+ ", "
				+ printExpr(length)
				+ (capacity == null ? "" : ", " + printExpr(capacity))
				+ ")";
			case GoFuncLiteral(params, results, body):
				var out = new StringBuf();
				out.add("func(");
				out.add(printParams(params));
				out.add(")");
				if (results.length == 1) {
					out.add(" ");
					out.add(results[0].render());
				} else if (results.length > 1) {
					out.add(" (");
					out.add(renderTypes(results));
					out.add(")");
				}
				out.add(" {\n");
				for (stmt in body) {
					out.add("\t");
					out.add(printStmt(stmt));
					out.add("\n");
				}
				out.add("}");
				out.toString();
			case GoRaw(code): code;
			case GoTypeAssert(inner, typeName):
				printExpr(inner) + ".(" + typeName.render() + ")";
			case GoRecvExpr(channel):
				"<-" + printExpr(channel);
			case GoUnary(op, inner): op.token() + printExpr(inner);
			case GoBinary(op, left, right): "(" + printExpr(left) + " " + op.token() + " " + printExpr(right) + ")";
			case GoCall(callee, args):
				var renderedArgs = [for (arg in args) printExpr(arg)].join(", ");
				printExpr(callee) + "(" + renderedArgs + ")";
		}
	}

	static function assignmentToken(op:Null<GoAssignmentOperator>):String {
		return op == null ? GoAssignmentOperator.Assign.token() : op.token();
	}

	/**
		What: Render the closed statement subset admitted by a classic `for` clause.
		Why: Initializer and post positions share most syntax, but Go specifically
		forbids a short declaration in the post position.
		How: Render each typed variant and apply the position-sensitive rejection
		before generated Go reaches `gofmt` or `go test`.
	**/
	static function printSimpleStmt(stmt:GoSimpleStmt, isPost:Bool):String {
		return switch (stmt) {
			case GoSimpleDeclare(name, value):
				if (isPost) {
					throw "Invalid Go for post statement: short declarations are not permitted";
				}
				name + " := " + printExpr(value);
			case GoSimpleAssign(left, right, op):
				printExpr(left) + " " + assignmentToken(op) + " " + printExpr(right);
			case GoSimpleIncDec(target, op):
				printExpr(target) + op.token();
			case GoSimpleExpr(expr):
				printExpr(expr);
			case GoSimpleSend(channel, value):
				printExpr(channel) + " <- " + printExpr(value);
		};
	}

	/**
		What: Render one positional, expression-keyed, or named-field literal entry.
		Why: Expression keys require traversal while struct field names require Go
		identifier validation; flattening both to text loses that distinction.
		How: Delegate expression text to `printExpr` and validate only the field-name
		variant before adding its colon.
	**/
	static function printCompositeElement(element:GoCompositeElement):String {
		return switch (element) {
			case GoCompositeValue(value): printExpr(value);
			case GoCompositeKeyValue(key, value): printExpr(key) + ": " + printExpr(value);
			case GoCompositeField(fieldName, value):
				if (!GoPackageName.isIdentifier(fieldName)) {
					throw 'Invalid Go composite field name "' + fieldName + '"';
				}
				fieldName + ": " + printExpr(value);
		};
	}

	static function renderTypes(types:Array<GoType>):String {
		return [for (type in types) type.render()].join(", ");
	}

	static function escapeString(value:String):String {
		return value.split("\\")
			.join("\\\\")
			.split("\"")
			.join("\\\"")
			.split("\n")
			.join("\\n")
			.split("\r")
			.join("\\r")
			.split("\t")
			.join("\\t");
	}
}
