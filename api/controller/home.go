package controller

import (
	"net/http"

	"github.com/Joel-K-Muraguri/Go-Crud/api/responses"
)


func (s *Server) Home(w http.ResponseWriter , r *http.Request){
	responses.JSON(w, http.StatusOK, "Welcome to my GAMES CRUD API")
	
}