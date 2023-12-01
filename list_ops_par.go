package functional

import (
	"sync"
)

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

func PMake[T any](n int, f func(int) T, workers int) []T {
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
