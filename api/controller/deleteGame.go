package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Joel-K-Muraguri/Go-Crud/api/models"
	"github.com/Joel-K-Muraguri/Go-Crud/api/responses"
	"github.com/gorilla/mux"
)

func (s *Server) DeleteGame(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)

	game := models.Game{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = game.DeleteAGame(s.DB, uint32(uid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	responses.JSON(w, http.StatusNoContent, "")
	// responses.JSON(w, http.StatusOK, "Delete all Games")
	
}