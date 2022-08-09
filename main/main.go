package main

import (
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

func initalize() *mux.Router {
	log.Println("initalizing...")
	kokotController := kokots.NewKokotController()
	m := mux.NewRouter()
	m.Handle("/kokots/v1", handler(kokotController)).Methods(http.MethodGet, http.MethodPost)
	m.Handle("/kokots/v1/{id}", handler(kokotController)).Methods(http.MethodGet, http.MethodPut, http.MethodDelete)
	return m
}

func run(router *mux.Router, port string) {
	log.Printf("Running on %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
func App(port string) {
	run(initalize(), port)
}
func main() {
	App("5050")
}
