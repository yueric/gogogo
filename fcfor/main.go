package main

import "fmt"

func main() {
	/*goto:无条件的跳转到本函数内的某个标签，执行标签下面的语句*/
	var i, sum1 int
	for {
		i++
		if i > 5 {
			goto label
		}
		sum1 += i
	}
label:
	fmt.Printf("i = %d, sum1 = %d\n", i, sum1)

	/*break:跳出循环，并执行循环之后的语句*/
	var j, sum2 int
	for j = 1; j <= 10; j++ {
		if j > 5 {
			break
		}
		sum2 += j
	}
	fmt.Printf("j = %d, sum2 = %d\n", j, sum2)

	/*continue:跳过当前循环，并执行下一循环*/
	var k, sum3 int
	for k = 1; k <= 10; k++ {
		if k == 5 {
			continue
		}
		sum3 += k
	}
	fmt.Printf("k = %d, sum3 = %d\n", k, sum3)
}
