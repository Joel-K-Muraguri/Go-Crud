package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Joel-K-Muraguri/Go-Crud/api/models"
	"github.com/Joel-K-Muraguri/Go-Crud/api/responses"
	"github.com/gorilla/mux"
)

func (s *Server) UpdateGame(w http.ResponseWriter , r *http.Request){
	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	game := models.Game{}
	err = json.Unmarshal(body, &game)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	game.Prepare()
	err = game.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := game.UpdateAGame(s.DB, uint32(uid))
	if err != nil {
		formattedError := responses.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedUser)
	// responses.JSON(w, http.StatusOK, "Update Game")
	
}