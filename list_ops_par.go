package functional

import (
	"sync"
)

// PDo applies a function to each element of a slice in parallel.
func PDo[T any](xs []T, f func(T), workers int) {
	inputChan := make(chan int, len(xs))
	wg := sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		go func() {
			for {
				index, ok := <-inputChan
				if !ok {
					return
				}
				f(xs[index])
				wg.Done()
			}
		}()
	}
	for xi := range xs {
		wg.Add(1)
		inputChan <- xi
	}
	close(inputChan)
	wg.Wait()
}

// PMap applies a function to each element of a slice in parallel and returns a new slice with the results.
func PMap[T, U any](xs []T, f func(T) U, workers int) []U {
	inputChan := make(chan int, len(xs))
	ys := make([]U, len(xs))
	wg := sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		go func() {
			for {
				index, ok := <-inputChan
				if !ok {
					return
				}
				ys[index] = f(xs[index])
				wg.Done()

			}
		}()
	}
	for xi := range xs {
		wg.Add(1)
		inputChan <- xi
	}
	close(inputChan)
	wg.Wait()
	return ys
}

// PMake creates a slice of length n by applying a function to each index in parallel.
func PMake[T any](n int, f func() T, workers int) []T {
	inputChan := make(chan int, n)
	ys := make([]T, n)
	wg := sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		go func() {
			for {
				index, ok := <-inputChan
				if !ok {
					return
				}
				ys[index] = f()
				wg.Done()

			}
		}()
	}
	for xi := 0; xi < n; xi++ {
		wg.Add(1)
		inputChan <- xi
	}
	close(inputChan)
	wg.Wait()
	return ys
}

// PMake creates a slice of length n by applying a function to each index in parallel.
func PMakeIdx[T any](n int, f func(int) T, workers int) []T {
	inputChan := make(chan int, n)
	ys := make([]T, n)
	wg := sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		go func() {
			for {
				index, ok := <-inputChan
				if !ok {
					return
				}
				ys[index] = f(index)
				wg.Done()

			}
		}()
	}
	for xi := 0; xi < n; xi++ {
		wg.Add(1)
		inputChan <- xi
	}
	close(inputChan)
	wg.Wait()
	return ys
}

// PFilter filters a slice in parallel using a function and returns a new slice with the results.
func PFilter[T any](xs []T, f func(T) bool, workers int) []T {
	inputChan := make(chan int, len(xs))
	ys := make([]T, 0, len(xs))
	wg := sync.WaitGroup{}
	for i := 0; i < workers; i++ {
		go func() {
			for {
				index, ok := <-inputChan
				if !ok {
					return
				}
				if f(xs[index]) {
					ys = append(ys, xs[index])
				}
				wg.Done()

			}
		}()
	}
	for xi := range xs {
		wg.Add(1)
		inputChan <- xi
	}
	close(inputChan)
	wg.Wait()
	return ys
}
