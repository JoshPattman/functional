package functional

// Is returns a function that checks if a value is equal to the provided value.
func Is[T comparable](x T) func(T) bool {
	return func(y T) bool {
		return x == y
	}
}

// IsNot returns a function that checks if a value is not equal to the provided value.
func IsNot[T comparable](x T) func(T) bool {
	return func(y T) bool {
		return x != y
	}
}
