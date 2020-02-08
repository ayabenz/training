package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	var address = ":8099"
	fmt.Println("Server Started")

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		var Data = map[string]string{
			"Title":   "Template",
			"Name":    "Benny",
			"Message": "Hello",
		}

		var t, err = template.ParseFiles("template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(writer, Data)
	})

	http.ListenAndServe(address, nil)
}
