package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"restful/kokots"
)

func Post(k kokots.KokotController, w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var name string
	err = json.Unmarshal(b, &name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(string(b))
		fmt.Println(err)
		return
	}
	// doesn't exist. Create new
	added := k.Add(name)
	if added == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		if j, err := json.Marshal(added); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusCreated)
			w.Write(j)
			return
		}
	}
}
