package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/monirz/gojwt"
	"github.com/monirz/gojwt/response"
	"github.com/monirz/gojwt/utils"
)

func (s *Server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := &gojwt.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusBadRequest, "invalid json input", err.Error())
		return
	}

	// Validate input
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		log.Println(err)

		response.ResponseError(w, http.StatusBadRequest, "validation error", err.Error())
		return
	}

	passHash, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println("err")
		log.Println(err)

		response.ResponseError(w, http.StatusBadRequest, "validation error", err.Error())
		return
	}

	user.UUID = uuid.NewString()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Password = passHash

	_, err = s.UserService.CreateUser(user)
	if err != nil {
		log.Println(err)
		response.ResponseError(w, http.StatusInternalServerError, "error creating user", err.Error())
		return
	}

	log.Println(err)
	response.JSONResponse(w, "user created successfully", http.StatusCreated, 201, &gojwt.User{
		UUID:      user.UUID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})
}
