package hxrt

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
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
