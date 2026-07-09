# Stdlib Helpers — Quick Reference

Common IO & file helpers
- `os.ReadFile(path)` -> []byte, err
- `os.WriteFile(path, data, 0644)`
- `io.Copy(dst, src)` to stream data
- `filepath.Join(dir, file)` to build paths

Collections & slices
- `slices` and `maps` helpers in the standard library (Go 1.21+)

HTTP
- `http.Get(url)` and `http.ListenAndServe(addr, handler)` basics
- Prefer `http.Client{Timeout: ...}` for timeouts

Time
- `time.After`, `time.Sleep`, `time.Tick` (use with caution)

Encoding
- `encoding/json` for JSON marshal/unmarshal

Useful utilities
- `strings` package for common string ops
- `strconv` for conversions
- `math/rand` and `crypto/rand` depending on needs
