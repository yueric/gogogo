package main

import (
	"fmt"
	"math/rand"
)

const (
	Left = iota
	Right
	GoOn
	Arrive =40
)

type car  struct{
	key uint
	isStarted bool
}

func (c *car) start(key uint){
	if key == c.key{
		c.isStarted = true
		fmt.Println("car started")
	}
}

func (c *car)goOn(){
	fmt.Println("go on")
}

func (c *car) back(){
	fmt.Println("go back")
}

func (c *car)turnLeft(){
	fmt.Println("turn left")
}
func (c *car)turnRight(){
	fmt.Println("turn right")
}
func (c *car)CarStop(){
	fmt.Println("arrive")
}
type director struct{

}

func (d *director) getNextDirect() int{
	v := rand.Intn(50)
	if v!=40 {
		v = v%3
	}
	return v
}




type person struct{
	name string
	key uint
	c *car
	d *director
}

func (p *person) dirveCar(){
	p.c.start(p.key)
	if p.c.isStarted{
		for{
			v := p.d.getNextDirect()
			switch v {
			case GoOn:
				p.c.goOn()
			case Left:
				p.c.turnLeft()
			case Right:
				p.c.turnRight()
			case Arrive:
				p.c.CarStop()
				return
			}
		}
	}
}

func main(){
	xiaoming := person{name:"xiaoming"}
	car := car{100,false}
	xiaoming.c = &car
	xiaoming.key = 100
	d := director{}
	xiaoming.d = &d
	xiaoming.dirveCar()

}