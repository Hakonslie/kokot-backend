package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"restful/kokots"
)

// Put Updates a current kokot or creates new if no id present. Returns Kokot
func Put(k kokots.KokotController, w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var ko kokots.Kokot
	err = json.Unmarshal(b, &ko)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if ok := k.Update(ko); !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		return
	}
}
