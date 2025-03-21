package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main(){

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/hello", basicHandler)

	server := &http.Server{
		Addr: ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil{ 
		fmt.Println(err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet{
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}
	w.Write([]byte("Hello World"))
}