package main

import (
	"fmt"
	"net/http"
)

var portName = "8081"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(5, 3)
	fmt.Fprintln(w, "This is the About Page! The sum is:", sum)
}

func addValues(a, b int) int {
	return a + b
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println("Server is running on port", portName)
	http.ListenAndServe(":"+portName, nil)
}
