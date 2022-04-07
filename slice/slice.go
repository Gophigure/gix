// Package slice offers functions to mutate or create new slices from existing
// ones. None of these functions mutate the original slice, instead returning a new
// slice with the changes applied.
package slice

// Pop removes and returns the last item from the slice.
func Pop[T any](slice []T) (popped T, altered []T) { return slice[len(slice)-1], slice[:len(slice)-1] }

// Shift removes and returns the first item from the slice.
func Shift[T any](slice []T) (shifted T, altered []T) { return slice[0], slice[1:] }

// Prepend is the reverse of the built-in append.
func Prepend[T any](slice []T, v ...T) []T { return append(v, slice...) }

// Map calls a function on every member of the slice, returning a slice created from the mutated
// values.
func Map[T any](slice []T, f func(int, T) T) []T {
	l := len(slice)
	altered := make([]T, l, l)
	for i, v := range slice {
		altered[i] = f(i, *&v)
	}
	return altered
}

// Reduce calls a function on every member of a slice, returning a slice created only from "correct"
// values.
func Reduce[T any](slice []T, f func(int, T) bool) []T {
	l, n := len(slice), 0
	altered := make([]T, l, l)
	for i, v := range slice {
		if f(i, v) {
			altered[n] = v
			n++
		}
	}
	return altered[:n]
}
