package Register

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func (t *HTTPTransport) AddUser(w http.ResponseWriter, r *http.Request) {
	var user UserReq
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

		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errRes); err != nil {
			fmt.Println("failed to encode message", err)
		}
	}
}
