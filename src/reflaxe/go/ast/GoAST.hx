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

typedef GoInterfaceMethod = {
  final name:String;
  final params:Array<GoParam>;
  final results:Array<String>;
}

typedef GoSwitchCase = {
  final values:Array<GoExpr>;
  final body:Array<GoStmt>;
}

enum GoDecl {
  GoInterfaceDecl(name:String, methods:Array<GoInterfaceMethod>);
  GoStructDecl(name:String, fields:Array<GoParam>);
  GoGlobalVarDecl(name:String, typeName:String, value:Null<GoExpr>);
  GoFuncDecl(name:String, receiver:Null<GoParam>, params:Array<GoParam>, results:Array<String>, body:Array<GoStmt>);
}

enum GoStmt {
  GoVarDecl(name:String, typeName:Null<String>, value:Null<GoExpr>, useShort:Bool);
  GoAssign(left:GoExpr, right:GoExpr);
  GoExprStmt(expr:GoExpr);
  GoRaw(code:String);
  GoWhile(cond:GoExpr, body:Array<GoStmt>);
  GoIf(cond:GoExpr, thenBody:Array<GoStmt>, elseBody:Null<Array<GoStmt>>);
  GoSwitch(value:GoExpr, cases:Array<GoSwitchCase>, defaultBody:Null<Array<GoStmt>>);
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
  GoIndex(target:GoExpr, index:GoExpr);
  GoSlice(target:GoExpr, start:Null<GoExpr>, end:Null<GoExpr>);
  GoArrayLiteral(elementType:String, elements:Array<GoExpr>);
  GoFuncLiteral(params:Array<GoParam>, results:Array<String>, body:Array<GoStmt>);
  GoRaw(code:String);
  GoTypeAssert(expr:GoExpr, typeName:String);
  GoUnary(op:String, expr:GoExpr);
  GoBinary(op:String, left:GoExpr, right:GoExpr);
  GoCall(callee:GoExpr, args:Array<GoExpr>);
}
