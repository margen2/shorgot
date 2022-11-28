package repositories

import (
	"database/sql"

	"github.com/margen2/shorgot/src/models"
)

type Users struct {
	db *sql.DB
}

func NewUserRepositorie(db *sql.DB) *Users {
	return &Users{db}
}

func (repositorie Users) Create(user models.User) (uint64, error) {
	statment, err := repositorie.db.Prepare(
		"INSERT INTO users(email, password) VALUES($1, $2) RETURNING user_id",
	)
	if err != nil {
		return 0, err
	}
	defer statment.Close()

	id := 0
	err = statment.QueryRow(user.Email, user.Password).Scan(&id)
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

func (repositorie Users) UpdateEmail(ID uint64, user models.User) error {
	statment, err := repositorie.db.Prepare(
		"UPDATE users SET email = $1 WHERE user_id = $2",
	)
	if err != nil {
		return err
	}
	defer statment.Close()
	if _, err := statment.Exec(user.Email, ID); err != nil {
		return err
	}
	return nil
}

func (repositorie Users) DeleteUser(ID uint64) error {
	statment, err := repositorie.db.Prepare(
		"DELETE FROM users WHERE user_id = $1",
	)
	if err != nil {
		return err
	}
	defer statment.Close()
	if _, err := statment.Exec(ID); err != nil {
		return err
	}
	return nil
}

func (repositorie Users) SearchEmail(email string) (models.User, error) {
	line, err := repositorie.db.Query("SELECT user_id, password FROM users WHERE email = $1", email)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (repositorie Users) SearchPW(userID uint64) (string, error) {
	line, err := repositorie.db.Query("SELECT password FROM users WHERE user_id = $1", userID)
	if err != nil {
		return "", err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(&user.Password); err != nil {
			return "", err
		}
	}

	return user.Password, nil
}

func (repositorie Users) UpdatePW(userID uint64, pw string) error {
	statement, err := repositorie.db.Prepare("UPDATE users SET password = $1 WHERE user_id = $2")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(pw, userID); err != nil {
		return err
	}
	return nil
}
