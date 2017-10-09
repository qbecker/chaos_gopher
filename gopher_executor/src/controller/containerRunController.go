package controller

import (
	"log"
	"net/http"
)

func RunContainer(w http.ResponseWriter, r *http.Request) {
	log.Println("Running container.")
	payload := &Payload{nil, nil}
	status := http.StatusAccepted
	defer WriteJsonResponse(w, payload, status)
}
