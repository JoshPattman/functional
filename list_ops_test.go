package functional

import "testing"

type testStruct struct {
	data int
}

func TestDo(t *testing.T) {
	vs := []*testStruct{{1}, {2}, {3}}
	Do(vs, func(v *testStruct) {
		v.data += 1
	})
	if vs[0].data != 2 || vs[1].data != 3 || vs[2].data != 4 {
		t.Errorf("Do failed with list %v", vs)
	}
}

func TestPDo(t *testing.T) {
	vs := []*testStruct{{1}, {2}, {3}}
	PDo(vs, func(v *testStruct) {
		v.data += 1
	}, NumCPU())
	if vs[0].data != 2 || vs[1].data != 3 || vs[2].data != 4 {
		t.Errorf("PDo failed with list %v", vs)
	}
}

func TestMap(t *testing.T) {
	vs := []*testStruct{{1}, {2}, {3}}
	us := Map(vs, func(v *testStruct) int {
		return v.data
	})
	if us[0] != 1 || us[1] != 2 || us[2] != 3 {
		t.Errorf("Map failed with list %v", vs)
	}
}

func TestPMap(t *testing.T) {
	vs := []*testStruct{{1}, {2}, {3}}
	us := PMap(vs, func(v *testStruct) int {
		return v.data
	}, NumCPU())
	if us[0] != 1 || us[1] != 2 || us[2] != 3 {
		t.Errorf("Map failed with list %v", vs)
	}
}

func TestMake(t *testing.T) {
	vs := Make(3, func(i int) int {
		return i
	})
	if len(vs) != 3 || vs[0] != 0 || vs[1] != 1 || vs[2] != 2 {
		t.Errorf("Make failed with list %v", vs)
	}
}

func TestPMake(t *testing.T) {
	vs := PMake(3, func(i int) int {
		return i
	}, NumCPU())
	if len(vs) != 3 || vs[0] != 0 || vs[1] != 1 || vs[2] != 2 {
		t.Errorf("Make failed with list %v", vs)
	}
}

func TestAccumulate(t *testing.T) {
	vs := []int{1, 2, 3}
	sum := Accumulate(vs, 0, func(x, acc int) int { return x + acc })
	if sum != 6 {
		t.Errorf("Accumulate failed with list %v", vs)
	}
}

func TestFilter(t *testing.T) {
	vs := []int{1, 2, 3}
	us := Filter(vs, func(x int) bool { return x > 1 })
	if len(us) != 2 {
		t.Errorf("Filter failed with list %v", vs)
	}
}

func TestPFilter(t *testing.T) {
	vs := []int{1, 2, 3}
	us := PFilter(vs, func(x int) bool { return x > 1 }, NumCPU())
	if len(us) != 2 {
		t.Errorf("Filter failed with list %v", vs)
	}
}
