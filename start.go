package main

import "net/http"

func StartDbFunction(writer http.ResponseWriter, request *http.Request) {
	ConnectDatabase()
	CreateTableNews()
	CreateDBproducts()
}
