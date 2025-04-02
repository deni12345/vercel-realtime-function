package model

import "time"

type User struct {
	Email        string    `firestore:"email"`
	Username     string    `firestore:"username"`
	HashPassword string    `firestore:"hashPassword"` // Hash of the password
	Phone        string    `firestore:"phone"`
	PhotoURL     string    `firestore:"photoURL"`
	CreatedAt    time.Time `firestore:"createdAt"`
	LastLogin    time.Time `firestore:"lastLogin"`
}

func NewUser(name, email, hashPassword, phone, photoURL string) *User {
	return &User{
		Email:        email,
		Username:     name,
		HashPassword: hashPassword,
		CreatedAt:    time.Now(),
		LastLogin:    time.Now(),
		Phone:        phone,
		PhotoURL:     photoURL,
	}
}
