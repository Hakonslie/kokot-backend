package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"restful/kokots"
)

func Get(id string, kokotController kokots.KokotController, w http.ResponseWriter) {
	// if ID is present - get one kokot
	if id != "" {
		kokot := kokotController.GetOne(id)
		if kokot == nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Kokot with ID: %s not found", id)
			return
		} else {
			// Put Kokot in json
			m, err := json.Marshal(kokot)
			if err != nil {
				// failed
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
				return
			}
			w.WriteHeader(200)
			w.Write(m)
		}
	} else {
		// get All kokots
		k, err := json.Marshal(kokotController.GetAll())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			return
		}
		w.WriteHeader(200)
		w.Write(k)
		return
	}
}
