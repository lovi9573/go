package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func hello(output http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path)
	if(request.URL.Path == "/"){
		request.URL.Path = "/index.html"
	}
	file, error := os.Open("html"+request.URL.Path)
	if error != nil && error !=io.EOF {
		
		panic(error)
	}
	if(request.URL.Path[len(request.URL.Path)-3:] == "css"){
		header := output.Header()
		header.Set("Content-Type", "text/css")
	}
	buffer := make([]byte, 2048)
	
	for {
		n,error := file.Read(buffer)
		if error != nil && error !=io.EOF {
			panic(error)
		}
		if n == 0 {
			break
		}
		io.WriteString(output, string(buffer[:n]))
	}
}

func submit(output http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println(request.PostForm)
	http.Redirect(output, request, "/quotanizer/index.html", 1);
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/submit", submit)
	http.ListenAndServe(":8080", nil)
}
