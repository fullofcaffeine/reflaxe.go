package main

import "snapshot/hxrt"

func main() {
	path := hxrt.StringFromLiteral("./tmp_sys_file_smoke.txt")
	_ = path
	sys__io__File_saveContent(path, hxrt.StringFromLiteral("hello"))
	content := sys__io__File_getContent(path)
	_ = content
	hxrt.Println(content)
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
