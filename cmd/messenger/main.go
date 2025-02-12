package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"

	"messenger/internal/controller/auth"
	// "messenger/internal/repository/chat"
	// "messenger/internal/repository/messange"
	"messenger/internal/repository/user"
	"messenger/internal/usecase/authsvc"
	// "messenger/internal/usecase/chatsvc"
	// "messenger/internal/usecase/usersvc"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLO!")
}

func main() {
	userRepo := user.NewInMemoryRepository()
	// chatRepo := chat.NewInMemoryRepository()
	// messangeRepo := messange.NewInMemoryRepository()

	authService := authsvc.NewServiceAuth(userRepo)
	// userService := usersvc.NewServiceUser(userRepo)
	// chatService := chatsvc.NewServiceChat(chatRepo, messangeRepo)
	authController := auth.NewAuthController(authService)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"null"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	registerHandler := c.Handler(http.HandlerFunc(authController.RegisterHandler))
	http.Handle("/register", registerHandler)

	loginHandler := c.Handler(http.HandlerFunc(authController.LoginHandler))
	http.Handle("/login", loginHandler)

	// fs := http.FileServer(http.Dir("./frontend"))
	// http.Handle("/", fs)

	http.HandleFunc("/", hello)

	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", nil)
}
