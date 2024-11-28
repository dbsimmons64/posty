package services

import (
	"github.com/dbsimmons64/posty/models"
	"github.com/dbsimmons64/posty/repos"
)

type PostService interface {
	CreatePost(title, content string) error
	ListPosts() ([]models.Post, error)
}

type PostServiceImpl struct {
	Repo repositories.PostRepository
}

func (ps PostServiceImpl) ListPosts() ([]models.Post, error) {
	return ps.Repo.All()
}

func (ps PostServiceImpl) CreatePost(title, content string) error {
	return ps.Repo.Insert(title, content)
}
