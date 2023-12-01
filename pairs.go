package functional

type Pair[T, U any] struct {
	First  T
	Second U
}

func Zip[T, U any](xs []T, ys []U) []Pair[T, U] {
	length := max(len(xs), len(ys))
	zipped := make([]Pair[T, U], length)
	for i := range zipped {
		zipped[i] = Pair[T, U]{xs[i], ys[i]}
	}
	return zipped
}

func ZipPair[T, U any](xys Pair[[]T, []U]) []Pair[T, U] {
	return Zip(xys.First, xys.Second)
}

func Unzip[T, U any](zipped []Pair[T, U]) ([]T, []U) {
	xs := make([]T, len(zipped))
	ys := make([]U, len(zipped))
	for i, p := range zipped {
		xs[i] = p.First
		ys[i] = p.Second
	}
	return xs, ys
}

func UnzipPair[T, U any](zipped []Pair[T, U]) Pair[[]T, []U] {
	return NewPair(Unzip(zipped))
}

func First[T, U any](p Pair[T, U]) T {
	return p.First
}

func Second[T, U any](p Pair[T, U]) U {
	return p.Second
}

func NewPair[T, U any](x T, y U) Pair[T, U] {
	return Pair[T, U]{x, y}
}

func NewFirstPair[T, U any](x T) Pair[T, U] {
	var y U
	return Pair[T, U]{x, y}
}

func NewSecondPair[T, U any](y U) Pair[T, U] {
	var x T
	return Pair[T, U]{x, y}
}

func Split[T, U any](p Pair[T, U]) (T, U) {
	return p.First, p.Second
}
