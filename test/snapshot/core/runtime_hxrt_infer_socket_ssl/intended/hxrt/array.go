package hxrt

// Array is the shared runtime identity for a portable Haxe Array<T>.
//
// What: Keep one mutable, nil-capable sequence behind a pointer that generated
// assignments, fields, calls, callbacks, and erased generic APIs can all share.
// Why: A Go slice value copies its length/capacity header, so an append or shrink
// through one alias is invisible through another. Haxe Array is instead a mutable
// reference object, and sparse writes also preserve null holes even for Array<Int>.
// How: Store elements in one []any owned by this explicit runtime boundary. The
// Haxe compiler retains source element types and performs typed coercion only at
// operations that require it; erased storage does not leak into native go.Slice
// APIs or explicit @:goNative boundaries.
type Array struct {
	values []any
}

func init() {
	// Optional runtime files compile independently of this carrier. Registering the
	// constructor here lets those files create portable Arrays once feature planning
	// has selected array.go, without forcing every generated program to copy it.
	portableArrayFromValues = func(values []any) any {
		return ArrayFromValues(values)
	}
}

// NewArray constructs a distinct portable Array identity from source values.
func NewArray(values ...any) *Array {
	return ArrayFromValues(values)
}

// ArrayFromValues copies an erased value slice into a new portable Array identity.
// It is used when a source-owned generic helper already produced []any storage.
func ArrayFromValues(values []any) *Array {
	copied := make([]any, len(values))
	copy(copied, values)
	return &Array{values: copied}
}

// Len returns the current logical length shared by every alias.
func (array *Array) Len() int {
	return len(array.values)
}

// Get implements Haxe's nil-returning out-of-range indexed read.
func (array *Array) Get(index int) any {
	if index < 0 || index >= len(array.values) {
		return nil
	}
	return array.values[index]
}

// Set implements Haxe indexed assignment, including sparse null-filled growth.
// The assigned value is returned so expression-form assignment can preserve the
// Haxe result without evaluating the right-hand side twice.
func (array *Array) Set(index int, value any) any {
	if index < 0 {
		return value
	}
	if index >= len(array.values) {
		array.values = append(array.values, make([]any, index-len(array.values)+1)...)
	}
	array.values[index] = value
	return value
}

// Push appends values to the shared sequence and returns the new Haxe length.
func (array *Array) Push(values ...any) int {
	array.values = append(array.values, values...)
	return len(array.values)
}

// Pop removes and returns the last value, or nil when the Array is empty.
func (array *Array) Pop() any {
	length := len(array.values)
	if length == 0 {
		return nil
	}
	last := length - 1
	value := array.values[last]
	array.values[last] = nil
	array.values = array.values[:last]
	return value
}

// Insert applies Haxe's negative-position and clamped oversized-position rules.
func (array *Array) Insert(position int, value any) {
	length := len(array.values)
	if position < 0 {
		position = length + position
		if position < 0 {
			position = 0
		}
	}
	if position > length {
		position = length
	}
	array.values = append(array.values, nil)
	copy(array.values[position+1:], array.values[position:])
	array.values[position] = value
}

// RemoveAt removes one known in-range slot while preserving the Array identity.
// Equality and first-match policy stay in typed compiler lowering.
func (array *Array) RemoveAt(index int) {
	if index < 0 || index >= len(array.values) {
		return
	}
	last := len(array.values) - 1
	copy(array.values[index:], array.values[index+1:])
	array.values[last] = nil
	array.values = array.values[:last]
}

// SetLength truncates or null-extends the same shared Array object.
func (array *Array) SetLength(length int) {
	if length < 0 {
		length = 0
	}
	if length <= len(array.values) {
		for index := length; index < len(array.values); index++ {
			array.values[index] = nil
		}
		array.values = array.values[:length]
		return
	}
	array.values = append(array.values, make([]any, length-len(array.values))...)
}

// Copy creates a distinct Array identity with the same current element values.
func (array *Array) Copy() *Array {
	return ArrayFromValues(array.values)
}

// Values exposes the live element slice to narrow compiler/runtime bridges such
// as StringJoinAny. Callers may read or replace existing slots but must use Array
// methods for length changes so aliases keep one authoritative header.
func (array *Array) Values() []any {
	return array.values
}

// ValuesCopy returns detached erased storage for an explicit native boundary.
// Mutating the returned slice cannot change this portable Array's element slots.
func (array *Array) ValuesCopy() []any {
	copied := make([]any, len(array.values))
	copy(copied, array.values)
	return copied
}
