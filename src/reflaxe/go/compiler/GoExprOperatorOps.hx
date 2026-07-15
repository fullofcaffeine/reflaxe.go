package reflaxe.go.compiler;

#if macro
import haxe.macro.Context;
import haxe.macro.Expr.Binop;
import haxe.macro.Expr.Unop;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoBinaryOperator;
import reflaxe.go.ast.GoUnaryOperator;

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

	public static function unitStepExpr(target:GoExpr, opSymbol:GoBinaryOperator, wrapInt32:Bool):GoExpr {
		var stepped = GoExpr.GoBinary(opSymbol, target, GoExpr.GoIntLiteral(1));
		return wrapInt32 ? wrapInt32Expr(stepped) : stepped;
	}

	public static function binopSymbol(op:Binop):GoBinaryOperator {
		return switch (op) {
			case OpAdd:
				GoBinaryOperator.Add;
			case OpMult:
				GoBinaryOperator.Multiply;
			case OpDiv:
				GoBinaryOperator.Divide;
			case OpSub:
				GoBinaryOperator.Subtract;
			case OpMod:
				GoBinaryOperator.Remainder;
			case OpEq:
				GoBinaryOperator.Equal;
			case OpNotEq:
				GoBinaryOperator.NotEqual;
			case OpGt:
				GoBinaryOperator.GreaterThan;
			case OpGte:
				GoBinaryOperator.GreaterThanOrEqual;
			case OpLt:
				GoBinaryOperator.LessThan;
			case OpLte:
				GoBinaryOperator.LessThanOrEqual;
			case OpBoolAnd:
				GoBinaryOperator.LogicalAnd;
			case OpBoolOr:
				GoBinaryOperator.LogicalOr;
			case OpAnd:
				GoBinaryOperator.BitwiseAnd;
			case OpOr:
				GoBinaryOperator.BitwiseOr;
			case OpXor:
				GoBinaryOperator.BitwiseXor;
			case OpShl:
				GoBinaryOperator.ShiftLeft;
			case OpShr:
				GoBinaryOperator.ShiftRight;
			case OpUShr:
				GoBinaryOperator.ShiftRight;
			case OpAssign:
				Context.fatalError("Assignment is handled at statement level", Context.currentPos());
			case _:
				Context.fatalError("Unsupported binary operator", Context.currentPos());
		};
	}

	public static function unopSymbol(op:Unop):GoUnaryOperator {
		return switch (op) {
			case OpNot:
				GoUnaryOperator.LogicalNot;
			case OpNeg:
				GoUnaryOperator.Negate;
			case OpNegBits:
				GoUnaryOperator.BitwiseNot;
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
