package main

func SyscallOpen_install(value SyscallOpenPort) bool {
	SyscallOpen_port = value
	return true
}

func SyscallOpen_openReadOnly(path *string, perm uint32) bool {
	var selected SyscallOpenPort = SyscallOpen_port
	return ((selected == nil) || selected.openReadOnly(path, perm))
}

var SyscallOpen_port SyscallOpenPort = nil

type SyscallOpenPort interface {
	openReadOnly(path *string, perm uint32) bool
}
