package reflaxe.go.ast.transformers.registry;

#if macro
import haxe.macro.Context;
#end
import reflaxe.go.CompilationContext;
import reflaxe.go.ast.GoAST.GoFile;

interface IGoASTPass {
  public function getName():String;
  public function getDependencies():Array<String>;
  public function run(file:GoFile, context:CompilationContext):GoFile;
}

class RegistryCore {
  public static function validateAndOrder(passes:Array<IGoASTPass>):Array<IGoASTPass> {
    var byName = new Map<String, IGoASTPass>();
    for (pass in passes) {
      var name = pass.getName();
      if (byName.exists(name)) {
        fail('Duplicate Go AST pass name: "' + name + '"');
      }
      byName.set(name, pass);
    }

    for (pass in passes) {
      for (dependency in pass.getDependencies()) {
        if (!byName.exists(dependency)) {
          fail('Missing Go AST pass dependency "' + dependency + '" required by "' + pass.getName() + '"');
        }
      }
    }

    var state = new Map<String, Int>();
    var visiting = new Array<String>();
    var ordered = new Array<IGoASTPass>();

    function visit(name:String):Void {
      var currentState = state.exists(name) ? state.get(name) : 0;
      if (currentState == 2) {
        return;
      }
      if (currentState == 1) {
        var cycleStart = visiting.indexOf(name);
        var cyclePath = cycleStart >= 0 ? visiting.slice(cycleStart) : [name];
        cyclePath.push(name);
        fail('Go AST pass dependency cycle: ' + cyclePath.join(" -> "));
      }

      state.set(name, 1);
      visiting.push(name);

      var pass = byName.get(name);
      for (dependency in pass.getDependencies()) {
        visit(dependency);
      }

      visiting.pop();
      state.set(name, 2);
      ordered.push(pass);
    }

    for (pass in passes) {
      visit(pass.getName());
    }

    return ordered;
  }

  static function fail(message:String):Void {
    #if macro
    Context.fatalError(message, Context.currentPos());
    #else
    throw message;
    #end
  }
}
