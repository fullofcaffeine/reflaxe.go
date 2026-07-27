package reflaxe.std;

/** Fixture-supplied portable result used to prove the native boundary. */
enum Result<T, E> {
	Ok(value:T);
	Err(error:E);
}
