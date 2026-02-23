package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.Expr.Binop;
import haxe.macro.Expr.Unop;
import reflaxe.go.ast.GoAST.GoExpr;

class GoExprOperatorOps {
	public static function lowerHaxeInt32BinopExpr(op:Binop, leftExpr:GoExpr, rightExpr:GoExpr):GoExpr {
		var leftInt32 = GoExpr.GoCall(GoExpr.GoIdent("hxrt.Int32Wrap"), [leftExpr]);
		return switch (op) {
			case OpUShr:
				var shifted = GoExpr.GoBinary(">>", GoExpr.GoCall(GoExpr.GoIdent("uint32"), [leftInt32]), GoExpr.GoCall(GoExpr.GoIdent("uint"), [rightExpr]));
				wrapInt32Expr(GoExpr.GoCall(GoExpr.GoIdent("int32"), [shifted]));
			case OpShl:
				wrapInt32Expr(GoExpr.GoBinary("<<", leftInt32, GoExpr.GoCall(GoExpr.GoIdent("uint"), [rightExpr])));
			case OpShr:
				wrapInt32Expr(GoExpr.GoBinary(">>", leftInt32, GoExpr.GoCall(GoExpr.GoIdent("uint"), [rightExpr])));
			case OpAdd | OpSub | OpMult | OpMod | OpAnd | OpOr | OpXor:
				var rightInt32 = GoExpr.GoCall(GoExpr.GoIdent("hxrt.Int32Wrap"), [rightExpr]);
				wrapInt32Expr(GoExpr.GoBinary(binopSymbol(op), leftInt32, rightInt32));
			case _:
				GoExpr.GoBinary(binopSymbol(op), leftExpr, rightExpr);
		};
	}

	public static function wrapInt32Expr(expr:GoExpr):GoExpr {
		return GoExpr.GoCall(GoExpr.GoIdent("int"), [GoExpr.GoCall(GoExpr.GoIdent("int32"), [expr])]);
	}

	public static function floatOperandExpr(expr:GoExpr, shouldKeepFloat:Bool):GoExpr {
		return shouldKeepFloat ? expr : GoExpr.GoCall(GoExpr.GoIdent("float64"), [expr]);
	}

	public static function unitStepExpr(target:GoExpr, opSymbol:String, wrapInt32:Bool):GoExpr {
		var stepped = GoExpr.GoBinary(opSymbol, target, GoExpr.GoIntLiteral(1));
		return wrapInt32 ? wrapInt32Expr(stepped) : stepped;
	}

	public static function binopSymbol(op:Binop):String {
		return switch (op) {
			case OpAdd:
				"+";
			case OpMult:
				"*";
			case OpDiv:
				"/";
			case OpSub:
				"-";
			case OpMod:
				"%";
			case OpEq:
				"==";
			case OpNotEq:
				"!=";
			case OpGt:
				">";
			case OpGte:
				">=";
			case OpLt:
				"<";
			case OpLte:
				"<=";
			case OpBoolAnd:
				"&&";
			case OpBoolOr:
				"||";
			case OpAnd:
				"&";
			case OpOr:
				"|";
			case OpXor:
				"^";
			case OpShl:
				"<<";
			case OpShr:
				">>";
			case OpUShr:
				">>";
			case OpAssign:
				Context.fatalError("Assignment is handled at statement level", Context.currentPos());
			case _:
				Context.fatalError("Unsupported binary operator", Context.currentPos());
		};
	}

	public static function unopSymbol(op:Unop):String {
		return switch (op) {
			case OpNot:
				"!";
			case OpNeg:
				"-";
			case OpNegBits:
				"^";
			case OpIncrement:
				Context.fatalError("Increment operator is not supported yet", Context.currentPos());
			case OpDecrement:
				Context.fatalError("Decrement operator is not supported yet", Context.currentPos());
			case _:
				Context.fatalError("Unsupported unary operator", Context.currentPos());
		};
	}
}
#end
