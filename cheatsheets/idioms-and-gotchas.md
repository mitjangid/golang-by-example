# Idioms & Gotchas — Fast Reference

Defer
- `defer` runs at the end of the surrounding function. Use `defer file.Close()` after successful open.

Zero values & constructors
- Every Go type has a zero value. Prefer `var x T` when zero value is sensible.
- Provide `NewX` constructor if initialization is non-trivial.

Pointer vs value receiver
- Use pointer receivers to modify state or avoid copying large structs.
- Use value receivers for small immutable values.

Nil interface vs nil pointer
- `var i interface{}` is nil; `var p *T` is nil pointer. An interface holding a nil pointer is NOT nil.

Common bug
- Returning an interface containing a nil pointer: the interface is non-nil — checks like `if err != nil` may be true when you expect false. Use typed nil checks or return nil interface.
