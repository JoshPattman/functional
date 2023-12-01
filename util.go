package functional

import "runtime"

func NumCPU() int {
	return runtime.NumCPU()
}
