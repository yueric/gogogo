package main

import (
	"fmt"
	"reflect"
)
type student struct{
	Name string
	Age uint
}

func (stu *student) SetName(nam string){
	stu.Name = nam
}
func main(){
	stu := student{Name: "zhangsan", Age: 25}

	//设置stu的age值为100
	v := reflect.ValueOf(&stu)
	v = v.Elem()
	field := v.FieldByName("Age")
	field.SetUint(100)

	//调用stu的setName方法
	v = reflect.ValueOf(&stu)
	m := v.MethodByName("SetName")
	args := []reflect.Value{reflect.ValueOf("liSi")}
	m.Call(args)
	fmt.Println(stu)

}