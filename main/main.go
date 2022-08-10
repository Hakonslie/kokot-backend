package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restful/handlers"
	"restful/kokots"
	"restful/pages"
)

func handler(kokotController kokots.KokotController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("incoming %s request: %v\n", r.Method, r.URL)
		vars := mux.Vars(r)
		switch r.Method {
		case http.MethodGet:
			handlers.Get(vars["id"], kokotController, w)
			return
		case http.MethodPut:
			handlers.Put(kokotController, w, r)
			return
		case http.MethodPost:
			handlers.Post(kokotController, w, r)
			return
		case http.MethodDelete:
			handlers.Delete(vars["id"], kokotController, w)
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	}
}

func initialize(port string) *http.Server {
	log.Println("Initializing backend...")
	kokotController := kokots.NewKokotController()

	m := mux.NewRouter()
	m.Handle("/kokots/v1", handler(kokotController)).Methods(http.MethodGet, http.MethodPost)
	m.Handle("/kokots/v1/{id}", handler(kokotController)).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)

	srv := http.Server{
		Addr:    ":" + port,
		Handler: m,
	}
	return &srv
}
func initializeFrontend(port string) *http.Server {
	log.Println("Initializing frontend...")
	m := mux.NewRouter()
	m.Handle("/", pages.Router()).Methods(http.MethodGet, http.MethodPost)
	srv := http.Server{
		Addr:    ":" + port,
		Handler: m,
	}
	return &srv
}

func Run(port string, c chan int) {
	srv := initialize(port)
	frt := initializeFrontend("4040")
	go srv.ListenAndServe()
	fmt.Println("Backend running.")
	go frt.ListenAndServe()
	fmt.Println("Frontend running.")
	for {
		select {
		case <-c:
			frt.Shutdown(context.Background())
			srv.Shutdown(context.Background())
			return
		}
	}
}
func main() {
	c := make(chan int)
	Run("5050", c)

}
