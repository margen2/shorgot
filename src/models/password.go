package models

type Password struct {
	NewPW string `json:"new"`
	OldPW string `json:"old"`
}

type Login struct {
	JWT string `json:"jwt"`
	User
}
