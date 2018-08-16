package main
import "fmt"
func main() {
	s := []int{1,2,6,7,8}
	index := 2
	/*先将切片在该位置后面的全部元素都添加到一个空切片，形成一个新切片*/
	s2 := append([]int{},s[index:]...)
	/*将包含这个位置及该位置之前的所有元素截取，制成一个新切片s1 */
	s1 := append(s[:index])
	s1=append(s1,3,4,5)           //在s1后面添加想要插入的元素
	/*将s2中的所有元素都添加到s1中，形成的切片s就包含想要插入的元素*/
	s = append(s1, s2...)
	fmt.Printf("s=%v\n",s)
}