package math

type (
	// Int constrains to only builtin signed integer types.
	Int interface {
		int | int8 | int16 | int32 | int64
	}
	// Uint constrains to only builtin unsigned integer types.
	Uint interface {
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
	}
	// Float constrains to only builtin float types.
	Float interface {
		float32 | float64
	}
	// Complex constrains to only builtin complex types.
	Complex interface {
		complex64 | complex128
	}
	// Number constrains to only builtin types supporting arithmetic.
	Number interface {
		Int | Uint | Float
	}
)

// Min returns the smaller of the two numbers.
func Min[T Number](a, b T) T {
	if a > b {
		return b
	}
	return a
}

// Max returns the larger of the two numbers.
func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Sum returns the sum of all numbers passed.
func Sum[T Number](numbers ...T) (sum T) {
	for _, num := range numbers {
		sum = sum + num
	}
	return
}
