package hxrt

import (
	"bufio"
	"os/exec"
)

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
