package main
import "fmt"
func main(){
	var m1 map[int]string
	fmt.Println(m1)
	fmt.Println(m1 == nil)     //仅声明的字典是空字典，打印的结果是：m1=map[]
	m1[1]="a"					//m1=map[int]string{}
	fmt.Println(m1)          //空字典是不能存储键值对的，程序出现panic
}
