package models

type User struct {
    ID     int    `json:"id"`          // ID пользователя
    Name   string `json:"user_name"`   // Имя пользователя
    Number string `json:"user_number"` // Номер телефона пользователя
}
