package main

import (
	"fmt"
	"net/http"
	"shulyakav/goweb/handlers"
)

var portName = "8081"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println("Server is running on port", portName)
	http.ListenAndServe(":"+portName, nil)
}
