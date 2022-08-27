package main

import (
	"fmt"
	"net/http"
)

func StartDbFunction(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("appLogin: %s\nappPassword: %s\nappIp: %s\nappPort: %s\nappDBname: %s\n", APP_LOGIN, APP_PASSWORD, APP_IP, APP_PORT, APP_DBNAME)
	ConnectDatabase()
	CreateTableNews()
	CreateDBproducts()
	fmt.Fprintln(writer, "База успешно создана")
}
