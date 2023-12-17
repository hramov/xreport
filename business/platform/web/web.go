package web

import (
	"encoding/json"
	"net/http"
	"time"
)

type response struct {
	Dt     time.Time `json:"dt"`
	Status int       `json:"status"`
	Data   any       `json:"data"`
	Error  string    `json:"error"`
}

func SendResponse(w http.ResponseWriter, data any, status int) error {
	resp := response{
		Dt:     time.Now(),
		Status: status,
		Data:   data,
	}

	body, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	w.WriteHeader(status)
	_, err = w.Write(body)
	if err != nil {
		return err
	}
	return nil
}

func SendError(w http.ResponseWriter, err error, status int) error {
	resp := response{
		Dt:     time.Now(),
		Status: status,
		Error:  err.Error(),
	}

	body, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	w.WriteHeader(status)
	_, err = w.Write(body)
	if err != nil {
		return err
	}
	return nil
}
