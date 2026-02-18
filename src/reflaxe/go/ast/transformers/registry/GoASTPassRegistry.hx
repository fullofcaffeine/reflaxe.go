package reflaxe.go.ast.transformers.registry;

#if macro
import haxe.macro.Context;
#end
import reflaxe.go.ast.transformers.passes.CollectImportsPass;
import reflaxe.go.ast.transformers.passes.CyclicAlphaPass;
import reflaxe.go.ast.transformers.passes.CyclicBetaPass;
import reflaxe.go.ast.transformers.passes.NormalizeNamesPass;
import reflaxe.go.ast.transformers.registry.groups.GranularBundle;
import reflaxe.go.ast.transformers.registry.groups.LeanBundle;
import reflaxe.go.ast.transformers.registry.RegistryCore.IGoASTPass;

class GoASTPassRegistry {
  static inline final GRANULAR_DEFINE = "go_granular_pass_registry";
  static inline final TEST_DEFINE = "reflaxe_go_test_registry_case";

  public static function resolve():Array<IGoASTPass> {
    var passes = #if macro
      Context.defined(GRANULAR_DEFINE) ? GranularBundle.build() : LeanBundle.build();
    #else
      LeanBundle.build();
    #end

    #if macro
    var testCase = Context.definedValue(TEST_DEFINE);
    if (testCase == null || testCase == "") {
      return passes;
    }

    return switch (testCase) {
      case "duplicate":
        [
          new NormalizeNamesPass(),
          new NormalizeNamesPass()
        ];
      case "missing_dep":
        [new CollectImportsPass()];
      case "cycle":
        [
          new CyclicAlphaPass(),
          new CyclicBetaPass()
        ];
      case _:
        Context.fatalError('Unknown Go AST registry test case "' + testCase + '"', Context.currentPos());
        passes;
    };
    #else
    return passes;
    #end
  }
}
