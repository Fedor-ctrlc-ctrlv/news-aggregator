package main

import (
	"fmt"
	"net/http"

	_ "github.com/gorilla/mux"
)

func main() {
	//router := mux.NewRouter()
	fmt.Println("Server six seven")
	http.ListenAndServe(":6767", nil)
}
