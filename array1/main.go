package main

import "fmt"

func main(){
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println("a = ", a)
	for i := 0; i < 3; i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
	a[1] = 100
	fmt.Println("a = ", a)
	b := [5]int{2: 100, 4: 200}
	fmt.Println("b = ", b)
	for j := 0; j < 5; j++ {
		fmt.Printf("b[%d] = %d\n", j, b[j])
	}
	b[1] = 1
	fmt.Println("b = ",b)

}
