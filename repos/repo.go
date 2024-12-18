package repositories

import (
	"database/sql"
	"github.com/dbsimmons64/posty/models"
)

type PostRepository interface {
	Insert(title, content string) error
	All() ([]models.Post, error)
}

type PostRepositoryDB struct {
	DB *sql.DB
}

func (m PostRepositoryDB) Insert(title, content string) error {
	stmt := `
    INSERT INTO posts (title, content, createdAt)
    VALUES (?, ?, datetime('now'))
  `
	_, err := m.DB.Exec(stmt, title, content)

	return err

}

func (m PostRepositoryDB) All() ([]models.Post, error) {
	stmt := `SELECT id, title, content, createdAt FROM posts ORDER BY id DESC`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	posts := []models.Post{}

	for rows.Next() {
		p := models.Post{}
		err := rows.Scan(&p.Id, &p.Title, &p.Content, &p.CreatedAt)
		if err != nil {
			return nil, err
		}

		posts = append(posts, p)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return posts, nil
}
