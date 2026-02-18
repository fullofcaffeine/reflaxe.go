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
  GoExprStmt(expr:GoExpr);
  GoReturn(expr:Null<GoExpr>);
}

enum GoExpr {
  GoIdent(name:String);
  GoStringLiteral(value:String);
  GoCall(callee:GoExpr, args:Array<GoExpr>);
}
