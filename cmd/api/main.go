package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	//Mux is router
	mux := http.NewServeMux()
	//Add router it is mainly used for a 200 status code
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("all ok"))
	})
	srv := http.Server{
		Addr:         ":1999",
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Server faild %V", err)
	}

}
