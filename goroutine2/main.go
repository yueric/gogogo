package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i = i + 1 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			//defer wg.Add(-1)
			EchoNumber(n)
		}(i)
	}

	wg.Wait()
	fmt.Println("End")
}

func EchoNumber(i int) {
	//time.Sleep(3e9)
	fmt.Println(i)
}
