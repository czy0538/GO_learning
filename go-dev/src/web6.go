package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		fmt.Println(err)
	}
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	strTemplate := "<div><h3>欢迎光临！</h3></div>"
	t := template.New("test") //建立一个模板
	//解析模板
	t, err := t.Parse(strTemplate) //模板内容为strTemplate
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(w, nil) //将内容输出到模板中，这里为nil
	if err != nil {
		fmt.Println(err)
	}
}
