package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/MYK12397/handlers"
	"github.com/gorilla/mux"
)

func main() {

	sm := mux.NewRouter()

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/v1/users", handlers.GetUsers)
	getR.HandleFunc("/v1/users/{id}", handlers.GetID)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/v1/users", handlers.CreateUser)

	putR := sm.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/v1/users/{id}", handlers.UpdateUser)

	deleteR := sm.Methods(http.MethodDelete).Subrouter()
	deleteR.HandleFunc("/v1/users/{id}", handlers.DeleteUser)

	s := &http.Server{
		Addr:         ":8000",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			panic(err)
		}

	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <-sigChan

	fmt.Println("Received Terminal signal. Shutting down. ", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)

}
