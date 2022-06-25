package api

import "github.com/google/uuid"

type Message struct {
	Message string `json:"message"`
}

type ErrorMessage struct {
	Message string `json:"error"`
}

type byUUIDRequest struct {
	Id uuid.UUID `json:"id" binding:"required"`
}
