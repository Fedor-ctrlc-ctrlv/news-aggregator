package main

import (
	"fmt"
	"net/http"
	"log"
	"news_scroller/internal/storage"
	_ "github.com/gorilla/mux"

)

func main() {
	//router := mux.NewRouter()
	db,err:=storage.NewStorage()
	if err!=nil{
		log.Fatalf("Ошибка в подключении к бд %v", err)
	}
	defer db.Close()
	
	if err:= db.Migrate();err!=nil{
		log.Fatalf("Ошибка в миграции %v", err)
	}
	fmt.Println("Server six seven")
	http.ListenAndServe(":6767", nil)
}
