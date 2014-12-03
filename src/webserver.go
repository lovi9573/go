package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
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

	fmt.Printf("FilePath: %s\nMinimum File Size: %s\n", request.PostForm["dir"][0], request.PostForm["minSize"][0])

	//Run external program and wait for it to finish
	cmd := exec.Command("sleep", "5")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}

	//Send browser the the generated page
	http.Redirect(output, request, "/index.html", 302)
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/submit", submit)
	http.ListenAndServe(":8080", nil)
}
