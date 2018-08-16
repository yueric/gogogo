package main

import (
	"math"
	"fmt"
)

type Point struct{
	X float64
	Y float64
}
func Distance(p,q Point) float64{
	return math.Hypot(p.X-q.X, q.Y-p.Y)
}
func (p Point) Distance(q Point) float64{
	return math.Hypot(p.X-q.X, q.Y-p.Y)
}
func (p Point) changeX(val float64){
	p.X = val
}

func (p *Point) changeY(val float64){
	p.Y = val
}
func main(){
	po := Point{12.1,1.2}
	qo := Point{33.11,23.1}
	fmt.Println(po.Distance(qo))
	fmt.Println(Distance(po,qo))

	distanceFromP := po.Distance
	fmt.Println(distanceFromP(qo))

	distance := Point.Distance
	fmt.Println(distance(po,qo))

	po.changeY(13)
	fmt.Println(po)
}

