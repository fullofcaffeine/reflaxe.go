package reflaxe.std;

/** Fixture-supplied portable result value for fallback evidence. */
enum Result<T, E> {
	Ok(value:T);
	Err(error:E);
}
