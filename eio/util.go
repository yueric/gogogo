package main
import (
	"strings"
	"io/ioutil"
	"fmt"
)

func main() {
	//NopCloser返回一个读取对象的ReadCloser接口
	//用于提供Close方法
	r := strings.NewReader("hello")
	rcl := ioutil.NopCloser(r)
	defer rcl.Close()

	//ReadAll读取所有数据
	r2 := strings.NewReader("1234567890")
	p, _ := ioutil.ReadAll(r2)
	fmt.Println(string(p))

	//读取目录下信息
	fileInfo, _ := ioutil.ReadDir("./")
	for _, v := range fileInfo {
		fmt.Println(v.Name())
	}

	//读取整个文件数据
	data, _ := ioutil.ReadFile("./1.txt")
	fmt.Println(string(data))

	//向指定文件写入数据，如果文件不存在，则创建文件，写入数据之前清空文件
	ioutil.WriteFile("./xxx.txt", []byte("hello,world"), 0655)

	//在当前目录下，创建一个以test为前缀的临时文件夹，并返回文件夹路径
	name, _ := ioutil.TempDir("./", "test")
	fmt.Println(name)

	//在当前目录下，创建一个以test为前缀的文件，并以读写模式打开文件，并返回os.File指针
	file, _ := ioutil.TempFile("./", "test")
	file.WriteString("写入字符串")
	file.Close()
}