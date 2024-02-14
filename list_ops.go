package functional

// Do applies a function to each element of a list.
func Do[T any](xs []T, f func(T)) {
	for xi := range xs {
		f(xs[xi])
	}
}

// Map applies a function to each element of a list and returns a new list with the results.
func Map[T, U any](xs []T, f func(T) U) []U {
	ys := make([]U, len(xs))
	for xi := range xs {
		ys[xi] = f(xs[xi])
	}
	return ys
}

// Make creates a list of length n by applying a function to each index.
func Make[T any](n int, f func(int) T) []T {
	ys := make([]T, n)
	for i := 0; i < n; i++ {
		ys[i] = f(i)
	}
	return ys
}

// Accumulate applies a function to each element of a list and returns the accumulated result.
func Accumulate[T, U any](xs []T, acc U, f func(T, U) U) U {
	for _, x := range xs {
		acc = f(x, acc)
	}
	return acc
}

// Filter returns a new list containing only the elements of the original list for which the provided function returns true.
func Filter[T any](xs []T, f func(T) bool) []T {
	ys := make([]T, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}
