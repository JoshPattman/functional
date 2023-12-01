package functional

func Do[T any](xs []T, f func(T)) {
	for xi := range xs {
		f(xs[xi])
	}
}

func Map[T, U any](xs []T, f func(T) U) []U {
	ys := make([]U, len(xs))
	for xi := range xs {
		ys[xi] = f(xs[xi])
	}
	return ys
}

func Make[T any](n int, f func(int) T) []T {
	ys := make([]T, n)
	for i := 0; i < n; i++ {
		ys[i] = f(i)
	}
	return ys
}

func Accumulate[T, U any](xs []T, acc U, f func(T, U) U) U {
	for _, x := range xs {
		acc = f(x, acc)
	}
	return acc
}

func Filter[T any](xs []T, f func(T) bool) []T {
	ys := make([]T, 0, len(xs))
	for _, x := range xs {
		if f(x) {
			ys = append(ys, x)
		}
	}
	return ys
}
