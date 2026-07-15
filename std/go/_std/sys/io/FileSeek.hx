package sys.io;

/**
	What
	- Defines the Haxe 4.3.7 file seek origins used by staged Go file streams.

	Why
	- The mainstream Haxe stdlib implementation cannot be used unchanged on `haxe.go`
	  while `FileInput` and `FileOutput` are target overrides: the enum must be emitted
	  with those source-owned modules instead of synthesized as a compiler carrier.

	How
	- Preserve the upstream constructors exactly; the staged stream methods map each
	  constructor to Go's start, current, or end seek origin.
**/
enum FileSeek {
	SeekBegin;
	SeekCur;
	SeekEnd;
}
