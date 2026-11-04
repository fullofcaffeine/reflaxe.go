package hxrt

func unwrapHaxeException(recovered any) (any, bool) {
	switch v := recovered.(type) {
	case HaxeException:
		return v.Value, true
	case *HaxeException:
		return v.Value, true
	default:
		return nil, false
	}
}

// UnwrapException converts only the carrier created by Throw. A foreign Go
// panic is a backend/native failure, not a Haxe value, so it resumes unwinding
// unchanged instead of entering a Haxe catch clause.
func UnwrapException(recovered any) any {
	value, ok := unwrapHaxeException(recovered)
	if !ok {
		panic(recovered)
	}
	return value
}

func TryCatch(tryFn func(), catchFn func(any)) {
	defer func() {
		if recovered := recover(); recovered != nil {
			catchFn(UnwrapException(recovered))
		}
	}()
	tryFn()
}
