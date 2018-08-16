package main

import "fmt"

func main() {
	c := make(chan int,2)

	go func() {
		defer fmt.Println("goroutine结束")

		fmt.Println("goroutine正在运行……")

		c <- 666 //666发送到c
		c <- 777
	}()

	num := <-c //从c中接收数据，并赋值给num

	fmt.Println("num = ", num)

	num = <-c
	fmt.Println("num2 = ", num)


	fmt.Println("main结束")
}

