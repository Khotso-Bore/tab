package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Khotso-Bore/tab/pkg/config"
	"github.com/Khotso-Bore/tab/pkg/routes"
	"github.com/gorilla/mux"
)

func main(){

	config.Connect()
	
	r := mux.NewRouter()
	http.Handle("/",r)
	routes.RegisterUserRoutes(r)
	routes.RegisterTabRoutes(r)
	fmt.Println("starting on 8000...")
	log.Fatal(http.ListenAndServe(":8000",r))
}