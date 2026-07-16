package hxrt

// IntMapCell owns native integer-keyed storage for staged haxe.ds.IntMap.
// Public collection behavior remains in Haxe source; order records the first
// insertion of each live key so iterator snapshots are deterministic.
type IntMapCell struct {
	values map[int]any
	order  []int
}

func IntMapNew() *IntMapCell {
	return &IntMapCell{values: make(map[int]any)}
}

func IntMapSet(cell *IntMapCell, key int, value any) {
	if _, exists := cell.values[key]; !exists {
		cell.order = append(cell.order, key)
	}
	cell.values[key] = value
}

func IntMapGet(cell *IntMapCell, key int) any {
	return cell.values[key]
}

func IntMapExists(cell *IntMapCell, key int) bool {
	_, exists := cell.values[key]
	return exists
}

func IntMapRemove(cell *IntMapCell, key int) bool {
	if _, exists := cell.values[key]; !exists {
		return false
	}
	delete(cell.values, key)
	for index, orderedKey := range cell.order {
		if orderedKey == key {
			cell.order = append(cell.order[:index], cell.order[index+1:]...)
			break
		}
	}
	return true
}

func IntMapKeys(cell *IntMapCell) []int {
	return append([]int{}, cell.order...)
}

func IntMapClear(cell *IntMapCell) {
	cell.values = make(map[int]any)
	cell.order = nil
}

// IntMapSnapshot is a narrow serializer bridge retained until Serializer itself
// moves to staged source. Returning a copy prevents serializer code from
// observing or mutating the live map carrier.
func IntMapSnapshot(cell *IntMapCell) map[int]any {
	out := make(map[int]any, len(cell.values))
	for key, value := range cell.values {
		out[key] = value
	}
	return out
}
