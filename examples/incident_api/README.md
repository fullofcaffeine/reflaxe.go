# incident_api

A small HTTP incident-management service written in Haxe stdlib code and compiled to Go.

## Why this example exists

- Shows a real local network service without using `go.*`, Go externs, or raw `__go__`.
- Demonstrates config files, file-backed state, JSON payloads, and loopback sockets.
- Proves the same Haxe source can compile under `portable` and `metal` profiles.

## What it does

The service exposes a minimal incident API:

- `GET /health`
- `GET /incidents`
- `POST /incidents`
- `POST /incidents/{id}/ack`
- `POST /incidents/{id}/resolve`
- `GET /metrics`

The implementation uses Haxe-facing APIs only:

- `sys.net.Socket` for the tiny HTTP/1.1 server
- `sys.io.File` and `sys.FileSystem` for config/state files
- `haxe.Json` for request/config/state parsing

## Profile support

| Profile | Supported | Practical meaning |
| --- | --- | --- |
| `portable` | Yes | Default showcase: ordinary Haxe stdlib code becomes a runnable Go service. |
| `metal` | Yes | Stricter profile lane proving the same service avoids raw/native shortcuts. |

`metal` is not required for good Go output. This example uses `metal` as an audit lane, not as permission to write Go-native app code.

If you want an example where `metal` unlocks more visible Go-native authoring choices, use `examples/worker_pool_select`. This service is intentionally different: it proves a useful network app can stay on Haxe-facing APIs.

## Compile

```bash
haxe compile.portable.hxml
haxe compile.metal.hxml
```

CI lanes:

```bash
haxe compile.portable.ci.hxml
haxe compile.metal.ci.hxml
```

## Run scripted contract

```bash
(cd out_portable && go run . --scripted)
(cd out_metal && go run . --scripted)
```

The scripted mode starts a loopback socket server, sends real HTTP requests through `sys.net.Socket`, verifies deterministic responses, and exits.

## Run live service

```bash
(cd out_portable && go run . init-config --config config.json)
(cd out_portable && go run . serve --config config.json)
```

Then in another terminal:

```bash
curl http://127.0.0.1:8080/health
curl -X POST -d '{"title":"Database lag","severity":"high"}' http://127.0.0.1:8080/incidents
curl http://127.0.0.1:8080/incidents
curl -X POST http://127.0.0.1:8080/incidents/1/ack
curl -X POST http://127.0.0.1:8080/incidents/1/resolve
curl http://127.0.0.1:8080/metrics
```

If the generated config keeps `port` as `0`, the service prints the chosen ephemeral port at startup. Set `port` to `8080` for the exact curl commands above.

## HTTP subset

This is intentionally not a web framework. It supports:

- one request per connection
- `GET` and `POST`
- `Content-Length` request bodies
- JSON responses
- `Connection: close`

It does not support TLS, chunked encoding, keep-alive, routing middleware, or full HTTP compliance.

## Expected output

Scripted outputs are validated by:

- `expected/portable.stdout`
- `expected/portable.ci.stdout`
- `expected/metal.stdout`
- `expected/metal.ci.stdout`

## Related docs

- [`docs/examples-matrix.md`](../../docs/examples-matrix.md)
- [`docs/profiles.md`](../../docs/profiles.md)
- [`docs/profile-semantics-guide.md`](../../docs/profile-semantics-guide.md)
- [`docs/feature-support-matrix.md`](../../docs/feature-support-matrix.md)
