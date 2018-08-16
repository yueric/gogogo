package main

import (
	"runtime"
	"fmt"
)

func main(){
	n := runtime.GOMAXPROCS(4)
	fmt.Println(n)
	for  {
		go fmt.Print(0)
		go fmt.Print(1)
		go fmt.Print(2)
		fmt.Print(3)
	}
}