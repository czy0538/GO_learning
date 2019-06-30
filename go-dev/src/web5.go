package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	if "POST" == r.Method { //如果提交了form
		//接收文件 file：文件句柄，Handler：文件的信息,
		file, handler, err := r.FormFile("file")
		if err != nil { //获取上传文件错误
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Println(handler.Header) //输出消息头部
		defer file.Close()
		f, err := os.OpenFile("./"+handler.Filename,
			os.O_WRONLY|os.O_CREATE, os.ModePerm)
		// ./ 当前路径 即当前项目的GOPATH
		//O_WRONLY以只写方式打开 WR==WRITE
		//O_CREATE 创建文件
		//os.Modeperm 每个人都有读和写以及执行的权限 0777
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		size, err := io.Copy(f, file)
		//io.Copy适合于实时传输占用资源大
		//可以用ioutil.ReadAll
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "上传文件的大小为: %d", size)
		return
	}
	// 上传页面
	//向http的响应头中加入内容
	w.Header().Add("Content-Type", "text/html;charset=UTF-8")
	//根据HTTP State Code来写Response的Header
	w.WriteHeader(200) //写入http的状态码到Header
	//200 代表请求已成功，请求所希望的响应头或数据体将随此响应返回。
	html := `
	<form enctype="multipart/form-data" action="/" method="POST">
	请选择要上传的文件:	<input name="file" type="file" /><br/>
	<input type="submit" value="上传文件" />
	</form>`
	io.WriteString(w, html) //将html串写入w

}

/*
	type ResponseWriter interface {
	    // Header returns the header map that will be sent by WriteHeader.
	    // Changing the header after a call to WriteHeader (or Write) has
	    // no effect.
	    Header() Header

	    // Write writes the data to the connection as part of an HTTP reply.
	    // If WriteHeader has not yet been called, Write calls WriteHeader(http.StatusOK)
	    // before writing the data.  If the Header does not contain a
	    // Content-Type line, Write adds a Content-Type set to the result of passing
	    // the initial 512 bytes of written data to DetectContentType.
	    Write([]byte) (int, error)

	    // WriteHeader sends an HTTP response header with status code.
	    // If WriteHeader is not called explicitly, the first call to Write
	    // will trigger an implicit WriteHeader(http.StatusOK).
	    // Thus explicit calls to WriteHeader are mainly used to
	    // send error codes.
	    WriteHeader(int)
	}

*/
