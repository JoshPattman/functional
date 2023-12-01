package functional

func Is[T comparable](x T) func(T) bool {
	return func(y T) bool {
		return x == y
	}
}

func IsNot[T comparable](x T) func(T) bool {
	return func(y T) bool {
		return x != y
	}
}
