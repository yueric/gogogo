package main

import "fmt"

func testRecover() {

	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()

	panic("error happend")
	fmt.Println("1111111")
}
func main() {
	testRecover()
	fmt.Println("I am  alive")
}
