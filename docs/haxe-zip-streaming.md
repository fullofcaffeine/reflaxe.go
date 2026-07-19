# `haxe.zip` streaming on Go

`haxe.zip.Compress` and `haxe.zip.Uncompress` support repeated partial-buffer
execution on `haxe.go`. This is the portable Haxe API: callers keep source and
destination cursors in Haxe, while an opaque typed Go handle retains the native
codec state.

## What `execute` guarantees

Each call returns `{done, read, write}`:

- `read` is the number of bytes consumed starting at `srcPos`.
- `write` is the number of bytes written starting at `dstPos`; it never exceeds
  the remaining destination capacity.
- `done` becomes `true` only after the stream has ended and all pending output
  has reached the caller.

A codec may consume all available source while producing only part of its
output. A later call can therefore validly return `read == 0` and `write > 0`.
Advance each cursor only by the corresponding result field and continue until
`done`.

Inflation applies output backpressure to the native decoder. Highly compressed
input therefore cannot expand into an unbounded hidden buffer: pending decode
state stays within the current or still-active destination allowance.

```haxe
var codec = new haxe.zip.Compress(6);
codec.setFlushMode(haxe.zip.FlushMode.FINISH);

var sourcePosition = 0;
var output = new haxe.io.BytesBuffer();
var done = false;
while (!done) {
	var destination = haxe.io.Bytes.alloc(4096);
	var step = codec.execute(source, sourcePosition, destination, 0);
	sourcePosition += step.read;
	output.addBytes(destination, 0, step.write);
	done = step.done;
}
codec.close();
var compressed = output.getBytes();
```

A zero-capacity destination returns zero progress. Supply a non-empty
destination before continuing. Positions outside `0...bytes.length` fail with
`haxe.io.Error.OutsideBounds`; the end position itself is valid for draining
pending output.

## Flush-mode policy

Go's public standard-library compressor exposes exact no-flush, sync-flush, and
finish operations. It does not expose zlib's dictionary-reset or block-boundary
controls, so `haxe.go` reports those limitations instead of silently weakening
their meaning.

| `FlushMode` | Go behavior |
| --- | --- |
| `NO` | Accept input without forcing buffered output. |
| `SYNC` | Flush output so a decoder can consume the stream so far. |
| `FINISH` | Finish the stream; call `execute` again as needed to drain output. |
| `FULL` | `setFlushMode` throws an explicit unsupported-mode error. |
| `BLOCK` | `setFlushMode` throws an explicit unsupported-mode error. |

`Uncompress` accepts `NO`, `SYNC`, and `FINISH` as progressive decode policy;
`FULL` and `BLOCK` are rejected because Go cannot promise their zlib boundary
behavior. This target policy follows the capabilities documented by Go's
[`compress/flate.Writer`](https://pkg.go.dev/compress/flate#Writer) rather than
pretending that `SYNC` is equivalent to the missing operations.

## Lifecycle and raw DEFLATE

`close()` is idempotent. Calling `execute` or `setFlushMode` after close throws a
deterministic Haxe exception. Close an incomplete `Uncompress` instance when it
will not receive more input; its native reader can otherwise remain paused
waiting for the next fragment.

Passing a negative `windowBits` value, conventionally `-15`, creates a raw
DEFLATE inflater for ZIP entry payloads. Positive and omitted values select a
zlib-wrapped stream. Static `Compress.run`, `Uncompress.run`, and
`haxe.zip.Tools` retain their source-owned one-shot behavior.

## Ownership and implementation boundary

Staged Haxe source owns `Bytes` conversion, offsets, result records, flush
selection, bounds, lifecycle errors, static helpers, and raw-stream selection.
`std/hxrt/zip` exposes only two opaque handles and a typed scalar/byte-slice
step carrier. `runtime/hxrt/zip.go` owns the live zlib writer or inflater.

The inflater uses a fragment feeder that implements Go's exact byte-reader
path. It pauses the real decoder at temporary input or output boundaries,
resumes without replaying earlier compressed data, and leaves trailing bytes
unread after the end of one stream. No generated `haxe.io.Bytes` layout,
`Dynamic` native handle, reflection, `unsafe`, raw Go injection, compiler shim,
profile branch, or whole-program IR is involved.
