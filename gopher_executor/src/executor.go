package main

import (
	"./controller"
	//	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	//	"runtime"
	//	"time"
)

func main() {
	defer log.Println("Exiting")
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// go func() {
	// 	for {
	// 		log.Println(fmt.Sprintf("Number of Goroutines: %v", runtime.NumGoroutine()))
	// 		time.Sleep(time.Second * 5)
	// 	}
	// }()

	log.Println("Executor inializing...")
	router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc("/container/run", controller.RunContainer).Methods("POST")
	http.Handle("/", router)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
