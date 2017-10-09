package main

import (
	"./controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	defer log.Println("Exiting")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Executor inializing...")
	router := mux.NewRouter()
	router.StrictSlash(true)
	router.HandleFunc("/container/run", controller.RunContainer).Methods("POST")
	http.Handle("/", router)
	// go func() {
	// 	for {
	// 		time.Sleep(time.Second * 5)
	// 		log.Println(fmt.Sprintf("Number of Goroutines: %v", runtime.NumGoroutine()))
	// 	}
	// }()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}
