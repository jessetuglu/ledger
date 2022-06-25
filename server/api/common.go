package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)


type ErrorMessage struct {
	Message string `json:"message"`
}

type StatusError struct {
	status  int
	message string
}

type Error func (w http.ResponseWriter, r *http.Request) error


func (s *Server) internalServerError(w http.ResponseWriter, r *http.Request) {
	s.sendResponse(
		http.StatusInternalServerError,
		ErrorMessage{
			Message: internalServerError(r).message,
		},
		w, r,
	)
}

func internalServerError(r *http.Request) StatusError {
	return StatusError{status: http.StatusInternalServerError,
		message: "Interval server error: " +
			r.Context().Value(RequestIdKey).(uuid.UUID).String()}
}

func (s *Server) sendResponse(code int, data interface{}, w http.ResponseWriter, r *http.Request) error {
	var body []byte
	if data != nil {
		w.Header().Add("Content-Type", "application/json")
		var err error
		body, err = json.Marshal(data)
		if err != nil {
			s.logger.Errorw("failed to marshal response",
			RequestIdKey, r.Context().Value(RequestIdKey),
				"err", err,
			)
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}
	}

	w.WriteHeader(code)
	_, err := w.Write(body)
	if err != nil {
		s.logger.Warnw("failed to write response to client",
		RequestIdKey, r.Context().Value(RequestIdKey),
			"err", err,
		)
	}
	return err
}