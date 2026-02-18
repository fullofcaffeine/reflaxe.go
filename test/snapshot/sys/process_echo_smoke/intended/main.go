package main

import (
	"bufio"
	"os"
	"os/exec"
	"snapshot/hxrt"
)

func main() {
	p := New_sys__io__Process(hxrt.StringFromLiteral("echo"), []*string{hxrt.StringFromLiteral("hi")})
	_ = p
	line := p.stdout.readLine()
	_ = line
	hxrt.Println(line)
	p.close()
}

type Sys struct {
}

type sys__io__File struct {
}

type sys__io__ProcessOutput struct {
	scanner *bufio.Scanner
}

type sys__io__Process struct {
	cmd    *exec.Cmd
	stdout *sys__io__ProcessOutput
}

func Sys_getCwd() *string {
	cwd, err := os.Getwd()
	if err != nil {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.StringFromLiteral(cwd)
}

func Sys_args() []*string {
	args := os.Args
	if len(args) <= 1 {
		return []*string{}
	}
	out := make([]*string, 0, len(args)-1)
	for _, arg := range args[1:] {
		out = append(out, hxrt.StringFromLiteral(arg))
	}
	return out
}

func sys__io__File_saveContent(path *string, content *string) {
	_ = os.WriteFile(*hxrt.StdString(path), []byte(*hxrt.StdString(content)), 0o644)
}

func sys__io__File_getContent(path *string) *string {
	raw, err := os.ReadFile(*hxrt.StdString(path))
	if err != nil {
		return hxrt.StringFromLiteral("")
	}
	return hxrt.StringFromLiteral(string(raw))
}

func New_sys__io__Process(command *string, args []*string) *sys__io__Process {
	cmd := exec.Command(*hxrt.StdString(command), hxrt.StringSlice(args)...)
	stdoutPipe, _ := cmd.StdoutPipe()
	_ = cmd.Start()
	scanner := bufio.NewScanner(stdoutPipe)
	return &sys__io__Process{cmd: cmd, stdout: &sys__io__ProcessOutput{scanner: scanner}}
}

func (self *sys__io__ProcessOutput) readLine() *string {
	if self == nil || self.scanner == nil {
		return hxrt.StringFromLiteral("")
	}
	if self.scanner.Scan() {
		return hxrt.StringFromLiteral(self.scanner.Text())
	}
	return hxrt.StringFromLiteral("")
}

func (self *sys__io__Process) close() {
	if self == nil || self.cmd == nil {
		return
	}
	if self.cmd.Process != nil {
		_ = self.cmd.Process.Kill()
	}
	_ = self.cmd.Wait()
}
