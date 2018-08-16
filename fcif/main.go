package main

import "fmt"

func main() {
	var a int = 10 /*if语句*/
	if a > 5 {
		fmt.Println("a比5大")
	}
	/*if...else语句*/
	if a > 10 {
		fmt.Println("a比10大")
	} else {
		fmt.Println("a小于等于10")
	}
	/*if...else if...else语句*/
	if a > 10 {
		fmt.Println("a比10大")
	} else if a == 10 {
		fmt.Println("a等于10")
	} else {
		fmt.Println("a比10小")
	}
	/*if嵌套语句*/
	if a > 5 {
		if a == 10 {
			fmt.Printf("a=%d\n", a)
		}
	}
}
