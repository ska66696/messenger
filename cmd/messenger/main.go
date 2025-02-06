package main

import (
	"fmt"
	"net/http"

	"messenger/internal/repository/chat"
	"messenger/internal/repository/messange"
	"messenger/internal/repository/user"
	"messenger/internal/usecase/authsvc"
	"messenger/internal/usecase/chatsvc"
	"messenger/internal/usecase/usersvc"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO!")
}

func main() {
	userRepo := user.NewInMemoryRepository()
	chatRepo := chat.NewInMemoryRepository()
	messangeRepo := messange.NewInMemoryRepository()

	authService := authsvc.NewServiceAuth(userRepo)
	userService := usersvc.NewServiceUser(userRepo)
	chatService := chatsvc.NewServiceChat(chatRepo, messangeRepo)

	registredUser, err := authService.ResterUser("testuser", "test@test.com", "pasword")
	if err != nil {
		fmt.Println("Ошибка регистрации:", err)
	} else {
		fmt.Println("Пользователь зарегистрирован:", registredUser)
	}

	foundUser, err := userRepo.GetByUsername("testuser")
	if err != nil {
		fmt.Println("Ошибка при получении пользователя:", err)
	} else {
		fmt.Println("Пользователь получен:", foundUser)
	}

	loggedUser, err := authService.LoginUser("testuser", "pasword")
	if err != nil {
		fmt.Println("Ошибка при входе пользователя:", err)
	} else {
		fmt.Println("Пользователь вошел:", loggedUser)
	}

	foundUserByUsername, err := userService.FindUserByUsername("testuser")
	if err != nil {
		fmt.Println("Ошибка при поиске пользователя:", err)
	} else {
		fmt.Println("Пользователь найден по имени:", foundUserByUsername)
	}

	user1, _ := userService.FindUserByUsername("testuser")
	user2, err := authService.ResterUser("testuser2", "test2@test.com", "pasword2")

	chatUsers := []string{user1.ID, user2.ID}
	newChat, err := chatService.CreateChat(chatUsers)
	if err != nil {
		fmt.Println("Ошибка при создании чата:", err)
	} else {
		fmt.Println("чат успешно создан:", newChat)
	}

	sentMessage, err := chatService.SendMessage(newChat.ID, user1.ID, "Привет, testuser2!")
	if err != nil {
		fmt.Println("Ошибка при отправке сообщения:", err)
	} else {
		fmt.Println("Сообщение отправлено:", sentMessage)
	}

	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)

	// http.HandleFunc("/", hello)

	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8081", nil)
}
