package hxrt

import "os"

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
