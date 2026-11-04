package hxrt

import (
	"os"
	"path/filepath"
	"time"
)

// FileSystemStat is the typed runtime carrier for portable sys.FileStat data.
//
// What: Carries filesystem metadata from Go's os.FileInfo into staged Haxe.
// Why: The public anonymous FileStat record belongs in Haxe source, while Go's
// native metadata and time values cannot cross that boundary directly.
// How: Store portable scalar values and epoch milliseconds; the staged
// sys.FileSystem override constructs the upstream FileStat record and Dates.
type FileSystemStat struct {
	Gid     int
	Uid     int
	AtimeMs float64
	MtimeMs float64
	CtimeMs float64
	Size    int
	Dev     int
	Ino     int
	Nlink   int
	Rdev    int
	Mode    int
}

// FileSystemExists reports whether a filesystem entry is visible at path.
//
// What: Implements the non-throwing existence probe used by sys.FileSystem.
// Why: Go exposes missing and inaccessible paths as errors, but the Haxe exists
// contract reduces every failed stat probe to false.
// How: Delegate to os.Stat and report only its success bit.
func FileSystemExists(path *string) bool {
	_, err := os.Stat(*StdString(path))
	return err == nil
}

// FileSystemRename moves one filesystem entry and preserves native failures.
//
// What: Implements sys.FileSystem.rename through os.Rename.
// Why: Rename errors are catchable Haxe exceptions, not silent failures or
// compiler-emitted Go control flow.
// How: Throw the typed native error through the existing hxrt exception boundary.
func FileSystemRename(path *string, newPath *string) {
	if err := os.Rename(*StdString(path), *StdString(newPath)); err != nil {
		Throw(err)
	}
}

// FileSystemStatPath reads portable metadata for a file or directory.
//
// What: Supplies the scalar fields needed to construct an upstream sys.FileStat.
// Why: os.FileInfo and time.Time are Go-native values, while the public Haxe API
// requires an anonymous record containing Date values.
// How: Read one os.FileInfo, convert its modification time to milliseconds, and
// use conservative cross-platform values for metadata not exposed portably by
// os.FileInfo. Target-specific enrichment can remain behind this typed carrier.
func FileSystemStatPath(path *string) *FileSystemStat {
	info, err := os.Stat(*StdString(path))
	if err != nil {
		Throw(err)
		return &FileSystemStat{}
	}
	modifiedMs := float64(info.ModTime().UnixNano()) / float64(time.Millisecond)
	return &FileSystemStat{
		AtimeMs: modifiedMs,
		MtimeMs: modifiedMs,
		CtimeMs: modifiedMs,
		Size:    int(info.Size()),
		Nlink:   1,
		Mode:    int(info.Mode()),
	}
}

// FileSystemFullPath returns a canonical absolute path for an existing entry.
//
// What: Implements sys.FileSystem.fullPath, including symlink resolution.
// Why: filepath.Abs alone does not satisfy the upstream canonicalization
// contract, and this target-specific operation does not belong in compiler AST.
// How: Make the path absolute, resolve symlinks, normalize it, and translate any
// native failure through the Haxe exception boundary.
func FileSystemFullPath(path *string) *string {
	absolute, err := filepath.Abs(*StdString(path))
	if err != nil {
		Throw(err)
		return StringFromLiteral("")
	}
	resolved, err := filepath.EvalSymlinks(absolute)
	if err != nil {
		Throw(err)
		return StringFromLiteral("")
	}
	return StringFromLiteral(filepath.ToSlash(filepath.Clean(resolved)))
}

// FileSystemAbsolutePath returns an absolute path without requiring it to exist.
//
// What: Implements the Haxe 4.3.7 sys.FileSystem.absolutePath API.
// Why: The earlier compiler shim omitted this method and canonicalization would
// incorrectly reject paths that have not been created yet.
// How: Use filepath.Abs and Clean without resolving filesystem links.
func FileSystemAbsolutePath(path *string) *string {
	absolute, err := filepath.Abs(*StdString(path))
	if err != nil {
		Throw(err)
		return StringFromLiteral("")
	}
	return StringFromLiteral(filepath.ToSlash(filepath.Clean(absolute)))
}

// FileSystemIsDirectory reports whether an existing entry is a directory.
//
// What: Implements sys.FileSystem.isDirectory through os.FileInfo.
// Why: Haxe 4.3.7's eval target returns false for failed stat probes despite the
// upstream API documentation describing an exception; portable semantic-diff
// parity preserves the interpreter's observable behavior.
// How: Collapse stat failures to false and otherwise return FileInfo.IsDir.
func FileSystemIsDirectory(path *string) bool {
	info, err := os.Stat(*StdString(path))
	if err != nil {
		return false
	}
	return info.IsDir()
}

// FileSystemCreateDirectory recursively creates a directory hierarchy.
//
// What: Implements sys.FileSystem.createDirectory.
// Why: Recursive creation is native filesystem behavior and should remain in a
// small typed runtime helper instead of a compiler-owned library algorithm.
// How: Delegate to os.MkdirAll with portable owner-write permissions and throw
// any native error.
func FileSystemCreateDirectory(path *string) {
	if err := os.MkdirAll(*StdString(path), 0o755); err != nil {
		Throw(err)
	}
}

// FileSystemDeleteFile removes one file and preserves native failures.
//
// What: Implements sys.FileSystem.deleteFile.
// Why: Deletion is target-specific I/O, but its public API and ownership remain
// in staged Haxe source.
// How: Use os.Lstat so symlinks remain file-like entries, reject directories,
// then delegate to os.Remove and translate native failures through hxrt.Throw.
func FileSystemDeleteFile(path *string) {
	entryPath := *StdString(path)
	info, err := os.Lstat(entryPath)
	if err != nil {
		Throw(err)
		return
	}
	if info.IsDir() {
		Throw(StringFromLiteral("sys.FileSystem.deleteFile expected a file: " + entryPath))
		return
	}
	if err := os.Remove(entryPath); err != nil {
		Throw(err)
	}
}

// FileSystemDeleteDirectory removes one empty directory.
//
// What: Implements sys.FileSystem.deleteDirectory.
// Why: os.Remove already enforces the empty-directory constraint and returns a
// native error that Haxe callers must be able to catch.
// How: Use os.Lstat to reject files and symlinks, then delegate empty-directory
// enforcement to os.Remove and throw on failure.
func FileSystemDeleteDirectory(path *string) {
	entryPath := *StdString(path)
	info, err := os.Lstat(entryPath)
	if err != nil {
		Throw(err)
		return
	}
	if !info.IsDir() {
		Throw(StringFromLiteral("sys.FileSystem.deleteDirectory expected a directory: " + entryPath))
		return
	}
	if err := os.Remove(entryPath); err != nil {
		Throw(err)
	}
}

// FileSystemReadDirectory lists direct child names for one directory.
//
// What: Implements sys.FileSystem.readDirectory without synthetic compiler code.
// Why: Directory enumeration requires Go's filesystem API, while the resulting
// Array<String> remains part of the ordinary Haxe stdlib contract.
// How: Read sorted os.DirEntry values and convert only their names to hxrt strings.
func FileSystemReadDirectory(path *string) []*string {
	entries, err := os.ReadDir(*StdString(path))
	if err != nil {
		Throw(err)
		return []*string{}
	}
	out := make([]*string, 0, len(entries))
	for _, entry := range entries {
		out = append(out, StringFromLiteral(entry.Name()))
	}
	return out
}
