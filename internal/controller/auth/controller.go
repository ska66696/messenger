package auth

import (
	"encoding/json"
	"net/http"

	"messenger/internal/usecase/authsvc"
)

type AuthController struct {
	authUseCase authsvc.UseCase
}

func NewAuthController(useCase authsvc.UseCase) *AuthController {
	return &AuthController{
		authUseCase: useCase,
	}
}

func (c *AuthController) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		c.writeJSONError(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Pasword  string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		c.writeJSONError(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	user, err := c.authUseCase.ResterUser(requestBody.Username, requestBody.Email, requestBody.Pasword)
	if err != nil {
		c.writeJSONError(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := map[string]interface{}{
		"status":  "success",
		"message": "Пользователь успешно зарегестрирован",
		"user_id": user.ID,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (c *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		c.writeJSONError(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		Username string `json:"username"`
		Pasword  string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		c.writeJSONError(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	user, err := c.authUseCase.LoginUser(requestBody.Username, requestBody.Pasword)
	if err != nil {
		c.writeJSONError(w, err.Error(), http.StatusUnauthorized)
		return
	}
	response := map[string]interface{}{
		"status":   "success",
		"message":  "Вход выполнен успешно",
		"user_id":  user.ID,
		"username": user.Username,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (c *AuthController) writeJSONError(w http.ResponseWriter, message string, status int) {
	response := map[string]interface{}{
		"status":  "error",
		"message": message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
