package main

import "fmt"

func main() {
	m := map[int]string{1: "mike", 2: "yoyo"}
	value1, ok := m[1] //ok的结果有两种：true或value
	if ok {            //如果ok为true，执行括号内代码
		fmt.Printf("找到了：key = 1, value = %v\n", value1)
	}

	if value2, ok := m[3]; ok { //如果ok为true，执行此括号内代码
		fmt.Printf("找到了：key = 3, value = %v\n", value2)
	} else { //如果ok为false，执行此括号内代码
		fmt.Printf("没有找到！\n")
	}
}
