package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/monirz/gojwt/response"
	"github.com/monirz/gojwt/utils"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusBadRequest, "invalid input", nil)
		return
	}

	email, ok := data["email"].(string)
	password, ok2 := data["password"].(string)

	if !ok && !ok2 {
		response.ResponseError(w, http.StatusBadRequest, "email or password is required", nil)
		return
	}

	// Call the data layer function to retrieve the user by email
	user, err := s.UserService.FindByEmail(email)
	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error retrieving user", nil)
		return
	}

	if user == nil {
		response.ResponseError(w, http.StatusUnauthorized, "User not found", nil)
		return
	}

	// Check if the provided password matches the user's stored password
	if !utils.CheckPasswordHash(password, user.Password) {
		response.ResponseError(w, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}

	accessTokenKey := utils.RandStr(10)
	refreshTokenKey := utils.RandStr(10)

	// Generate access and refresh tokens for the user
	accessToken, refreshToken, err := utils.GenerateTokens(email, accessTokenKey, refreshTokenKey)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusInternalServerError, "Error generating access token", nil)
		return
	}

	// Return tokens as JSON response
	tokens := map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	response.JSONResponse(w, "login successful", http.StatusOK, 0, tokens)
}
