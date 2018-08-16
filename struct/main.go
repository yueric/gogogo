package main

import "fmt"

type class struct {
	className    string
	number       int
	maleNumber   int
	femaleNumber int
}

func main() {
	/*1. 打印结构体成员变量值*/
	var c1 *class = &class{"高一(一)班", 50, 25, 25}
	fmt.Printf("className=%s,number=%d,maleNumber=%d,femaleNumber=%d\n", c1.className, (*c1).number, c1.maleNumber, (*c1).femaleNumber)
	/*2. 对结构体成员变量赋值或修改*/
	c2:=new(class)
	c2.className = "高三(一)班"
	(*c2).number=45
	c2.maleNumber=45
	//没有赋值的成员变量取零值
	fmt.Printf("*c2=%v\n",*c2)
	//显示成员变量名
	fmt.Printf("*c2=%+v\n",c2)
	//对成员变量值进行修改
	c2.className="高三(八)班"
	fmt.Printf("*c2=%+v\n",*c2)
}
