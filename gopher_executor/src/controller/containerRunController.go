package controller

import (
	"../DAO/database/entities"
	"encoding/json"
	"log"
	"net/http"
)

func RunContainer(w http.ResponseWriter, r *http.Request) {
	log.Println("Running container.")
	payload := &Payload{nil, nil}
	status := http.StatusAccepted
	defer WriteJsonResponse(w, payload, status)
	containerRun := &entities.ContainerRunEntity{}
	if err := json.NewDecoder(r.Body).Decode(containerRun); err != nil {
		log.Println("Error ")
		status = http.StatusBadRequest
		payload.Error = err
		return
	}
	log.Println("Now to figure out how to run a docker container")
}
