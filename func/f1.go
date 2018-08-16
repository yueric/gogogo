package main

import "fmt"

func Test1() {
	fmt.Println("无参无返回值")
}
func Test2(v1, v2 int) {
	fmt.Printf("有参(普通参数)无返回值:v1=%d, v2= %d\n", v1, v2)
}
func Test3(args ...int) { //不定参数，args本质是一个切片
	fmt.Println("有参(不定参数)无返回值:", args)
}
func Test4() (value int) {
	value = 250
	return value
}
func Test5() (a int, str string) {
	b := 250
	c := "傻子"
	return b, c
}
func MinAndMax(num1 int, num2 int) (min int, max int) {
	if num1 > num2 {
		min = num2
		max = num1
	} else {
		max = num2
		min = num1
	}
	return
}
func main() {
	Test1()
	Test2(11, 22)
	Test3()
	Test3(1)
	Test3(1, 2, 3, 4)
	v4 := Test4()
	fmt.Println("无参有返回值(一个返回值)：", v4)
	_, v5 := Test5()
	fmt.Println("无参有返回值(多个返回值)：", v5)
	min, max := MinAndMax(33, 22)
	fmt.Printf("有参有返回值：min = %d, max = %d\n", min, max)
}
