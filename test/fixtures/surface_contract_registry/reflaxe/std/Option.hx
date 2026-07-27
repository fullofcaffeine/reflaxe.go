package reflaxe.std;

/**
	Fixture-supplied portable presence contract.

	Why / What / How
	- `Some(value)` and `None` are distinct source cases, even when `value` is null.
	- The real user-facing package remains external to the compiler package.
	- This local definition lets the registry test exact typed identity without
	  falsely claiming that `reflaxe.go` publishes `reflaxe.std`.
**/
enum Option<T> {
	Some(value:T);
	None;
}
