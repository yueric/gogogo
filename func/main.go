package main

import "fmt"

func swap1(a1, b1 int) { //值传递
	a1, b1 = b1, a1
}
func swap2(c1, d1 *int) { //引用传递：c1、d1为指针变量
	*c1, *d1 = *d1, *c1 //不能交换地址，即：c1, d1 = d1, c1 错误
}
func main() {
	var a int = 2
	var b int = 1
	swap1(a, b) //a、b的值传递给a1、b1
	fmt.Printf("值传递：a = %d, b = %d\n", a, b)
	var c int = 2
	var d int = 1
	swap2(&c, &d) //c、d的地址传递给c1、d1
	fmt.Printf("指针传递：c = %d, d = %d\n", c, d)
}
