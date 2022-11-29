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

	"github.com/gorilla/mux"
)

// CreateLink creates a link for the given user
func CreateLink(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		answers.Error(w, http.StatusUnauthorized, err)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		answers.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var link models.Link
	if err = json.Unmarshal(requestBody, &link); err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	link.AuthorID = userID

	db, err := db.ConnectDB()
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewLinksRepositorie(db)
	err = repositorie.Create(link)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusCreated, link)
}

// SearchLink returns the complete link information from the given shortened url
func SearchLink(w http.ResponseWriter, r *http.Request) {
	shortened := mux.Vars(r)["link"]

	db, err := db.ConnectDB()
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewLinksRepositorie(db)
	link, err := repositorie.SearchShortenedURL(shortened)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	err = repositorie.AddClick(link.ID)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}
	answers.JSON(w, http.StatusOK, link)
}

// SearchLinks returns all the active links for the given user ID
func SearchLinks(w http.ResponseWriter, r *http.Request) {
	userIDToken, err := auth.ExtractUserID(r)
	if err != nil {
		answers.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := db.ConnectDB()
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewLinksRepositorie(db)
	links, err := repositorie.SearchLinkByUserID(userIDToken)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, links)
}

// UpdateLink updates the given link
func UpdateLink(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		answers.Error(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	linkID, err := strconv.ParseUint(parameters["linkID"], 10, 64)
	if err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConnectDB()
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewLinksRepositorie(db)
	linkOnDB, err := repositorie.SearchLinkByID(linkID)
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	if userID != linkOnDB.AuthorID {
		answers.Error(w, http.StatusForbidden, errors.New("you can't change someone'else link"))
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	var link models.Link
	if err = json.Unmarshal(requestBody, &link); err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = repositorie.UpdateLink(linkID, link); err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}

// DeleteLink deletes the given link from the database
func DeleteLink(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		answers.Error(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	linkID, err := strconv.ParseUint(parameters["linkID"], 10, 64)
	if err != nil {
		answers.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.ConnectDB()
	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repositorie := repositories.NewLinksRepositorie(db)
	linkOnDB, err := repositorie.SearchLinkByID(linkID)

	if err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	if userID != linkOnDB.AuthorID {
		answers.Error(w, http.StatusForbidden, errors.New("you can't change someone else's link"))
		return
	}

	if err = repositorie.DeleteLink(linkID); err != nil {
		answers.Error(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusNoContent, nil)
}
