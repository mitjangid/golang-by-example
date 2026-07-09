# Slices, Arrays & Maps — Quick Cheatsheet

Summary
- Use arrays when size is fixed and you need a true array type. Use slices for flexible lists. Use maps for key/value lookup.

Slices
- Create literal: `a := []int{1,2,3}`
- Make with capacity: `s := make([]int, 0, 8)`
- Append: `s = append(s, 4)`
- Sub-slice: `s[1:3]` // start inclusive, end exclusive
- Copy: `dst := make([]int, len(src)); copy(dst, src)`
- Compare: `import "slices"; slices.Equal(a,b)`

Len vs Cap
- `len(s)` is number of elements; `cap(s)` is how many elements the backing array can hold before growth.
- Appending beyond cap allocates a new backing array — existing slices sharing the old backing array are not updated.

Maps
- Create: `m := map[string]int{"x":1}`
- Add/Update: `m["k"] = 5`
- Check existence: `v, ok := m["k"]`
- Delete: `delete(m, "k")`
- Zero value: `var m map[string]int` is nil — writing to it panics; initialize with `make`.

Common pitfalls
- Sharing backing arrays unintentionally: `s2 := s[:2]` may mutate same storage.
- Out-of-bounds indexing causes panic. Prefer safe checks or append.
- Confusing nil slice vs empty slice: `var s []int` vs `s := []int{}`

Try it
1) Capacity growth: create `s := make([]int, 0, 2)` and append five values; print `len` and `cap` after each append.
2) Safe map getter: `func GetAge(m map[string]int, name string) (int, bool) { v, ok := m[name]; return v, ok }`
