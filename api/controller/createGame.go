package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Joel-K-Muraguri/Go-Crud/api/models"
	"github.com/Joel-K-Muraguri/Go-Crud/api/responses"
	"github.com/Joel-K-Muraguri/Go-Crud/api/utils"
)

func (s *Server) CreateGame(w http.ResponseWriter , r *http.Request){
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	game := models.Game{}
	err = json.Unmarshal(body, &game)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	game.Prepare()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := game.SaveGame(s.DB)

	if err != nil {

		formattedError := utils.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	responses.JSON(w, http.StatusCreated, userCreated)
	// responses.JSON(w, http.StatusOK, "Create Game")
	
}