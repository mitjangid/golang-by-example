# Generics — Quick Patterns

When to use
- Generics simplify reusable algorithms (stacks, sets, maps helpers) while preserving type-safety.

Simple generic function
```
func IndexOf[T comparable](slice []T, v T) int {
  for i, x := range slice { if x == v { return i } }
  return -1
}
```

Generic stack type
```
type Stack[T any] struct{ elems []T }
func (s *Stack[T]) Push(v T) { s.elems = append(s.elems, v) }
func (s *Stack[T]) Pop() (T, bool) {
  if len(s.elems) == 0 { var z T; return z, false }
  v := s.elems[len(s.elems)-1]
  s.elems = s.elems[:len(s.elems)-1]
  return v, true
}
```

Constraints examples
- `comparable` for equality checks.
- Define custom constraints using interfaces if needed.

Pitfalls
- Avoid over-abstracting; generics are best for clear reusable behaviors.
- Type inference works in many cases; explicit type parameters can be used when needed.
