package controllers

import (
	"github.com/inventory-app/backend/golang/common/log"
	"net/http"
)

type Controller struct {
}

func (c *Controller) HealthCheck(w http.ResponseWriter, r *http.Request) {
	applog.Log.Info("App is up and running!")
	response := []byte("{\"Message\":\"success\"}")

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	return
}
