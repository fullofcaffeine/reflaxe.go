package reflaxe.std;

/**
	Fixture-supplied portable success-or-error contract.

	Why / What / How
	- `Ok(value)` and `Err(error)` preserve independent `T` and `E` parameters.
	- `E` is not implicitly replaced with `go.Error` or Go's `error` interface.
	- This local definition exercises compiler admission without claiming package
	  ownership for the future standalone `reflaxe.std` release.
**/
enum Result<T, E> {
	Ok(value:T);
	Err(error:E);
}
