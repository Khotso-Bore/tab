package routes

import (
	"github.com/Khotso-Bore/tab/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterTabRoutes = func(router *mux.Router){
	router.HandleFunc("/tab",controllers.CreateTab).Methods("POST")
	router.HandleFunc("/tab/{userid}/{contactid}",controllers.GetTabForUserContact).Methods("GET")
	router.HandleFunc("/tab/paytab",controllers.PayTab).Methods("POST")
}