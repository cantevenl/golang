package main
//
//import (
//	"html/template"
//
//	"log"
//	"net/http"
//	"os"
//)
//
//func testTemplateStr() {
//	//数据
//	name := "zhu"
//	//定义模板
//	muban := "hello, {{.}}"
//	//解析模板
//	tmpl, err := template.New("test").Parse(muban)
//	if err != nil {
//		log.Fatal(err)
//	}
//	//执行模板，输出到终端
//	err = tmpl.Execute(os.Stdout, name)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//type Person struct {
//	Name string
//	Age  int
//}
//
//func testTemplateStruct() {
//	person := Person{"tom", 20}
//	muban := "hello, {{.Name}}, Your age {{.Age}}"
//	tmpl, err := template.New("test").Parse(muban)
//	if err != nil {
//		log.Fatal(err)
//	}
//	tmpl.Execute(os.Stdout, person)
//}
//
//func testTemplateHtml() {
//	// 请求处理函数
//	f := func(resp http.ResponseWriter, req *http.Request) {
//		t1, err := template.ParseFiles("test3.html")
//		if err != nil {
//			log.Fatal(err)
//		}
//		t1.Execute(resp,"hello kubernetes")
//	}
//	// 响应路径,注意前面要有斜杠 /
//	http.HandleFunc("/hello", f)
//	// 设置监听端口，并监听，注意前面要有冒号:
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func testTemplateRange() {
//	f := func(resp http.ResponseWriter,req *http.Request) {
//		files, err := template.ParseFiles("test5.html")
//		if err != nil {
//			log.Fatal(err)
//		}
//		s := []string{"牛", "golang ", "k8s","猪"}
//		files.Execute(resp, s)
//	}
//
//	server:=http.Server{
//		Addr: "127.0.0.1:8080",
//	}
//	http.HandleFunc("/test",f)
//	server.ListenAndServe()
//}
//
//func main() {
//	//testTemplateStr()
//	//testTemplateStruct()
//	//testTemplateHtml()
//	testTemplateRange()
//
//}
