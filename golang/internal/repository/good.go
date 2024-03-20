package repository

import (
	"github.com/oooiik/test_09.03.2024/internal/database"
	"github.com/oooiik/test_09.03.2024/internal/model"
)

type Good interface {
	ListWithPagination(limit, offset uint32) ([]*model.Good, error)
	GetById(id uint32) (*model.Good, error)
	Create(model *model.Good) (*model.Good, error)
	Update(model *model.Good) (*model.Good, error)
	Delete(model *model.Good) error
}

type good struct {
	sql   database.Interface
	table string
}

func NewGood(db database.Interface) Good {
	return &good{
		sql:   db,
		table: "goods",
	}
}
func (g good) ListWithPagination(limit, offset uint32) ([]*model.Good, error) {
	//TODO implement me
	panic("implement me")
}

func (g good) GetById(id uint32) (*model.Good, error) {
	//TODO implement me
	panic("implement me")
}

func (g good) Create(model *model.Good) (*model.Good, error) {
	//TODO implement me
	panic("implement me")
}

func (g good) Update(model *model.Good) (*model.Good, error) {
	//TODO implement me
	panic("implement me")
}

func (g good) Delete(model *model.Good) error {
	//TODO implement me
	panic("implement me")
}
