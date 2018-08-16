package main

import "fmt"

type Element interface{}

type Person struct {
	name string
	age  int
}

func main() {
	list := make([]Element, 3)
	list[0] = 1       // an int
	list[1] = "Hello" // a string
	list[2] = Person{"mike", 18}

	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is [%s, %d]\n", index, value.name, value.age)
		} else {
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}

}