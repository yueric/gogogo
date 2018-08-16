package main
import "fmt"
func main() {
	s1:=[]int{1,2}
	s2:=[]int{3}
	copy(s1,s2)             //len(s1)>len(s2)，只覆盖s1前1个元素，余下元素保持不变
	fmt.Printf("s1=%v\n",s1)

	s3:=[]int{3}
	s4:=[]int{1,2}
	rt := copy(s1,s4)             //len(s3)<len(s4)，s4中只有一个元素能覆盖s3
	fmt.Printf("s3=%v\n",s3)
	fmt.Println(rt)
}
