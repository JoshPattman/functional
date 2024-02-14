package functional

// Pair is a tuple of two values.
type Pair[T, U any] struct {
	First  T
	Second U
}

// Zip combines two lists into a list of pairs.
func Zip[T, U any](xs []T, ys []U) []Pair[T, U] {
	length := max(len(xs), len(ys))
	zipped := make([]Pair[T, U], length)
	for i := range zipped {
		zipped[i] = Pair[T, U]{xs[i], ys[i]}
	}
	return zipped
}

// ZipPair combines two lists (stored in a pair) into a list of pairs.
func ZipPair[T, U any](xys Pair[[]T, []U]) []Pair[T, U] {
	return Zip(xys.First, xys.Second)
}

// Unzip separates a list of pairs into two lists.
func Unzip[T, U any](zipped []Pair[T, U]) ([]T, []U) {
	xs := make([]T, len(zipped))
	ys := make([]U, len(zipped))
	for i, p := range zipped {
		xs[i] = p.First
		ys[i] = p.Second
	}
	return xs, ys
}

// UnzipPair separates a list of pairs into two lists (stored in a pair).
func UnzipPair[T, U any](zipped []Pair[T, U]) Pair[[]T, []U] {
	return NewPair(Unzip(zipped))
}

// First gets the first element of a pair.
func First[T, U any](p Pair[T, U]) T {
	return p.First
}

// Second gets the second element of a pair.
func Second[T, U any](p Pair[T, U]) U {
	return p.Second
}

// MustPair returns the first element if the second element is nil, otherwise it panics with the second element.
func MustPair[T any](p Pair[T, error]) T {
	if p.Second != nil {
		panic(p.Second)
	}
	return p.First
}

// NewPair creates a pair.
func NewPair[T, U any](x T, y U) Pair[T, U] {
	return Pair[T, U]{x, y}
}

// NewFirstPair creates a pair with the first element set.
func NewFirstPair[T, U any](x T) Pair[T, U] {
	var y U
	return Pair[T, U]{x, y}
}

// NewSecondPair creates a pair with the second element set.
func NewSecondPair[T, U any](y U) Pair[T, U] {
	var x T
	return Pair[T, U]{x, y}
}

// Split separates a pair into two elements.
func Split[T, U any](p Pair[T, U]) (T, U) {
	return p.First, p.Second
}
