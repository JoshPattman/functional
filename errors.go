package functional

// These Musts are a bit janky but i think they are the most efficient way to do this for now

// Must returns the first value if the error is nil, otherwise it panics with the error.
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

// Must2 returns the first and second values if the error is nil, otherwise it panics with the error.
func Must2[T, U any](t T, u U, err error) (T, U) {
	if err != nil {
		panic(err)
	}
	return t, u
}

// Must3 returns the first, second and third values if the error is nil, otherwise it panics with the error.
func Must3[T, U, V any](t T, u U, v V, err error) (T, U, V) {
	if err != nil {
		panic(err)
	}
	return t, u, v
}

// Must4 returns the first, second, third and fourth values if the error is nil, otherwise it panics with the error.
func Must4[T, U, V, W any](t T, u U, v V, w W, err error) (T, U, V, W) {
	if err != nil {
		panic(err)
	}
	return t, u, v, w
}
