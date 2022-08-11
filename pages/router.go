package pages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"restful/kokots"
	"strings"
)

func Router() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			if r.URL.String() == "/" {
				Front(w)
				return
			}
			w.WriteHeader(http.StatusNotFound)
			return
		case http.MethodPost:
			if err := r.ParseForm(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			switch r.FormValue("action") {
			case "add":
				name, err := json.Marshal(r.FormValue("name"))
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				_, err = http.Post("http://localhost:5050/kokots/v1", "application/json", bytes.NewReader(name))
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			case "delete":
				id, err := json.Marshal(r.FormValue("id"))
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				cleaned := strings.Replace(string(id), "\"", "", -1)
				req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:5050/kokots/v1/%s", cleaned), nil)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				_, err = http.DefaultClient.Do(req)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			case "put":
				bod, err := json.Marshal(kokots.Kokot{
					ID:   r.FormValue("id"),
					Name: r.FormValue("name"),
				})
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:5050/kokots/v1"), bytes.NewReader(bod))
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				_, err = http.DefaultClient.Do(req)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}

			}
			http.Redirect(w, r, "http://localhost:4040", 303)
			return
		}
		return
	}
}
