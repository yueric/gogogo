package main

import "fmt"

func squares() (func() int) {
	var x int
	fmt.Println("we chekc ",x)
	return func() int {
		x++

		return x*x
	}
}

func main(){
	f1 := squares()

	fmt.Println(squares()())
	fmt.Println(squares()())
	fmt.Println(squares()())
	fmt.Println(squares()())

	fmt.Println(f1())
	fmt.Println(f1())
	fmt.Println(f1())

}