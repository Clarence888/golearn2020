package main

import "sync"

//示例：

func main() {
	sizes := make(chan int64)

	var wg sync.WaitGroup //
	for f := range xxxxxx {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//xxxxx
			//xxx

		}()
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
