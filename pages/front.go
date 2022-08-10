package pages

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"restful/kokots"
)

func Front(w http.ResponseWriter) {
	r, err := http.Get("http://localhost:5050/kokots/v1")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var AllKokots []kokots.Kokot
	err = json.Unmarshal(b, &AllKokots)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data := struct {
		AllKokots []kokots.Kokot
	}{AllKokots: AllKokots}
	t, err := template.ParseFiles("./pages/front.html")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = t.Execute(w, data)
	return
}
