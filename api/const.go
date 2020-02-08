package main

import (
	"encoding/json"
)

const baseUrl string = ":69"
const serverName string = "localhost"
const databaseName string = "1000Salesman"
const usernameDB string = "sa"
const passwordDB string = "abc123"


func EncodeJSON(enc *json.Encoder, v interface{}) {
	_ = enc.Encode(v)
}