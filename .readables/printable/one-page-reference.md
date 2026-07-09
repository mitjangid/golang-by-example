# One-page Reference — Printable

Top rules
- Prefer small, focused functions.
- Check and handle errors immediately (wrap with %w for context).
- Use slices for lists; maps for lookup; channels for communication.
- Use `defer` to release resources.
- Prefer `context` for cancellation/timeouts in concurrent code.

Commands
- Run: `go run <file>`
- Test: `go test ./...`
- Format: `gofmt -w .`
- Tidy: `go mod tidy`

Quick idioms
- Table-driven tests
- `slices.Equal` for slice comparisons
- `errors.Is` / `errors.As` for error inspection

Keep this sheet nearby while exploring examples.
