package repository

import (
	"errors"
	"github.com/oooiik/test_09.03.2024/internal/database"
	"github.com/oooiik/test_09.03.2024/internal/model"
)

var (
	ErrNotFound = errors.New("not found model")
)

type Interface interface {
	Init(table string, database database.Interface)

	tableName() string

	GetAll() ([]model.Interface, error)
	GetById(int32) (model.Interface, error)
	Create(model.Interface) (model.Interface, error)
	Update(model.Interface) (model.Interface, error)
	Delete(model.Interface) error
}

type repository struct {
	sql   database.Interface
	table string
	model model.Interface
}

func (r *repository) Init(table string, database database.Interface, model model.Interface) {
	r.table = table
	r.sql = database
}

func (r *repository) tableName() string {
	return r.table
}

func (r *repository) GetAll() ([]model.Interface, error) {
	// TODO
	return nil, nil
}
func (r *repository) GetById(id int32) (model.Interface, error) {
	// TODO
	return nil, nil
}
func (r *repository) Create(m model.Interface) (model.Interface, error) {
	// TODO
	return nil, nil
}
func (r *repository) Update(m model.Interface) (model.Interface, error) {
	// TODO
	return nil, nil
}
func (r *repository) Delete(m model.Interface) error {
	// TODO
	return nil
}
