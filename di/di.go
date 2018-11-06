package main

import (
	"net/http"
	"io"
	"fmt"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	http.ListenAndServe(":8009", http.HandlerFunc(MyGreeterHandler))
}