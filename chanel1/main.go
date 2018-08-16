package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	c := make(chan int, 0) //无缓冲的通道

	//内置函数 len 返回未被读取的缓冲元素数量， cap 返回缓冲区大小
	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))

	go func() {
		defer fmt.Println("goroutine结束")
		for i := 0; i < 3; i++ {
			c <- i
			//fmt.Printf("goroutine正在运行[%d]: len(c)=%d, cap(c)=%d\n", i, len(c), cap(c))
			fmt.Println(i)
		}
	}()


	go func() {
		defer fmt.Println("goroutine2结束")
		for i := 0; i < 3; i++ {
			num := <-c //从c中接收数据，并赋值给num
			fmt.Println("num = ", num)
		}
	}()
	time.Sleep(2 * time.Second) //延时2s
	fmt.Println("main结束")
}

