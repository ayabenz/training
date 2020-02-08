package handler

import (
	"fmt"
	"net/http"
)

func Welcome(writer http.ResponseWriter, request *http.Request) {
	if _, err :=  writer.Write([]byte("This is Welcome")); err != nil {
		fmt.Println(err.Error())
	}
	return
}