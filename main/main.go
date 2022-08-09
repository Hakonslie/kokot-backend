package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"restful/handlers"
	"restful/kokots"
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
	log.Println("initializing...")
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

func Run(port string, c chan int) {
	srv := initialize(port)
	go srv.ListenAndServe()
	for {
		select {
		case <-c:
			srv.Shutdown(context.Background())
			return
		}
	}
}
func Frontend(port string, c chan int) {

}
func main() {
	c := make(chan int)
	Run("5050", c)

}
