package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	connectSQL()
	router := mux.NewRouter()

	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/users/all", getData).Methods("GET")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(baseUrl, router))
}

func connectSQL() {
	db, err = sql.Open("sqlserver", "sqlserver://"+usernameDB+":"+passwordDB+"@"+serverName+"?database="+databaseName+"&connection+timeout=30")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if db != nil {
		fmt.Println("Connected")
	}
}

func getData(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var response http.Response
	var arrUser []User
	rows, err := db.Query("SELECT TOP 5 Name, Email, NoTelepon FROM tbl_mst_user WHERE Name IS NOT NULL AND Email IS NOT NULL")
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if rows != nil {
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.Fullname, &user.Email, &user.Phone); err != nil {
				fmt.Println(err.Error())
				return
			} else {
				arrUser = append(arrUser, user)
			}
		}
	}
	response.Status = "Success"
	response.StatusCode = 200
	if err = json.NewEncoder(writer).Encode(arrUser); err != nil {
		fmt.Println(err.Error())
		return
	}
}

func login(writer http.ResponseWriter, request *http.Request) {

}
