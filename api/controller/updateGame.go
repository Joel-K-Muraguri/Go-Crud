package controller

import (
	"net/http"

	"github.com/Joel-K-Muraguri/Go-Crud/api/responses"
)

func (s *Server) UpdateGame(w http.ResponseWriter , r *http.Request){
	responses.JSON(w, http.StatusOK, "Update Game")
	
}