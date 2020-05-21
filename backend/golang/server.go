package main

import (
	"github.com/inventory-app/backend/golang/common/conf"
	"github.com/inventory-app/backend/golang/common/context"
	"github.com/inventory-app/backend/golang/common/log"
	"github.com/inventory-app/backend/golang/router"
	"net/http"
	"time"
)

func main() {
	applog.InitializeLogging()
	time.Sleep(1 * time.Second)

	applog.Log.Info("Server Started .........")

	// if changing port change in homepage to serve the ui also(optional)
	router := router.NewRouter() // create routes

	router.Methods("GET", "POST", "DELETE", "UPDATE")

	context.ApplicationContext()

	serv := &http.Server{
		Handler: router,
		Addr:    conf.Cfg.Server.ServerAddress,

		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	applog.Log.Fatal(serv.ListenAndServe())

	defer applog.Log.Info("Server Closed !!! Please restart server...... ")
}
