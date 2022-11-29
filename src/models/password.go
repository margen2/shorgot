package models

// Passoword represents a password for the UpdatePassord function
type Password struct {
	NewPW string `json:"new"`
	OldPW string `json:"old"`
}

// Login represents the authentication response for a verified user
type Login struct {
	JWT string `json:"jwt"`
	User
}
