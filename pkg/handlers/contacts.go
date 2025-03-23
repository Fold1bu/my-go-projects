package handlers

import (
	"encoding/json"
	"go-project/pkg/models"
	"go-project/pkg/repository"
	"net/http"
	"log"
)

   // handlers.go

   func PostContact(repo *repository.PGRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		// Декодируем JSON из тела запроса
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		// Сохраняем пользователя в БД
		if err := repo.CreateUser(&user); err != nil {
			log.Printf("SQL Error: %v\n", err)
			http.Error(w, "Error saving user to database", http.StatusInternalServerError)
			return
		}

		// Успешный ответ
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user) // Возвращаем созданного пользователя
	}
}


