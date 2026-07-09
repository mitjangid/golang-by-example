# Testing & Benchmarks — Practical Cheatsheet

Table-driven tests
- Use slices of test cases and `t.Run` for clarity.

Example test template
```
func TestAdd(t *testing.T) {
  tests := []struct{ name string; a, b, want int }{
    {"simple", 1, 2, 3},
  }
  for _, tc := range tests {
    t.Run(tc.name, func(t *testing.T) {
      got := Add(tc.a, tc.b)
      if got != tc.want { t.Fatalf("want %d got %d", tc.want, got) }
    })
  }
}
```

Benchmarks
- Add `func BenchmarkX(b *testing.B)` and put the hot loop inside:
```
func BenchmarkConcat(b *testing.B) {
  for i := 0; i < b.N; i++ { _ = strings.Join(parts, ",") }
}
```
Run: `go test -bench . -benchmem`

Coverage
- `go test ./... -coverprofile=coverage.out`
- `go tool cover -html=coverage.out` to view in browser

Common gotchas
- Tests that rely on global state may be flaky; prefer isolated table-driven tests.
- Use `b.ReportAllocs()` to see allocations in benchmarks.

Try it
- Add a table-driven test for `packages/user` constructors and run `go test ./...`.
