// Package cond offers functions to conditionally apply, act, pass, or return
// values.
package cond

// Ternary is intended to behave like a ternary operator.
func Ternary[T any](cond bool, a, b T) T {
	if cond {
		return a
	}
	return b
}
