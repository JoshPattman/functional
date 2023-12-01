package functional

import (
	"errors"
	"testing"
)

func TestPair(t *testing.T) {
	p := NewPair(1, 2)
	if First(p) != 1 || Second(p) != 2 {
		t.Errorf("Pair failed with pair %v", p)
	}
	x, y := Split(p)
	if x != 1 || y != 2 {
		t.Errorf("Split failed with pair %v", p)
	}
}

func TestPairs(t *testing.T) {
	xs := []int{1, 2, 3}
	ys := []int{4, 5, 6}
	zipped := Zip(xs, ys)
	if len(zipped) != 3 || First(zipped[0]) != 1 || Second(zipped[0]) != 4 || First(zipped[1]) != 2 || Second(zipped[1]) != 5 || First(zipped[2]) != 3 || Second(zipped[2]) != 6 {
		t.Errorf("Zip failed with lists %v and %v", xs, ys)
	}
}

func TestMapPairs(t *testing.T) {
	xs := []int{1, 2, 3}
	ys := []int{4, 5, 6}

	zs := Map(Zip(xs, ys), func(p Pair[int, int]) int { return p.First + p.Second })
	if len(zs) != 3 || zs[0] != 5 || zs[1] != 7 || zs[2] != 9 {
		t.Errorf("Map failed with list %v", zs)
	}
}

func TestMapPairsErrors(t *testing.T) {
	xs := []int{1, 2, 3}
	ys := []int{4, 5, 6}

	zs := Map(Zip(xs, ys), func(p Pair[int, int]) int { return p.First + p.Second })

	zsWithErrs := Map(zs, func(v int) Pair[int, error] {
		if v == 7 {
			return NewSecondPair[int, error](errors.New("The number was 7!"))
		}
		return NewFirstPair[int, error](v)
	})

	errs := Filter(Map(zsWithErrs, Second), IsNot[error](nil))
	if len(errs) != 1 || errs[0].Error() != "The number was 7!" {
		t.Errorf("Map failed with list %v", errs)
	}
}
