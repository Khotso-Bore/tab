package routes

import (
	"github.com/Khotso-Bore/tab/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterUserRoutes = func(router *mux.Router){
	router.HandleFunc("/user",controllers.CreateUser).Methods("POST")
	router.HandleFunc("/user",controllers.GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}",controllers.GetUser).Methods("GET")
	router.HandleFunc("/user/{id}",controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}",controllers.DeleteUser).Methods("DELETE")
	router.HandleFunc("/user/addContact",controllers.AddContact).Methods("POST")
	router.HandleFunc("/user/getContatcs/{id}",controllers.GetContacts).Methods("GET")
	router.HandleFunc("/user/depositFunds",controllers.DepositFunds).Methods("POST")

}

