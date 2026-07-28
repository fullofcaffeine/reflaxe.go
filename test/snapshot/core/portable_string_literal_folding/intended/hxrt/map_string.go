package hxrt

// StringMapCell owns native string-keyed storage for staged haxe.ds.StringMap.
// Public collection behavior remains in Haxe source; order records the first
// insertion of each live key so iterator snapshots are deterministic.
type StringMapCell struct {
	values map[string]any
	order  []string
}

func StringMapNew() *StringMapCell {
	return &StringMapCell{values: make(map[string]any)}
}

func stringMapKey(key *string) string {
	return *StdString(key)
}

func StringMapSet(cell *StringMapCell, key *string, value any) {
	resolvedKey := stringMapKey(key)
	if _, exists := cell.values[resolvedKey]; !exists {
		cell.order = append(cell.order, resolvedKey)
	}
	cell.values[resolvedKey] = value
}

func StringMapGet(cell *StringMapCell, key *string) any {
	return cell.values[stringMapKey(key)]
}

func StringMapExists(cell *StringMapCell, key *string) bool {
	_, exists := cell.values[stringMapKey(key)]
	return exists
}

func StringMapRemove(cell *StringMapCell, key *string) bool {
	resolvedKey := stringMapKey(key)
	if _, exists := cell.values[resolvedKey]; !exists {
		return false
	}
	delete(cell.values, resolvedKey)
	for index, orderedKey := range cell.order {
		if orderedKey == resolvedKey {
			cell.order = append(cell.order[:index], cell.order[index+1:]...)
			break
		}
	}
	return true
}

func StringMapKeys(cell *StringMapCell) []*string {
	out := make([]*string, 0, len(cell.order))
	for _, key := range cell.order {
		out = append(out, StringFromLiteral(key))
	}
	return out
}

func StringMapClear(cell *StringMapCell) {
	cell.values = make(map[string]any)
	cell.order = nil
}
