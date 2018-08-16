package main
import "fmt"
func main() {
	language := "Python"
	languageType := "Static"
	m := language == "Go"
	n := languageType == "Static"
	if !m && n {
		language = "Go"
	}
	switch language {
	case "Python":
		fmt.Println("python language")
	case "Java":
		fmt.Println("java language")
	case "Go":
		fmt.Println("go language")
	default:
		fmt.Println("unknow language")
	}
}
