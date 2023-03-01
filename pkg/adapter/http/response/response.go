package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type message struct {
	Message string `json:"message"`
}

func newMessage(msg string) *message {
	return &message{
		Message: msg,
	}
}

func New(w http.ResponseWriter, body interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Println(err.Error())
	}
}

func BadRequestErr(w http.ResponseWriter, err error) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	body := newMessage(err.Error())
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Println(err.Error())
	}
}
