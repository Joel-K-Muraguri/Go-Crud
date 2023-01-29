package controller

import (
	"net/http"

	utils "github.com/Joel-K-Muraguri/Go-Crud/api/responses"
)


func (s *Server) Home(w http.ResponseWriter , r *http.Request){
	utils.JSON(w, http.StatusOK, "Welcome to my GAMES CRUD API")
	
	
}