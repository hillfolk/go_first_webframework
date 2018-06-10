package main

import (
	"fmt"
	"net/http"
) 

func main() {
	// 루트 경로 함수 지정
	http.HandleFunc("/" ,func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"welcome!")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,"about")
	})

	http.ListenAndServe(":8080",nil)
}
