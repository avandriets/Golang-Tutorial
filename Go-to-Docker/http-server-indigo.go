package main

/*
#cgo CFLAGS: -I./indigo/include
#include <indigo.h>

#include <stdio.h>
#include <stdlib.h>

static void myprint(char* s) {
  printf("%s\n", s);
}
*/
import "C"
import (
	"io"
	"net/http"
)

func helloWorld(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html><html><head><title>Hello World</title></head><body>Hello World!</body></html>`,
	)
}

func main() {
	http.HandleFunc("/hello", helloWorld)
	http.ListenAndServe(":9000", nil)
}
