# Portable HTTP client admission design

## Outcome

A portable Haxe caller should observe an HTTP exchange as it happens:

1. the server's status and headers become visible;
2. a caller-supplied `Output` is prepared when the body length is known;
3. bounded body chunks are written as they arrive;
4. a complete body is classified as success or an HTTP-status error;
5. a later transfer failure keeps every byte already delivered.

The request side has the same rule: `set*` and `add*` calls become the same
ordered, repeated fields on the wire. Multipart uploads are pumped by staged
Haxe on the synchronous caller, while Go owns only the native transport and a
cancellable byte sink.

This design does not itself admit HTTP into the beta release. HTTP remains
release-excluded until the implementation children pass their evidence gates
and `haxe_go-vfp.10.8.6` completes an independent operation-level admission
review.

## Why the old seam is insufficient

The current native call returns one status, one header map, and one fully
buffered body. Consider a server that flushes status `200` and the bytes
`hello`, then pauses before sending the rest:

- Haxe 4.3.7 calls `onStatus(200)` and writes `hello` before the pause ends.
- The current Go target calls neither because `io.ReadAll` has not returned.
- If the connection then fails, the current carrier discards `hello`.

The current upload direction has the inverse ownership problem. Go's
`net/http` transport asks an `io.Reader` for more request bytes from its own
write goroutine. That reader calls a generated-Haxe closure, so a transport
goroutine enters Haxe code and mutates captured Haxe state. A no-op request-body
closer also cannot unblock that closure if `Input.readBytes` waits forever.

Small completed-response fixtures did not expose either difference. They
proved final values after success, not the lifecycle that produces those
values.

## Contract authority

This decision uses:

- Haxe 4.3.7 `std/sys/Http.hx` and `std/haxe/http/HttpBase.hx` for the public
  source contract;
- `std/go/_std/sys/Http.hx` for target-staged policy;
- typed externs under `std/hxrt/http` for the native capability boundary;
- `runtime/hxrt/http.go` for Go resource ownership;
- the accepted findings in
  [the independent network disposition](reviews/network-admission-oracle-disposition-vfp-10.4.md).

The mainstream source is the default semantic authority, but it is not a
security authority. Where its raw HTTP/1 implementation accepts dangerous or
ambiguous request metadata, the Go target may reject the request before network
I/O. Such differences must be explicit below and in the operation manifest.

## Source and native ownership

The staged `sys.Http` implementation owns everything a Haxe program can
observe:

- the ordered parameter and header collections inherited from `HttpBase`;
- percent encoding and GET/form/multipart selection;
- explicit method and body selection;
- response maps and repeated header values;
- `onStatus`, `Output.prepare`, body writes, status classification,
  `Output.close`, `onError`, `onData`, and `onBytes`;
- partial `responseBytes` and lazy `responseData`;
- the source exception produced by a custom upload `Input` or response
  `Output`.

Typed `hxrt` owns resources that cannot be ordinary portable Haxe values:

- parsed URLs and one-use `http.Transport` values;
- proxy, dial, socket, TLS, and request context resources;
- a native upload pipe and its cancellation;
- the live `http.Response.Body`;
- progress deadlines and cleanup.

The boundary passes strings, scalars, immutable `ByteView` chunks, and opaque
typed handles. Native code never inspects a generated Haxe object, and
generated Haxe is never called from a native goroutine.

This staged-library and typed-runtime change does not require a universal compiler IR.
It also needs no compiler-owned `sys.Http` shim, raw `__go__` injection,
`Dynamic` native handle, or behavior selected by the `portable|metal`
compatibility preset.

## Lifecycle model

The positive lifecycle is:

```text
build ordered request
    -> start native exchange
    -> pump upload from the Haxe caller, when present
    -> await response headers
    -> stream bounded response chunks into the Haxe Output
    -> classify status
    -> close Output and native resources
```

Each opaque resource has one owner at a time:

| Handle | Created by | Owner and lifetime |
| --- | --- | --- |
| `HttpRequestHandle` | `newRequest` | Staged `sys.Http` owns the mutable builder until `startExchange` freezes and consumes it. |
| `HttpExchangeHandle` | `startExchange` | Staged `sys.Http` owns the public lifecycle; native code owns the context, transport, response, and cleanup state until `closeExchange` or `cancelExchange`. |
| `HttpUploadSinkHandle` | `startExchange` for multipart | Staged `sys.Http` is its only writer. Native transport is its only reader. Finishing, canceling, an early response, or a transport failure closes it exactly once. |
| `HttpReadResultHandle` | `readResponseChunk` | One immutable result describes a bounded response chunk plus its EOF or error state; it owns no live resource after its accessors return. |

The intended typed capabilities are:

| Capability | Positive contract |
| --- | --- |
| `newRequest` and ordered add/set functions | Build a source-ordered request without starting network I/O. |
| `startExchange` | Freeze the request, allocate native resources, and start a native exchange without calling Haxe. |
| `writeUploadChunk` | Copy one bounded immutable chunk from the Haxe caller into the native upload sink or return its synchronized terminal error. |
| `finishUpload` | Declare exact upload completion and close the writer side once. |
| `abortUpload` | Close the writer with the source-owned upload error so the transport stops reading. |
| `awaitResponse` | Wait for response headers or the exchange's transport error. |
| response header/content-length accessors | Expose native response facts after `awaitResponse`; they do not classify status. |
| `readResponseChunk` | Return at most the requested bounded response chunk and preserve bytes even when the same native read also reports EOF or failure. |
| `closeExchange` | Release a successfully completed response, transport, optional socket, and idle connections once. |
| `cancelExchange` | Cancel an incomplete exchange, close upload/response resources, and unblock native waits once. |

No public path may abandon an exchange without calling `closeExchange` or
`cancelExchange`.

## Request contract

### Parameters

`HttpBase` already implements the public distinction:

- `addParameter` appends a new ordered entry;
- `setParameter` replaces the first exact-name entry or appends when absent.

Staged `sys.Http` passes every remaining entry to `HttpRequestHandle` in that
order. It also passes the `StringTools.urlEncode` spelling of each name and
value, because URL encoding is Haxe-visible library policy rather than native
resource policy.

For GET-like requests, encoded entries append to the existing raw query with
`?` or `&`. Existing repetition, escaping, and spelling are not parsed,
sorted, collapsed, or re-encoded. For form POST requests, the body contains
only the configured parameters; URL query values stay in the URL.

Multipart fields use the raw names and values in the same source order.

### Headers

`addHeader` appends and `setHeader` replaces the first exact-name entry, just
as `HttpBase` specifies. Ordinary headers are installed with native `Add`
semantics so repeated values retain their per-name order.

Go-special request fields receive deliberate policy:

- `Host` sets `http.Request.Host` rather than an ineffective map entry;
- `Content-Length` must parse and equal the exact body length when that length
  is known, otherwise request validation fails before dialing;
- `Connection: close` sets the native close policy;
- unsupported `Connection` tokens, caller-provided `Transfer-Encoding`,
  `Trailer`, and `Upgrade` are rejected before dialing until they have their
  own typed implementation and evidence.

This first admission does not promise global header-line ordering, because
Go's transport owns serialization. It does promise every configured value and
the order of repeated values for a given name.

### Method and body

An explicit non-empty method token is preserved exactly; it is not uppercased.
Without one, the method is `POST` when the source selects post semantics and
`GET` otherwise.

An explicit `postData` or `postBytes` body is sent even when
`customRequest(false, ..., method)` supplied `false`. The Boolean still chooses
whether configured parameters become a query or form body when no explicit
body exists.

### Multipart framing

The native request builder creates one high-entropy boundary per request and
derives the body prefix, tail, exact content length, and complete
`multipart/form-data; boundary=...` content type from that same value.

A caller-provided multipart content type with a different boundary is rejected
before dialing. The staged class never duplicates a boundary constant.
Parameter names, upload field names, filenames, and media types reject CR, LF,
and NUL. Media types must parse as media types before any bytes are written.

## Response contract

Response processing uses one order for ordinary network and `data:` responses:

```text
headers -> `onStatus` -> `prepare` -> body writes
complete body -> status classification -> Output close
partial body -> transfer error -> `onError`
```

`prepare` runs only when a nonnegative declared content length fits the Haxe
`Int` accepted by `Output.prepare`. An oversized declared length fails before
body allocation rather than truncating the value.

Each `HttpReadResultHandle` can carry bytes and a terminal read state together.
Staged `sys.Http` writes the bytes first, then handles EOF or failure. This
ensures a successful partial native read is never discarded.

The two public entry points finish differently:

- `customRequest` writes directly to its caller-owned `Output`. It closes that
  Output exactly once only after a complete body accepted by the source status
  classifier. The final admitted policy is 200 through 399; the response slice
  still uses the pre-existing `< 400` rule until `haxe_go-vfp.10.8.5` adds the
  missing low-status classification. An HTTP-status error, transfer error, or
  Output exception leaves it open.
- `request()` wraps `customRequest` with a source-owned `BytesOutput`, like
  Haxe 4.3.7. On failure, `request()` retains every byte already written in
  `responseBytes`. On success it publishes the complete bytes through
  `responseBytes`, lazy `responseData`, `onData`, and `onBytes`.

An Output exception from `prepare` or `writeBytes` cancels the native exchange
and reaches `onError` once. If `close` itself throws, staged code does not try
to close the Output a second time; it releases native resources and reports
the close failure once.

If `onStatus` throws, the exchange is canceled and the source error path calls
`onError`, matching the upstream outer request guard. An exception thrown by
`onError` itself propagates. Successful `request()` callbacks (`onData` and
`onBytes`) run after `customRequest` completes; their own exceptions propagate
instead of being misclassified as transport failures.

### Implemented response slice

`haxe_go-vfp.10.8.3` implements this response contract with
`HttpExchangeHandle` and `HttpReadResultHandle`. The native start operation
returns after headers, retains no aggregate body, and exposes reads capped by
the staged 1024-byte chunk size. Cleanup is idempotent and owns the response
body, request context, one-use transport, and optional socket.

The decisive contracts prove first-chunk visibility before server completion,
known-length `prepare`, partial bytes plus transfer failure, partial
`request().responseBytes`, one error for throwing `onStatus`/`prepare`/
`writeBytes`/`close`, bounded large-response reads, and matching `data:` event
order. This does not complete HTTP admission: the upload callback still crosses
from the native transport until `haxe_go-vfp.10.8.4`, and the policy listed
under Remaining client policy still belongs to `haxe_go-vfp.10.8.5`.

## Upload contract

`startExchange` creates a native pipe when multipart upload data is present and
starts `net/http` entirely behind the native boundary. Staged `sys.Http` then:

1. asks its typed `Input` for at most the configured chunk size;
2. passes the returned immutable view to `writeUploadChunk`;
3. repeats until the exact declared size is written;
4. calls `finishUpload`.

Early EOF, zero progress, too many bytes, or a source exception calls
`abortUpload`. If the server responds early or the exchange times out, native
code closes the sink so the next `writeUploadChunk` fails promptly. A response
already published by native code is processed instead of being hidden behind a
generic closed-pipe error.

There is one unavoidable synchronous-source limitation. An arbitrary custom
`Input.readBytes` implementation can wait forever inside user Haxe code. It
cannot be forcibly interrupted without moving generated Haxe onto another
goroutine, which this design forbids. That case is explicitly excluded from
the cancellable-upload claim. Once `readBytes` returns, native cancellation
must stop the next sink write promptly. Because all reads occur on the public
caller, no upload work can continue after the public request returns.

## Error precedence

The exchange records concurrent native terminal state under synchronization:
first terminal event wins. Staged code reports exactly one public error by
evaluating these branches in lifecycle order:

1. request or multipart validation before dialing;
2. a source `Input` exception or declared-size violation already observed by
   the Haxe caller;
3. a response that native code has already published after an upload write
   closes, including an early HTTP error status;
4. the synchronized native timeout, cancellation, dial, TLS, proxy, socket, or
   upload-sink error;
5. an Output exception at the exact `prepare` or body-write operation where it
   occurs;
6. a response-body read error, after first attempting to write bytes returned
   with that error;
7. HTTP status below 200 or at least 400, after the complete body;
8. an Output `close` exception after an otherwise successful body and status.

An earlier Output exception remains the public error even if cancellation
causes a later native closed-resource error. Cleanup errors are retained for
diagnostics where useful but never cause a second `onError`.

## Timeout contract

`cnxTimeout` is a progress timeout, not one total wall-clock budget:

- `cnxTimeout < 0` means no native deadline;
- `cnxTimeout == 0` means an immediate native deadline;
- `cnxTimeout > 0` bounds connect/TLS setup, waiting for response headers, each
  upload sink write, and each response read separately.

Successful progress starts a fresh budget for the next blocking operation. A
slow response that continues producing chunks within the configured interval
does not fail merely because total request time exceeds `cnxTimeout`.

The deadline covers native waits. It cannot cover time spent blocked inside an
arbitrary custom `Input.readBytes`; that source-owned limitation is the
explicit upload exclusion above. The implementation must not silently replace
zero or negative values with ten seconds and must not use
`http.Client.Timeout` as a hidden total deadline.

## Remaining client policy

`haxe_go-vfp.10.8.5` owns behavior that is orthogonal to streaming:

- do not automatically follow redirects;
- disable transparent response decompression so body and headers retain their
  received meaning;
- classify status below 200 or at least 400 after streaming the body;
- make `requestUrl` throw its error instead of returning an error string;
- define proxy and custom-socket behavior separately for HTTP and HTTPS.

Received redirect and compressed response bodies remain available to the
caller's Output under these rules.

## Finding and evidence map

| Finding | Implementation owner | Decisive evidence |
| --- | --- | --- |
| H1 | `haxe_go-vfp.10.8.2` | Raw TCP capture for interleaved repeated parameters/headers, set/add order, existing escaped queries, query/form separation, Host, and special fields. |
| H2 | `haxe_go-vfp.10.8.2` | Distinct atomic boundaries, conflicting content type, delimiter-shaped payload, exact lengths, zero-byte upload, and multipart field order. |
| L2 | `haxe_go-vfp.10.8.2` | Hostile field name, filename, and media-type metadata fails before a server accepts a connection. |
| B1 | `haxe_go-vfp.10.8.3` | Header/first-chunk flush before server release, truncated partial body, `request()` partial bytes, bounded retention, and throwing Output methods. |
| B2 | `haxe_go-vfp.10.8.4` | Generated-Haxe early 413, timeout, server close, source failure, exact-size/EOF cases under race, with no native callback into Haxe. |
| H3 | `haxe_go-vfp.10.8.5` | Redirect destination not contacted, mixed-case method, body with `post == false`, low status, gzip, progress timeout, `requestUrl`, data URL, proxy, and custom socket. |

`haxe_go-vfp.10.8.6` owns repeated goroutine/connection/file-descriptor
convergence, full supported-toolchain evidence, compatibility regeneration, and
the independent xhigh admission review.

## Second-pass challenge review

The required xhigh second pass challenged the design against six plausible
ways it could still produce believable but incorrect evidence:

1. **A total client timer could still fail a progressing response.** The
   contract therefore forbids `http.Client.Timeout` and assigns a fresh native
   deadline to each connect, header, sink-write, and body-read wait.
2. **A response read could return bytes and an error together.** The
   `HttpReadResultHandle` carries both; staged code attempts the byte write
   before reporting the read failure. If that write throws, the Output error is
   the first Haxe-visible failure.
3. **Canceling native I/O could be mistaken for canceling user code.** The
   forever-blocking custom `Input.readBytes` case is an explicit exclusion. No
   native goroutine is allowed to enter Haxe merely to make that claim appear
   bounded.
4. **Go serialization could be described as exact global header order.** The
   contract promises ordered values within each repeated name, explicitly
   handles Go-special fields, and leaves global header-line ordering
   unadmitted.
5. **Cleanup could duplicate callbacks or close attempts.** Every live handle
   has one owner and one terminal path. Cleanup records diagnostics but cannot
   call Haxe, retry `Output.close`, or emit a second `onError`.
6. **A broken native reader could escape the bounded-read contract.** The
   implementation checkpoint found that an invalid negative or oversized read
   count would panic while constructing the immutable view. A red regression
   now proves both cases become typed read errors with no exposed bytes.

No challenge changed the staged-source/typed-runtime architecture. The review
did narrow the claims for custom blocking input and header-line order, fixed
the invalid-count panic before closure, and made Output-versus-read-error
precedence explicit.

## Admission boundary

Until the final review:

- HTTP remains release-excluded, even when an implementation fixture passes;
- the supported runtime platform remains Linux/amd64;
- Windows cross-builds remain compile-only evidence;
- Darwin developer runs do not become governed release evidence;
- arbitrary custom upload inputs that block forever inside `readBytes` remain
  explicitly excluded;
- public Internet, hostile-peer, proxy fleet, public CA-store, HTTP/2, and
  long-duration load guarantees remain excluded unless the final
  operation-level manifest names and proves them.

The final review may admit smaller operations independently. It must not turn a
passing `haxe.Http` module row into blanket admission of every method, error
path, custom source/sink, platform, or network environment.
