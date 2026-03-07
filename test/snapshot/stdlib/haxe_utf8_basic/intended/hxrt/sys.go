package hxrt

import (
	"os"
	"runtime"
	"strings"
)

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

func SysGetEnv(key *string) *string {
	if key == nil {
		return nil
	}
	value, ok := os.LookupEnv(*key)
	if !ok {
		return nil
	}
	return StringFromLiteral(value)
}

func SysPutEnv(key *string, value *string) {
	if key == nil {
		return
	}
	if value == nil {
		_ = os.Unsetenv(*key)
		return
	}
	_ = os.Setenv(*key, *value)
}

func SysEnvironment() map[string]string {
	out := map[string]string{}
	for _, entry := range os.Environ() {
		key, value, ok := strings.Cut(entry, "=")
		if !ok {
			out[entry] = ""
			continue
		}
		out[key] = value
	}
	return out
}

func SysSystemName() *string {
	switch runtime.GOOS {
	case "darwin":
		return StringFromLiteral("Mac")
	case "linux":
		return StringFromLiteral("Linux")
	case "windows":
		return StringFromLiteral("Windows")
	case "freebsd", "openbsd", "netbsd", "dragonfly":
		return StringFromLiteral("BSD")
	default:
		return StringFromLiteral(runtime.GOOS)
	}
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
