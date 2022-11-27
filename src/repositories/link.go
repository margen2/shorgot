package repositories

import (
	"database/sql"
	"time"

	"github.com/margen2/shorgot/api/src/models"
)

type Links struct {
	db *sql.DB
}

func NewLinksRepositorie(db *sql.DB) *Links {
	return &Links{db}
}

func (repositorie Links) Create(link models.Link) error {
	statment, err := repositorie.db.Prepare(
		`INSERT INTO links (target_url, shortened_url, created_on, user_id)
		VALUES (?, ?, ?, ?)`,
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

func (repositorie Links) SearchShortenedURL(linkID string) (models.Link, error) {
	line, err := repositorie.db.Query(`
			SELECT * FROM links WHERE shortened_url = ?`, linkID,
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

func (repositorie Links) UpdateLink(linkID uint64, link models.Link) error {
	statement, err := repositorie.db.Prepare("UPDATE links SET target_url = ?, shortened_url = ? WHERE link_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(link.Target, link.Shortened, linkID); err != nil {
		return err
	}

	return nil
}

func (repositorie Links) DeleteLink(postID uint64) error {
	statement, err := repositorie.db.Prepare("DELETE FROM links WHERE link_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(postID); err != nil {
		return err
	}

	return nil
}

func (repositorie Links) SearchLinkByUserID(userID uint64) (models.Link, error) {
	line, err := repositorie.db.Query(`
	SELECT * FROM links WHERE user_id = ?`,
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

func (repositorie Links) SearchLinkByID(linkID uint64) (models.Link, error) {
	line, err := repositorie.db.Query(`
	SELECT * FROM links WHERE link_id = ?`,
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

func (repositorie Links) AddClick(linkID uint64) error {
	statement, err := repositorie.db.Prepare("UPDATE links SET clicks = clicks +1 WHERE link_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(linkID); err != nil {
		return err
	}

	return nil
}
