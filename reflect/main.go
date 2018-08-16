package main

import (
	"reflect"
	"fmt"
)

type student struct{
	Name string
	age int
}

func (stu student) setName(nam string){
	stu.Name = nam
}

func main() {
	stu := student{Name: "zhangsan", age: 25}
	t := reflect.TypeOf(stu) //获取对象的类型名称
	fmt.Println("class name:", t.Name())
	fmt.Printf("%v \n",t)
	//获取stu对象的成员名称和值
	v := reflect.ValueOf(stu)
	fmt.Printf("%v \n",v)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println(field.Name, "type:", field.Type)
		fmt.Println(field.Name, "value:", v.Field(i))
	}
	//获取stu的方法
	for i := 0; i < t.NumMethod(); i++ {
		f := t.Method(i)
		fmt.Println(f.Name, f.Type)
	}
}