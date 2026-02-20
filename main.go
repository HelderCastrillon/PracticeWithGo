package main

import (
	"fmt"

	"github.com/HelderCastrillon/PracticeWithGo/internal/domain"
)

func main() {
	fmt.Println("Hola Mundo")

	user := domain.User{
		ID:       1,
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password123",
	}

	fmt.Println(user)
}
