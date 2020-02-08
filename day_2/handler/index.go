package handler

import (
	"fmt"
	"net/http"
)

func Index(writer http.ResponseWriter, request *http.Request) {
	if _, err := writer.Write([]byte("This is Index")); err != nil {
		fmt.Println(err.Error())
	}
	return
}
