package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/margen2/shorgot/src/answers"
	"github.com/margen2/shorgot/src/auth"
	"github.com/margen2/shorgot/src/db"
	"github.com/margen2/shorgot/src/models"
	"github.com/margen2/shorgot/src/repositories"
	"github.com/margen2/shorgot/src/security"
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

	if err = security.VerifyPW(savedUser.Password, user.Password); err != nil {
		answers.JSON(w, http.StatusUnauthorized, nil)
		return
	}
	var login models.Login
	login.ID = savedUser.ID
	login.JWT, err = auth.CreateToken(savedUser.ID)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusAccepted, login)
}
