package repository

import (
	"database/sql"
	"linkbio-go/src/model"
)

type LinkRepository struct {
	DB *sql.DB
}

func NewLinkRepository(db *sql.DB) *LinkRepository {
	return &LinkRepository{DB: db}
}

func (r *LinkRepository) GetAll() ([]model.Link, error) {
	query := "SELECT id, title, url, created_at FROM links ORDER BY created_at DESC"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []model.Link
	for rows.Next() {
		var link model.Link
		if err := rows.Scan(&link.ID, &link.Title, &link.URL, &link.CreatedAt); err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	return links, nil
}

func (r *LinkRepository) Create(link model.Link) (model.Link, error) {
	query := "INSERT INTO links (title, url) VALUES ($1, $2) RETURNING id, created_at"
	err := r.DB.QueryRow(query, link.Title, link.URL).Scan(&link.ID, &link.CreatedAt)
	if err != nil {
		return model.Link{}, err
	}
	return link, nil
}