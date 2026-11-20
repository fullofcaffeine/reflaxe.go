import Shapes;

/**
	The stability test recompiles this module with two edits:
	`recovery_names_reorder` swaps the first use of two helpers, and
	`recovery_names_edit` demands helpers that no other module uses.
**/
class Edited {
	public static function run(first:Dynamic, second:Dynamic, callback:Dynamic):Int {
		#if recovery_names_reorder
		var circle:Circle = second;
		var square:Square = first;
		#else
		var square:Square = first;
		var circle:Circle = second;
		#end
		var total = square.size + circle.size;
		#if recovery_names_edit
		var measure:String->Int = callback;
		if (Std.isOfType(first, Shape))
			total += measure("shape");
		#end
		return total;
	}
}
