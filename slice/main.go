package main
import "fmt"
func main() {
	s :=[]int{1, 2, 3}                     //切片元素类型为整数型
	for i, v := range s{
		fmt.Printf("s[%d] = %d\n", i, v)
	}
	slice1:=append(s,1,2) //将同类型元素1、2追加到切片s的尾部，形成新切片slice1
	fmt.Println(slice1)
	slice2:=[]int{3,4}
	slice3:=append(s,slice2...)    //将slice2整个切片的元素都追加到切片s的尾部，形成新切片slice3
	fmt.Println(slice3)
}