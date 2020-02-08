package main

import (
	"fmt"
	"net/http"
	"training/day_2/handler"
)

func main() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/welcome", handler.Welcome)
	http.HandleFunc("/hello", handler.Hello)
	if err := http.ListenAndServe("localhost:69", nil); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Server Started")
	}
}
