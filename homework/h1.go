package main

import (
	"net/http"
	"fmt"
	"log"
	"html"
	"strconv"
)

type student struct{
	name string
	birthday string
	class int
	no int
}
var stus map[int]student

func Indexhandler(w http.ResponseWriter,r *http.Request)  {
	defer func() {
		if err := recover(); err != nil {
			//log.Printf("Work failed with %s in %v", err, )
			fmt.Fprintf(w, "warning!!!! \n %v", err)
		}
	}()
	params := r.URL.Query()
	if len(params)>0 {
		stype := params["type"][0]
		switch stype {
		case "s"://创建
			cno, _ := strconv.Atoi(params["class"][0])
			sno, _ := strconv.Atoi(params["no"][0])
			nstu := student{params["name"][0], params["birthday"][0], cno, sno}
			stus[sno] = nstu
			fmt.Fprintf(w, "we create a new student: %v", nstu)
		case "g"://view
			sno, _ := strconv.Atoi(params["no"][0])
			fmt.Fprintf(w, "%v", stus[sno])
		case "list"://列表
			fmt.Fprintf(w, "%v", stus)
		default:
			fmt.Fprintln(w, "wrong command!")
		}
	}else{
		fmt.Fprintln(w, "hello world! pls use right url path.")
	}
}



func main(){

	stus = map[int]student{}

	http.HandleFunc("/",Indexhandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hello, %q, %v", html.EscapeString(r.URL.Path),r.URL.Query())

	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}