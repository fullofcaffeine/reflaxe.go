package reflaxe.std;

/** Portable presence value used as supplied-package contract evidence. */
enum Option<T> {
	Some(value:T);
	None;
}
