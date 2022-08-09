package handlers

import (
	"fmt"
	"net/http"
	"restful/kokots"
)

func Delete(id string, k kokots.KokotController, w http.ResponseWriter) {
	k.Delete(id)
	fmt.Printf("Deleting id: %s \n", id)
	if exists := k.GetOne(id); exists != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		return
	}
}
