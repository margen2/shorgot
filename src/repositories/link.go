package repositories

import (
	"database/sql"
	"time"

	"github.com/margen2/shorgot/src/models"
)

type Links struct {
	db *sql.DB
}

func NewLinksRepositorie(db *sql.DB) *Links {
	return &Links{db}
}

// Create creates a new user on the database
func (repositorie Links) Create(link models.Link) error {
	statment, err := repositorie.db.Prepare(
		`INSERT INTO links (target_url, shortened_url, created_on, user_id)
		VALUES ($1, $2, $3, $4)`,
	)
	if err != nil {
		return err
	}
	defer statment.Close()

	_, err = statment.Exec(link.Target, link.Shortened, time.Now(), link.AuthorID)
	if err != nil {
		return err
	}
	return nil
}

// SearchShortenedURL searches for the full entry based on the given shortened url
func (repositorie Links) SearchShortenedURL(shortened string) (models.Link, error) {
	line, err := repositorie.db.Query(`
			SELECT * FROM links WHERE shortened_url = $1`, shortened,
	)
	if err != nil {
		return models.Link{}, err
	}
	defer line.Close()

	var link models.Link

	if line.Next() {
		if err = line.Scan(
			&link.ID,
			&link.Target,
			&link.Shortened,
			&link.CreatedOn,
			&link.Clicks,
			&link.AuthorID,
		); err != nil {
			return models.Link{}, err
		}
	}
	return link, nil
}

// UpdateLink updates the link based on the given linkID
func (repositorie Links) UpdateLink(linkID uint64, link models.Link) error {
	statement, err := repositorie.db.Prepare("UPDATE links SET target_url = $1, shortened_url = $2 WHERE link_id = $3")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(link.Target, link.Shortened, linkID); err != nil {
		return err
	}

	return nil
}

// DeleteLink deletes the link entry based on the given linkID
func (repositorie Links) DeleteLink(linkID uint64) error {
	statement, err := repositorie.db.Prepare("DELETE FROM links WHERE link_id = $1")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(linkID); err != nil {
		return err
	}

	return nil
}

// SearchLinkByUserID search all the links associated with the given userID
func (repositorie Links) SearchLinkByUserID(userID uint64) (models.Link, error) {
	line, err := repositorie.db.Query(`
	SELECT * FROM links WHERE user_id = $1`,
		userID)

	if err != nil {
		return models.Link{}, err
	}
	defer line.Close()

	var link models.Link

	if line.Next() {
		if err = line.Scan(
			&link.ID,
			&link.Target,
			&link.Shortened,
			&link.CreatedOn,
			&link.Clicks,
			&link.AuthorID,
		); err != nil {
			return models.Link{}, err
		}
	}
	return link, nil
}

// SearchLinkByID searches for the full entry link based on the given linkID
func (repositorie Links) SearchLinkByID(linkID uint64) (models.Link, error) {
	line, err := repositorie.db.Query(`
	SELECT * FROM links WHERE link_id = $1`,
		linkID)

	if err != nil {
		return models.Link{}, err
	}
	defer line.Close()

	var link models.Link

	if line.Next() {
		if err = line.Scan(
			&link.ID,
			&link.Target,
			&link.Shortened,
			&link.CreatedOn,
			&link.Clicks,
			&link.AuthorID,
		); err != nil {
			return models.Link{}, err
		}
	}
	return link, nil
}

// AddClick adds one click to the given linkID
func (repositorie Links) AddClick(linkID uint64) error {
	statement, err := repositorie.db.Prepare("UPDATE links SET clicks = clicks +1 WHERE link_id = $1")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(linkID); err != nil {
		return err
	}

	return nil
}
