# Errors and Logging — Quick Guide

Error wrapping & inspection
- Wrap with context: `fmt.Errorf("read config: %w", err)`
- Check errors: `errors.Is(err, target)` and `errors.As(err, &typed)`
- Prefer returning errors early and adding context at call sites.

Example
```
data, err := os.ReadFile(path)
if err != nil {
  return fmt.Errorf("read %s: %w", path, err)
}
```

Logging
- Small examples using the standard `log` package:
```
log.Printf("started service: %s", name)
log.Fatalf("fatal error: %v", err) // exits after print
```
- For structured logging prefer a small library (zerolog, zap) in real projects; for learning, stdlog is fine.

Best practices
- Do not swallow errors. Log or return them.
- Wrap errors with context close to the failing operation, not far away.
- Consider sentinel errors sparingly — prefer typed errors or error wrapping.

Try it
- Write a function that opens a file and returns wrapped errors. Add a test that asserts `errors.Is` or `errors.As` on the returned error.
