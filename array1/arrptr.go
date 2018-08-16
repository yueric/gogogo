package main
import "fmt"
func main() {
	a := [...]int{1, 2, 3}
	/*定义数组指针*/
	var p *[3]int = &a
	fmt.Printf("p=%p\n",p)
	fmt.Println("&a[0]=",&a[0])
	fmt.Println("&a[1]=",&a[1])
	fmt.Println("&a[2]=",&a[2])
	fmt.Println("*p=",*p, "a=",a)
	/*在数组指针这个变量前面加*可以获得该数组，再对数组进行遍历*/
	for index, value := range *p {
		fmt.Println(index ,":", value)
	}
	/*对数组指针遍历，也可以得到数组的index和value*/
	for index, value := range p {
		fmt.Println(index ,":",value)
	}
}
