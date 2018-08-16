package main

import "fmt"
type class2 struct {
	string
	int
}
func main() {
	var c1 class2 = class2{"高一(一)班", 50 }
	fmt.Println("c1 =",c1)
	var c2 class2
	c2.string = "高二(三)班"
	c2.int = 30
	fmt.Println("c2 =",c2)
}
