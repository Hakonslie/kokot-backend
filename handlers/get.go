package handlers

import (
	"encoding/json"
	"net/http"
	"restful/kokots"
)

func Get(id string, kokotController kokots.KokotController, w http.ResponseWriter) {
	// if ID is present - get one kokot
	if id != "" {
		kokot := kokotController.GetOne(id)
		if kokot == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			// Put Kokot in json
			m, err := json.Marshal(kokot)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
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
			return
		}
		w.WriteHeader(200)
		w.Write(k)
		return
	}
}
