package main

import (
	"fmt"
	"time"
)

func main(){
	c := make(chan int,0)
	chanSlice := []chan int{}
	for i := 1; i <= 11; i++{
		chanSlice = append(chanSlice, make(chan int))
	}
	for i:=0;i<10 ;i++  {
		go func(t int){
			<-c
			fmt.Println(t)
		}(i)
	}

	for j:=0;j<10 ;j++  {
		c<-j
	}

	time.Sleep(3*time.Second)
}
