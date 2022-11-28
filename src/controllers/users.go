package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/margen2/shorgot/src/answers"
	"github.com/margen2/shorgot/src/auth"
	"github.com/margen2/shorgot/src/db"
	"github.com/margen2/shorgot/src/models"
	"github.com/margen2/shorgot/src/repositories"
	"github.com/margen2/shorgot/src/security"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	if err = user.Prepare("signup"); err != nil {
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
	user.ID, err = repositorie.Create(user)
	if err != nil {
		answers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Password = ""
	answers.JSON(w, http.StatusCreated, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userid"], 10, 64)
	if err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	userIDToken, err := auth.ExtractUserID(r)
	if err != nil {
		answers.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDToken {
		answers.Error(w, http.StatusForbidden, errors.New("you can't update another user's account"))
		return
	}

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
	if err = user.Prepare("edit"); err != nil {
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
	if err = repositorie.UpdateEmail(userID, user); err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}
	answers.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userid"], 10, 64)
	if err != nil {
		answers.Error(w, http.StatusBadRequest, err)
	}

	userIDToken, err := auth.ExtractUserID(r)
	if err != nil {
		answers.Error(w, http.StatusUnauthorized, err)
		return
	}

	if userID != userIDToken {
		answers.Error(w, http.StatusForbidden, errors.New("you cannot delete someone's else account"))
		return
	}

	db, err := db.ConnectDB()
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()
	repositorie := repositories.NewUserRepositorie(db)
	if err := repositorie.DeleteUser(userID); err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := auth.ExtractUserID(r)
	if err != nil {
		answers.Error(w, http.StatusUnauthorized, err)
		return
	}
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["userid"], 10, 64)
	if err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}
	if userID != userIDToken {
		answers.Error(w, http.StatusForbidden, errors.New("you can't change someone's else password"))
		return
	}

	RequestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		answers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var password models.Password
	if err = json.Unmarshal(RequestBody, &password); err != nil {
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
	dbPW, err := repositorie.SearchPW(userID)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPW(dbPW, password.OldPW); err != nil {
		answers.Error(w, http.StatusUnauthorized, err)
		return
	}
	hashedPW, err := security.Hash(password.NewPW)
	if err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = repositorie.UpdatePW(userID, string(hashedPW)); err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}
