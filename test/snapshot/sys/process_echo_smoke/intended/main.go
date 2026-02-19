package main

import "snapshot/hxrt"

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
	impl *hxrt.ProcessOutput
}

type sys__io__Process struct {
	impl   *hxrt.Process
	stdout *sys__io__ProcessOutput
}

func Sys_getCwd() *string {
	return hxrt.SysGetCwd()
}

func Sys_args() []*string {
	return hxrt.SysArgs()
}

func sys__io__File_saveContent(path *string, content *string) {
	hxrt.FileSaveContent(path, content)
}

func sys__io__File_getContent(path *string) *string {
	return hxrt.FileGetContent(path)
}

func New_sys__io__Process(command *string, args []*string) *sys__io__Process {
	impl := hxrt.NewProcess(command, args)
	stdout := &sys__io__ProcessOutput{}
	if impl != nil {
		stdout.impl = impl.Stdout()
	}
	return &sys__io__Process{impl: impl, stdout: stdout}
}

func (self *sys__io__ProcessOutput) readLine() *string {
	if self == nil || self.impl == nil {
		return hxrt.StringFromLiteral("")
	}
	return self.impl.ReadLine()
}

func (self *sys__io__Process) close() {
	if self == nil || self.impl == nil {
		return
	}
	self.impl.Close()
}
