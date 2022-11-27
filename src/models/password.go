package models

type Password struct {
	NewPW string `json:"new"`
	OldPW string `json:"old"`
}
