package main

import (
	"fmt"
	"time"
)

func test(){
	fmt.Println("2222")
}
func main(){
	fmt.Println("1111")
	defer test()
	fmt.Println("33333")
	time.Sleep(3*time.Second)
	panic("66666")
	defer fmt.Println("44444")
	fmt.Println("555555")

}