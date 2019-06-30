package main

import "net/http"

type HttpHandler struct { //定义一个空类HttpHandler
}

//实现接口Handler中的ServeHTTP
func (this *HttpHandler) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	w.Write([]byte("<h1>在ServeHTTP 里</h1>"))
	w.Write([]byte(r.URL.Path))
}
func main() {
	handler := &HttpHandler{}
	http.ListenAndServe(":8899", handler)
}
