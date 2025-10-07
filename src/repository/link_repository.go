package repository

import (
	"database/sql"
	"linkbio-go/src/model"
)

// LinkRepository is the layer for data interaction related to links
type LinkRepository struct {
	DB *sql.DB
}

// NewLinkRepository creates a new instance of LinkRepository
func NewLinkRepository(db *sql.DB) *LinkRepository {
	return &LinkRepository{DB: db}
}

// GetAll retrieves all links from the database
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

// Create stores a new link in the database
func (r *LinkRepository) Create(link model.Link) (model.Link, error) {
	query := "INSERT INTO links (title, url) VALUES ($1, $2) RETURNING id, created_at"
	err := r.DB.QueryRow(query, link.Title, link.URL).Scan(&link.ID, &link.CreatedAt)
	if err != nil {
		return model.Link{}, err
	}
	return link, nil
}

// GetByID retrieves a single link by ID
func (r *LinkRepository) GetByID(id int) (model.Link, error) {
	var link model.Link
	query := "SELECT id, title, url, created_at FROM links WHERE id = $1"
	err := r.DB.QueryRow(query, id).Scan(&link.ID, &link.Title, &link.URL, &link.CreatedAt)
	if err == sql.ErrNoRows {
		return model.Link{}, err
	}
	if err != nil {
		return model.Link{}, err
	}
	return link, nil
}

// Update modifies an existing link in the database
func (r *LinkRepository) Update(id int, link model.Link) (model.Link, error) {
	query := "UPDATE links SET title = $1, url = $2 WHERE id = $3 RETURNING id, title, url, created_at"
	var updatedLink model.Link
	err := r.DB.QueryRow(query, link.Title, link.URL, id).Scan(&updatedLink.ID, &updatedLink.Title, &updatedLink.URL, &updatedLink.CreatedAt)
	if err != nil {
		return model.Link{}, err
	}
	return updatedLink, nil
}

// Delete removes a link from the database by ID
func (r *LinkRepository) Delete(id int) error {
	query := "DELETE FROM links WHERE id = $1"
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}