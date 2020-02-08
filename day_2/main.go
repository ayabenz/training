package main

import (
	"fmt"
	"github.com/salamander-labs/training/day_2/handler"
	"net/http"
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
