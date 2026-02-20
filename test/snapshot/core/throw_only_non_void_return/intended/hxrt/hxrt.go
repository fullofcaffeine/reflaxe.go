package hxrt

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"os/exec"
	"reflect"
	"sync"
	"sync/atomic"
)

func StringFromLiteral(value string) *string {
	copy := value
	return &copy
}

func StdString(value any) *string {
	switch v := value.(type) {
	case nil:
		return StringFromLiteral("null")
	case *string:
		if v == nil {
			return StringFromLiteral("null")
		}
		return v
	case string:
		return StringFromLiteral(v)
	default:
		return StringFromLiteral(fmt.Sprint(v))
	}
}

func StringSlice(values []*string) []string {
	out := make([]string, len(values))
	for i := 0; i < len(values); i++ {
		out[i] = *StdString(values[i])
	}
	return out
}

func StringConcatAny(left any, right any) *string {
	l := StdString(left)
	r := StdString(right)
	return StringFromLiteral(*l + *r)
}

func StringEqualAny(left any, right any) bool {
	l := StdString(left)
	r := StdString(right)
	return *l == *r
}

func stringValueOrNullToken(value *string) string {
	if value == nil {
		return "null"
	}
	return *value
}

func StringConcatStringPtr(left *string, right *string) *string {
	return StringFromLiteral(stringValueOrNullToken(left) + stringValueOrNullToken(right))
}

func StringEqualStringPtr(left *string, right *string) bool {
	return stringValueOrNullToken(left) == stringValueOrNullToken(right)
}

func StringLength(value any) int {
	runes := []rune(*StdString(value))
	return len(runes)
}

func StringCharAt(value any, index int) *string {
	runes := []rune(*StdString(value))
	if index < 0 || index >= len(runes) {
		return StringFromLiteral("")
	}
	return StringFromLiteral(string(runes[index]))
}

func StringCharCodeAt(value any, index int) int {
	runes := []rune(*StdString(value))
	if index < 0 || index >= len(runes) {
		return -1
	}
	return int(runes[index])
}

func StringSubstring(value any, start int, end int) *string {
	runes := []rune(*StdString(value))
	if start < 0 {
		start = 0
	}
	if end < 0 {
		end = 0
	}
	if start > len(runes) {
		start = len(runes)
	}
	if end > len(runes) {
		end = len(runes)
	}
	if end < start {
		end = start
	}
	return StringFromLiteral(string(runes[start:end]))
}

func FloatMod(a float64, b float64) float64 {
	return math.Mod(a, b)
}

func Int32Wrap(value int) int32 {
	return int32(value)
}

type AtomicIntCell struct {
	value atomic.Int64
}

func AtomicIntNew(value int) any {
	cell := &AtomicIntCell{}
	cell.value.Store(int64(value))
	return cell
}

func AtomicIntLoad(cell any) int {
	typed := cell.(*AtomicIntCell)
	return int(typed.value.Load())
}

func AtomicIntStore(cell any, value int) int {
	typed := cell.(*AtomicIntCell)
	typed.value.Store(int64(value))
	return value
}

func AtomicIntExchange(cell any, value int) int {
	typed := cell.(*AtomicIntCell)
	return int(typed.value.Swap(int64(value)))
}

func AtomicIntCompareExchange(cell any, expected int, replacement int) int {
	typed := cell.(*AtomicIntCell)
	expectedValue := int64(expected)
	replacementValue := int64(replacement)
	for {
		previous := typed.value.Load()
		if previous != expectedValue {
			return int(previous)
		}
		if typed.value.CompareAndSwap(previous, replacementValue) {
			return int(previous)
		}
	}
}

func AtomicIntAdd(cell any, value int) int {
	typed := cell.(*AtomicIntCell)
	delta := int64(value)
	return int(typed.value.Add(delta) - delta)
}

func AtomicIntSub(cell any, value int) int {
	typed := cell.(*AtomicIntCell)
	delta := int64(value)
	return int(typed.value.Add(-delta) + delta)
}

func AtomicIntAnd(cell any, value int) int {
	typed := cell.(*AtomicIntCell)
	mask := int64(value)
	for {
		previous := typed.value.Load()
		next := previous & mask
		if typed.value.CompareAndSwap(previous, next) {
			return int(previous)
		}
	}
}

func AtomicIntOr(cell any, value int) int {
	typed := cell.(*AtomicIntCell)
	mask := int64(value)
	for {
		previous := typed.value.Load()
		next := previous | mask
		if typed.value.CompareAndSwap(previous, next) {
			return int(previous)
		}
	}
}

func AtomicIntXor(cell any, value int) int {
	typed := cell.(*AtomicIntCell)
	mask := int64(value)
	for {
		previous := typed.value.Load()
		next := previous ^ mask
		if typed.value.CompareAndSwap(previous, next) {
			return int(previous)
		}
	}
}

type AtomicObjectCell struct {
	mu    sync.Mutex
	value any
}

func AtomicObjectNew(value any) any {
	return &AtomicObjectCell{value: value}
}

func AtomicObjectLoad(cell any) any {
	typed := cell.(*AtomicObjectCell)
	typed.mu.Lock()
	defer typed.mu.Unlock()
	return typed.value
}

func AtomicObjectStore(cell any, value any) any {
	typed := cell.(*AtomicObjectCell)
	typed.mu.Lock()
	defer typed.mu.Unlock()
	typed.value = value
	return value
}

func AtomicObjectExchange(cell any, value any) any {
	typed := cell.(*AtomicObjectCell)
	typed.mu.Lock()
	defer typed.mu.Unlock()
	previous := typed.value
	typed.value = value
	return previous
}

func AtomicObjectCompareExchange(cell any, expected any, replacement any) any {
	typed := cell.(*AtomicObjectCell)
	typed.mu.Lock()
	defer typed.mu.Unlock()
	previous := typed.value
	if atomicReferenceEqual(previous, expected) {
		typed.value = replacement
	}
	return previous
}

func atomicReferenceEqual(left any, right any) bool {
	leftValue := reflect.ValueOf(left)
	rightValue := reflect.ValueOf(right)
	if !leftValue.IsValid() || !rightValue.IsValid() {
		return !leftValue.IsValid() && !rightValue.IsValid()
	}
	if leftValue.Type() != rightValue.Type() {
		return false
	}

	switch leftValue.Kind() {
	case reflect.Interface:
		if leftValue.IsNil() || rightValue.IsNil() {
			return leftValue.IsNil() && rightValue.IsNil()
		}
		return atomicReferenceEqual(leftValue.Elem().Interface(), rightValue.Elem().Interface())
	case reflect.Ptr, reflect.UnsafePointer, reflect.Map, reflect.Slice, reflect.Func, reflect.Chan:
		if leftValue.IsNil() || rightValue.IsNil() {
			return leftValue.IsNil() && rightValue.IsNil()
		}
		return leftValue.Pointer() == rightValue.Pointer()
	default:
		if leftValue.Type().Comparable() {
			return left == right
		}
		return false
	}
}

func Println(value any) {
	fmt.Println(*StdString(value))
}

type HaxeException struct {
	Value any
}

type ExceptionValue struct {
	Value   any
	Message *string
}

func Throw(value any) {
	panic(HaxeException{Value: value})
}

func UnwrapException(recovered any) any {
	switch v := recovered.(type) {
	case HaxeException:
		return v.Value
	case *HaxeException:
		return v.Value
	default:
		return v
	}
}

func TryCatch(tryFn func(), catchFn func(any)) {
	defer func() {
		if recovered := recover(); recovered != nil {
			catchFn(UnwrapException(recovered))
		}
	}()
	tryFn()
}

func ExceptionCaught(value any) *ExceptionValue {
	switch v := value.(type) {
	case *ExceptionValue:
		return v
	case ExceptionValue:
		copy := v
		return &copy
	default:
		return &ExceptionValue{
			Value:   value,
			Message: StdString(value),
		}
	}
}

func ExceptionThrown(value any) any {
	return ExceptionCaught(value).Value
}

func ExceptionMessage(value any) *string {
	return ExceptionCaught(value).Message
}

func JsonParse(source *string) any {
	if source == nil {
		return nil
	}

	var decoded any
	if err := json.Unmarshal([]byte(*source), &decoded); err != nil {
		return nil
	}
	return decoded
}

func JsonStringify(value any) *string {
	encoded, err := json.Marshal(value)
	if err != nil {
		return StringFromLiteral("null")
	}
	return StringFromLiteral(string(encoded))
}

func SysGetCwd() *string {
	cwd, err := os.Getwd()
	if err != nil {
		return StringFromLiteral("")
	}
	return StringFromLiteral(cwd)
}

func SysArgs() []*string {
	args := os.Args
	if len(args) <= 1 {
		return []*string{}
	}
	out := make([]*string, 0, len(args)-1)
	for _, arg := range args[1:] {
		out = append(out, StringFromLiteral(arg))
	}
	return out
}

func FileSaveContent(path *string, content *string) {
	_ = os.WriteFile(*StdString(path), []byte(*StdString(content)), 0o644)
}

func FileGetContent(path *string) *string {
	raw, err := os.ReadFile(*StdString(path))
	if err != nil {
		return StringFromLiteral("")
	}
	return StringFromLiteral(string(raw))
}

type ProcessOutput struct {
	scanner *bufio.Scanner
}

func (self *ProcessOutput) ReadLine() *string {
	if self == nil || self.scanner == nil {
		return StringFromLiteral("")
	}
	if self.scanner.Scan() {
		return StringFromLiteral(self.scanner.Text())
	}
	return StringFromLiteral("")
}

type Process struct {
	cmd    *exec.Cmd
	stdout *ProcessOutput
}

func NewProcess(command *string, args []*string) *Process {
	cmd := exec.Command(*StdString(command), StringSlice(args)...)
	stdout := &ProcessOutput{}
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return &Process{cmd: cmd, stdout: stdout}
	}
	if err := cmd.Start(); err != nil {
		return &Process{cmd: cmd, stdout: stdout}
	}
	stdout.scanner = bufio.NewScanner(stdoutPipe)
	return &Process{cmd: cmd, stdout: stdout}
}

func (self *Process) Stdout() *ProcessOutput {
	if self == nil {
		return &ProcessOutput{}
	}
	if self.stdout == nil {
		self.stdout = &ProcessOutput{}
	}
	return self.stdout
}

func (self *Process) Close() {
	if self == nil || self.cmd == nil {
		return
	}
	if self.cmd.Process != nil {
		_ = self.cmd.Process.Kill()
	}
	_ = self.cmd.Wait()
}

func BytesFromString(value *string) []int {
	if value == nil {
		return []int{}
	}

	raw := []byte(*value)
	out := make([]int, len(raw))
	for i := 0; i < len(raw); i++ {
		out[i] = int(raw[i])
	}
	return out
}

func BytesToString(values []int) *string {
	raw := make([]byte, len(values))
	for i := 0; i < len(values); i++ {
		raw[i] = byte(values[i])
	}
	return StringFromLiteral(string(raw))
}

func BytesClone(values []int) []int {
	if len(values) == 0 {
		return []int{}
	}
	out := make([]int, len(values))
	copy(out, values)
	return out
}
