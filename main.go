package main

import (
	"fmt"
	controller "go_connectdb/controllers"
	"net/http"
)

func main() {
	fmt.Println("Hello")
	http.HandleFunc("/api/user/", controller.GetUser)
	http.ListenAndServe("127.0.0.1:2222", http.DefaultServeMux)
}
