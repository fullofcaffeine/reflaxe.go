package reflaxe.go;

#if macro
import reflaxe.go.ast.GoAST.GoDecl;
import reflaxe.go.ast.GoAST.GoExpr;
import reflaxe.go.ast.GoAST.GoFile;
import reflaxe.go.ast.GoAST.GoStmt;
import reflaxe.go.ast.GoASTPrinter;
#end

typedef GoGeneratedFile = {
  final relativePath:String;
  final contents:String;
}

class GoCompiler {
  public function new() {}

  #if macro
  public function compileModule():Array<GoGeneratedFile> {
    var mainFile:GoFile = {
      packageName: "main",
      imports: ["snapshot/hxrt"],
      decls: [
        GoDecl.GoFuncDecl("main", [], [], [
          GoStmt.GoExprStmt(GoExpr.GoCall(
            GoExpr.GoIdent("hxrt.Println"),
            [
              GoExpr.GoCall(
                GoExpr.GoIdent("hxrt.StringFromLiteral"),
                [GoExpr.GoStringLiteral("hi")]
              )
            ]
          ))
        ])
      ]
    };

    return [
      {
        relativePath: "main.go",
        contents: GoASTPrinter.printFile(mainFile)
      }
    ];
  }
  #end
}
