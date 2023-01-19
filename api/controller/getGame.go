package controller

import (
	"net/http"

	"github.com/Joel-K-Muraguri/Go-Crud/api/responses"
)

func (s *Server) GetGame(w http.ResponseWriter , r *http.Request){
	responses.JSON(w, http.StatusOK, "Get a game")
	
}