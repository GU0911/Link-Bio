package repository

import "linkbio-go/src/model"

type ILinkRepository interface {
	GetAll() ([]model.Link, error)
	GetByID(id int) (model.Link, error)
	Create(link model.Link) (model.Link, error)
	Update(id int, link model.Link) (model.Link, error)
	Delete(id int) error
}