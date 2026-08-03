# Appwrite CLI (Go) — preview build

A preview of the Appwrite CLI rewritten in Go, for testing before it replaces
the TypeScript/Bun CLI. Same command surface, same flags, same machine-readable
output — roughly 30× faster to start and a much smaller download.

> **This is a testing fork, not an official release.** It is built from an
> unmerged branch of [appwrite/sdk-generator](https://github.com/appwrite/sdk-generator)
> and published here so it can be tried out. Use the official
> [appwrite/sdk-for-cli](https://github.com/appwrite/sdk-for-cli) for anything real.

## Install

Download a binary from [Releases](https://github.com/ChiragAgg5k/appwrite-cli-go/releases),
or build from source:

```bash
go install github.com/ChiragAgg5k/appwrite-cli-go@latest
```

`go install` names the binary after the repository, so rename it if you want it
to stand in for the existing CLI:

```bash
mv "$(go env GOPATH)/bin/appwrite-cli-go" "$(go env GOPATH)/bin/appwrite"
```

### Installing a downloaded binary

```bash
chmod +x appwrite-cli-darwin-arm64
xattr -d com.apple.quarantine appwrite-cli-darwin-arm64 2>/dev/null || true
mv appwrite-cli-darwin-arm64 /usr/local/bin/appwrite
```

The macOS binaries carry an ad-hoc code signature, which is what the loader
requires. They are **not** notarised, so a browser download will be quarantined
— hence the `xattr` line. `curl` and the Releases API do not set that attribute.

## Try it

```bash
appwrite --version
appwrite --help

# Point it at an instance
appwrite client --endpoint https://cloud.appwrite.io/v1 --project-id <id> --key <key>

# Or log in interactively
appwrite login
appwrite whoami
```

Config lives in `~/.appwrite/prefs.json` and `appwrite.config.json`, in the same
format the TypeScript CLI uses, so it reads and writes an existing setup.
Credentials go to the OS keychain where there is one, and fall back to
`prefs.json` where there is not.

## What to test, and what to report

Most valuable: anything **stateful**. `push`, `pull`, `init` and `run` are where
a port can diverge in ways that still compile and still pass unit tests.

Worth reporting:

- a command that behaves differently from the TypeScript CLI
- different JSON from `--json` for the same call
- a different `appwrite.config.json` after the same `init` or `pull`
- a flag that exists on one CLI and not the other

Human-readable table output is **not** expected to match yet — per-resource
field formatting is not ported, so values render as raw JSON scalars. That is
known, and not worth reporting.

## Known gaps

- `run --with-variables` warns instead of fetching remote variables, and `run`
  has no JWT manager
- `init project` has no autopull prompt and does not install skills
- Per-resource output formatting is not ported
- Keyring round-trips are verified on macOS and Linux; **Windows is unverified**

## Build from source

```bash
git clone https://github.com/ChiragAgg5k/appwrite-cli-go
cd appwrite-cli-go
go build -o appwrite .
go test ./...
```

## Layout

| Path | |
|---|---|
| `internal/cmd` | commands — generated service commands plus the hand-written ones |
| `internal/appwritesdk` | the Appwrite Go SDK, vendored |
| `internal/config` | `appwrite.config.json` and `prefs.json`, order-preserving |
| `internal/sdk` | the one place an API client is built |

The SDK is vendored rather than depended on, because the published
`appwrite/sdk-for-go` does not yet contain three packages this CLI needs
(`migrations`, `notifications`, `vcs`). Vendoring keeps this repo installable
with a single `go install`. The real repository will depend on a released SDK.
