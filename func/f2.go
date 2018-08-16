package main
import "fmt"
func Factorial(n int64)(result int64) {
	//var result int64
	if (n > 0) {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
func main() {
	var i int64 = 30
	fmt.Printf("%d的阶乘是：%d\n", i, Factorial(i))
	var a int8 =5
	var b int32 =15
	fmt.Println(b-a)
	}
