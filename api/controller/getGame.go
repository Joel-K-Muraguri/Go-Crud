package controller

import (
	"net/http"
	"strconv"

	"github.com/Joel-K-Muraguri/Go-Crud/api/models"
	"github.com/Joel-K-Muraguri/Go-Crud/api/responses"
	"github.com/gorilla/mux"
)

func (s *Server) GetGame(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	games := models.Game{}
	userGotten, err := games.FindGameByID(s.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, userGotten)
	// responses.JSON(w, http.StatusOK, "Get a game")
	
}