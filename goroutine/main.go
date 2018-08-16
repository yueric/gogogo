package main

import (
	"fmt"
	"time"
	"runtime"
)

func main(){

	n := runtime.GOMAXPROCS(2)

	fmt.Println("i am here",n)

	go func(){
		fmt.Println("go routine1")
	}()
	go func(){
		fmt.Println("go routine2")
	}()
	go func(){
		fmt.Println("go routine3")
	}()
	go func(){
		fmt.Println("go routine4")
	}()

	//n = runtime.GOMAXPROCS(3)
	n = runtime.GOMAXPROCS(3)

	fmt.Println("i am there",n)

	fmt.Println("i am there")
	time.Sleep(2 * time.Second)
}

