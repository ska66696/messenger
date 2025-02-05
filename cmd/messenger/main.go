package main

import (
	"fmt"
	"net/http"

	"messenger/internal/repository/user"
	"messenger/internal/usecase/auth"
)

func main() {
	userRepo := user.NewInMemoryRepository()

	authService := auth.NewService(userRepo)

	registredUser, err := authService.ResterUser("testuser", "test@test.com", "pasword")
	if err != nil {
		fmt.Println("Ошибка регистрации:", err)
	} else {
		fmt.Println("Пользователь зарегистрирован:", registredUser)
	}

	foundUser, err := userRepo.GetByUsername("testuser")
	if err != nil {
		fmt.Println("Ошибка про получении пользователя:", err)
	} else {
		fmt.Println("Пользователь получен:", foundUser)
	}

	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", nil)
}
