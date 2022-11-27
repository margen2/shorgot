package controllers

import (
	"github.com/margen2/shorgot/api/src/answers"
	"github.com/margen2/shorgot/api/src/auth"
	"github.com/margen2/shorgot/api/src/db"
	"github.com/margen2/shorgot/api/src/models"
	"github.com/margen2/shorgot/api/src/opsec"
	"github.com/margen2/shorgot/api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		answers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConnectDB()
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewUserRepositorie(db)
	savedUser, err := repositorie.SearchEmail(user.Email)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = opsec.VerifyPW(savedUser.Password, user.Password); err != nil {
		answers.JSON(w, http.StatusUnauthorized, nil)
		return
	}

	token, err := auth.CreateToken(savedUser.ID)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusAccepted, token)
}
