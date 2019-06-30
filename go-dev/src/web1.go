package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/test", processReq)
	http.ListenAndServe(":8888", nil)
}
func processReq(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("<h1>第一个WEB应用</h1>"))
	rw.Write([]byte(req.URL.Path))
}
