package Register

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Komal-0110/User-Authentication-Service/models"
	"github.com/Komal-0110/User-Authentication-Service/sqlite"
	"github.com/gorilla/mux"
)

type HTTPTransport struct {
	service Service
}

func NewHttpTransport(service Service) *HTTPTransport {
	return &HTTPTransport{
		service: service,
	}
}

type ErrResponse struct {
	Error string `json:"error"`
}

func (t *HTTPTransport) Register(w http.ResponseWriter, r *http.Request) {
	var user models.UserReq
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errRes := ErrResponse{
			Error: "invalid json",
		}

		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errRes); err != nil {
			fmt.Println("failed to encode message", err)
		}
	}

	if err := t.service.AddUser(r.Context(), user); err != nil {
		errRes := ErrResponse{
			Error: "internal server error",
		}

		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(errRes); err != nil {
			fmt.Println("failed to encode message", err)
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (t *HTTPTransport) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := t.service.GetUsers(r.Context())
	if err != nil {
		errRes := ErrResponse{
			Error: "internal server error",
		}

		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(errRes); err != nil {
			fmt.Println("failed to encode message", err)
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		fmt.Println("failed to encode message", err)
	}
}

func (t *HTTPTransport) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["userID"]
	userID, err := strconv.Atoi(id)
	if err != nil {
		errRes := ErrResponse{
			Error: "invalid user id",
		}

		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errRes); err != nil {
			fmt.Println("failed to encode message", err)
		}
	}

	user, err := t.service.GetUserById(r.Context(), userID)
	if err != nil {
		if errors.Is(err, sqlite.ErrNotFound) {
			errRes := ErrResponse{
				Error: "user id not exists",
			}

			w.WriteHeader(http.StatusNotFound)
			if err := json.NewEncoder(w).Encode(errRes); err != nil {
				fmt.Println("failed to encode message", err)
			}
		}

		errRes := ErrResponse{
			Error: "internal server error",
		}

		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(errRes); err != nil {
			fmt.Println("failed to encode message", err)
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		fmt.Println("failed to encode message", err)
	}
}

func (t *HTTPTransport) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.UpdateUserReq
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		errRes := ErrResponse{
			Error: "invalid json",
		}

		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errRes); err != nil {
			fmt.Println("failed to encode message", err)
		}
	}

	if err := t.service.UpdateUser(r.Context(), user); err != nil {
		if errors.Is(err, sqlite.ErrNotFound) {
			errRes := ErrResponse{
				Error: "user not found",
			}

			w.WriteHeader(http.StatusBadRequest)
			if err := json.NewEncoder(w).Encode(errRes); err != nil {
				fmt.Println("failed to encode message", err)
			}
		}
		errRes := ErrResponse{
			Error: "internal server error",
		}

		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(errRes); err != nil {
			fmt.Println("failed to encode message", err)
		}
	}

	w.WriteHeader(http.StatusOK)
}

func (t *HTTPTransport) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["userID"]
	userID, err := strconv.Atoi(id)
	if err != nil {
		errRes := ErrResponse{
			Error: "invalid user id",
		}

		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errRes); err != nil {
			fmt.Println("failed to encode message", err)
		}
	}

	if err := t.service.DeleteUser(r.Context(), userID); err != nil {
		if errors.Is(err, sqlite.ErrNotFound) {
			errRes := ErrResponse{
				Error: "user id not exists",
			}

			w.WriteHeader(http.StatusNotFound)
			if err := json.NewEncoder(w).Encode(errRes); err != nil {
				fmt.Println("failed to encode message", err)
			}
		}

		errRes := ErrResponse{
			Error: "internal server error",
		}

		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(errRes); err != nil {
			fmt.Println("failed to encode message", err)
		}
	}

	w.WriteHeader(http.StatusOK)
}
