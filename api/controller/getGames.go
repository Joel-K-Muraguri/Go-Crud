package controller

import (
	"net/http"

	"github.com/Joel-K-Muraguri/Go-Crud/api/models"
	"github.com/Joel-K-Muraguri/Go-Crud/api/responses"
)

func (s *Server) GetGames(w http.ResponseWriter , r *http.Request){
	game := models.Game{}

	games, err := game.FindAllGames(s.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, games)
	// responses.JSON(w, http.StatusOK, "Get All Games")
	
}