package functional

import "runtime"

// NumCPU returns the number of CPUs on the machine.
// Useful for setting the number of workers.
func NumCPU() int {
	return runtime.NumCPU()
}
