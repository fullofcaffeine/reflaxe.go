package hxrt

import (
	"os"
	"runtime"
	"testing"
)

func TestFileTextHelpersPreserveNativeErrors(t *testing.T) {
	root := t.TempDir()
	missing := root + "/missing.txt"
	if content, err := FileGetContent(&missing); err == nil || content != nil {
		t.Fatalf("missing read = (%v, %v), want nil content and error", content, err)
	}
	content := "not a directory replacement"
	if err := FileSaveContent(&root, &content); err == nil {
		t.Fatal("writing content to a directory returned success")
	}
}

func TestFileTextHelpersPreservePermissionErrors(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("POSIX file modes are not authoritative on Windows")
	}
	if os.Geteuid() == 0 {
		t.Skip("root can bypass the permission fixture")
	}

	path := t.TempDir() + "/locked.txt"
	if err := os.WriteFile(path, []byte("secret"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(path, 0); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := os.Chmod(path, 0o600); err != nil {
			t.Error(err)
		}
	}()

	if _, err := FileGetContent(&path); err == nil {
		t.Fatal("permission-denied read returned success")
	}
	content := "replacement"
	if err := FileSaveContent(&path, &content); err == nil {
		t.Fatal("permission-denied write returned success")
	}
}

func TestSysPutEnvPreservesNativeError(t *testing.T) {
	key := "HAXE_GO=INVALID"
	value := "value"
	if err := SysPutEnv(&key, &value); err == nil {
		t.Fatal("invalid native environment key returned success")
	}
}
