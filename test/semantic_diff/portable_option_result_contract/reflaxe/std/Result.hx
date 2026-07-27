package reflaxe.std;

/** Portable success-or-error value that preserves its independent error type. */
enum Result<T, E> {
	Ok(value:T);
	Err(error:E);
}
