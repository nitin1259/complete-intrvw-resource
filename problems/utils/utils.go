package utils

// Comparable Interface: This interface defines a type constraint.
// The ~ operator indicates that the underlying type must be one of the listed types (integers, floats, or strings).
// This allows the Max function to work with any custom type that has an underlying type of int, float64, string, etc.
type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

func Max[T Comparable](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Comparable](a, b T) T {
	if a < b {
		return a
	}
	return b
}
