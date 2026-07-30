package main

import (
	"encoding/json"
	"log"
	"net/http"
	"news_scroller/internal/storage"
	"news_scroller/internal/user"

	"github.com/gorilla/mux"
)

func main() {

	db, err := storage.NewStorage()
	if err != nil {
		log.Fatalf("Ошибка в подключении к бд %v", err)
	}
	defer db.Close()

	if err := db.Migrate(); err != nil {
		log.Fatalf("Ошибка в миграции %v", err)
	}

	userRepo := user.NewRepository(db.GetDB())

	router := mux.NewRouter()

	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		var params user.CreateUserParam
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, "Неверный формат", http.StatusBadRequest)
			return
		}
		if err := userRepo.Create(r.Context(), params); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)

	}).Methods("POST")

	log.Fatal(http.ListenAndServe(":6767", router))
}
