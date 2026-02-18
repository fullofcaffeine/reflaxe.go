package reflaxe.go.ast;

typedef GoFile = {
  final packageName:String;
  final imports:Array<String>;
  final decls:Array<GoDecl>;
}

typedef GoParam = {
  final name:String;
  final typeName:String;
}

enum GoDecl {
  GoFuncDecl(name:String, params:Array<GoParam>, results:Array<String>, body:Array<GoStmt>);
}

enum GoStmt {
  GoVarDecl(name:String, typeName:Null<String>, value:Null<GoExpr>, useShort:Bool);
  GoAssign(left:GoExpr, right:GoExpr);
  GoExprStmt(expr:GoExpr);
  GoIf(cond:GoExpr, thenBody:Array<GoStmt>, elseBody:Null<Array<GoStmt>>);
  GoReturn(expr:Null<GoExpr>);
}

enum GoExpr {
  GoIdent(name:String);
  GoIntLiteral(value:Int);
  GoFloatLiteral(value:String);
  GoBoolLiteral(value:Bool);
  GoStringLiteral(value:String);
  GoNil;
  GoSelector(target:GoExpr, field:String);
  GoUnary(op:String, expr:GoExpr);
  GoBinary(op:String, left:GoExpr, right:GoExpr);
  GoCall(callee:GoExpr, args:Array<GoExpr>);
}
