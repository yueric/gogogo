package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := [3]*int{&a[0], &a[1], &a[2]}
	fmt.Println(b)
}
