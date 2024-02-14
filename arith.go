package functional

func Sum[T Number](vs []T) T {
	sum := T(0)
	for _, v := range vs {
		sum += v
	}
	return sum
}

func Prod[T Number](vs []T) T {
	prod := T(1)
	for _, v := range vs {
		prod *= v
	}
	return prod
}

func Mean[T Number](vs []T) T {
	return Sum(vs) / T(len(vs))
}

func Max[T Number](vs []T) T {
	max := vs[0]
	for _, v := range vs {
		if v > max {
			max = v
		}
	}
	return max
}

func Min[T Number](vs []T) T {
	min := vs[0]
	for _, v := range vs {
		if v < min {
			min = v
		}
	}
	return min
}

func ArgMax[T Number](vs []T) int {
	index := 0
	max := vs[0]
	for i, v := range vs {
		if v > max {
			max = v
			index = i
		}
	}
	return index
}

func ArgMin[T Number](vs []T) int {
	index := 0
	min := vs[0]
	for i, v := range vs {
		if v < min {
			min = v
			index = i
		}
	}
	return index
}
