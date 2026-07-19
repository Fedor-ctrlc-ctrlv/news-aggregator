package main

import (
	"fmt"
	"net/http"
	"log"
	"invest-news/internal/storage"
	_ "github.com/gorilla/mux"

)

func main() {
	//router := mux.NewRouter()
	db,err:=storage.NewStorage()
	if err!=nil{
		log.FatalF("Ошибка в подключении к бд %v", err)
	}
	defer db.Close()
	
	fmt.Println("Server six seven")
	http.ListenAndServe(":6767", nil)
}
