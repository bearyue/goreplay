# GitHub Actions Linux amd64 Tarball Packaging (Manual)

## Goal
Provide a GitHub Actions workflow that manually builds and packages the project into a Linux `amd64` tarball (`.tar.gz`) without requiring local build environment setup.

## Scope
- Add a new workflow triggered via `workflow_dispatch` only.
- Build Linux `amd64` binary on `ubuntu-latest`.
- Package as `dist/gor_<version>_linux_amd64.tar.gz`.
- Upload the tarball as a workflow artifact.

Out of scope:
- Release publishing
- Multi-platform builds
- `deb`/`rpm` packaging
- Docker image publishing

## Constraints and Assumptions
- Build must use `libpcap-dev` (cgo dependency via `gopacket/pcap`).
- Use Go version compatible with the project’s `go.mod`.
- The version string comes from `version.go` (`VERSION` constant).
- No changes to existing CI workflows unless necessary.

## Workflow Design

### Triggers
- `on: workflow_dispatch`

### Job: `build-linux-amd64`
Runs on `ubuntu-latest` and performs:
1. Checkout repository
2. Install Go (version from `go.mod` or fixed)
3. Install `libpcap-dev`
4. `go mod vendor`
5. Build:
   - `GOOS=linux`
   - `GOARCH=amd64`
   - `CGO_ENABLED=1`
   - `go build -mod=vendor -o gor ./cmd/gor/`
6. Package:
   - `mkdir -p dist`
   - `tar -czf dist/gor_<version>_linux_amd64.tar.gz gor`
7. Upload artifact `gor_<version>_linux_amd64.tar.gz`

### Version Resolution
Extract `VERSION` from `version.go` using a simple `grep/sed` in bash. If extraction fails, fall back to `dev`.

## Error Handling
- Fail fast on any command error (`set -euo pipefail`).
- Explicitly log the resolved version and artifact path.

## Testing Strategy
- Manual trigger in GitHub Actions.
- Confirm artifact exists and downloads successfully.
- Optionally run the binary on a Linux host to validate.

## Files to Change
- Add workflow: `.github/workflows/package-linux-amd64.yml`
- No changes to application source code.
