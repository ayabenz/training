package handler

import (
	"fmt"
	"net/http"
)

func Hello(writer http.ResponseWriter, request *http.Request) {
	if _, err :=  writer.Write([]byte("This is Hello World")); err != nil {
		fmt.Println(err.Error())
	}
	return
}
