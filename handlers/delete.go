package handlers

import (
	"net/http"
	"restful/kokots"
)

func Delete(id string, k kokots.KokotController, w http.ResponseWriter) {
	k.Delete(id)
	if exists := k.GetOne(id); exists != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		return
	}
}
