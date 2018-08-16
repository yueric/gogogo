package main

import "fmt"

type Humaner interface {
	SayHi()
}

type student struct {
	name string
}
type teacher struct {
	name string
}

func (s student) SayHi(){
	fmt.Println("hello student ",s.name)
}

func (t teacher) SayHi(){
	fmt.Println("hello teacher ",t.name)
}

func whoSayHi(h Humaner){
	h.SayHi()
}
func main(){

	var stu,tea Humaner
	stu = student{"tom"}
	tea = teacher{"jerry"}
	//方法调用
	stu.SayHi()
	tea.SayHi()

	//多态
	whoSayHi(stu)
	whoSayHi(tea)
	//
	////接口赋值
	hs := make([]Humaner,2)
	hs[0],hs[1] = stu,tea
	for _,value := range hs{
		value.SayHi()
	}
}