package main

import (
	"reflect"
	"fmt"
)
type Humaner interface {
	SayHi()
}
type MyInt int
type stud struct {
	name string
}
func (s stud) SayHi(){
	fmt.Println("hello student ",s.name)
}


func main(){
	var x  MyInt = 7
	v:=reflect.ValueOf(x)
	fmt.Println(v.Type())
	fmt.Println(v.Kind())

	var stu Humaner
	stu = stud{"eric"}
	y:=reflect.ValueOf(stu)
	fmt.Println(y.Type())
	fmt.Println(y.Kind())
}
