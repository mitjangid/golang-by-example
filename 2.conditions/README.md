# Conditions Examples

Run these files from the project root:

```bash
go run conditions/ifelse.go
go run conditions/switchcase.go
```

## What To Practice

- Use `if`, `else if`, and `else` for decision making.
- Combine checks with `&&`, `||`, and `!`.
- Use a short statement in `if` when a temporary value is needed only for that block.
- Use `switch` when one value has several possible branches.
- Use a type switch when an `interface{}` or `any` value can hold different types.

## Useful Predefined Functions

```go
time.Now() // Get the current time; useful for date/time decisions.
time.Now().Weekday() // Get the current day of the week for switch cases.
fmt.Println("allowed") // Print the branch result so you can see which path ran.
```

Use this folder when you want to practice writing clear branches before moving to loops.
