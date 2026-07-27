package reflaxe.std;

/** Fixture-supplied portable presence value for fallback evidence. */
enum Option<T> {
	Some(value:T);
	None;
}
