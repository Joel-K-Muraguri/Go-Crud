package controller

import (
	"net/http"

	"github.com/Joel-K-Muraguri/Go-Crud/api/responses"
)

func (s *Server) DeleteGame(w http.ResponseWriter , r *http.Request){
	responses.JSON(w, http.StatusOK, "Delete all Games")
	
}