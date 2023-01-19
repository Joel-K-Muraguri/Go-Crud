package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


type Server struct{
	Router *mux.Router
}


func (server *Server) Initialize(){

	server.Router = mux.NewRouter()

	server.initializeRoutes()

}


func (server *Server) Run(port string){
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(port, server.Router))
}