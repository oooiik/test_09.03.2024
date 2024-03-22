package repository

import (
	"fmt"
	"github.com/oooiik/test_09.03.2024/internal/database"
	"github.com/oooiik/test_09.03.2024/internal/logger"
	"github.com/oooiik/test_09.03.2024/internal/model"
	"strings"
)

type Good interface {
	ListWithPagination(limit, offset uint32) ([]*model.Good, error)
	ListWithFilters(model *model.Good) ([]*model.Good, error)
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
func (r good) ListWithPagination(limit, offset uint32) ([]*model.Good, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE removed = false LIMIT %d OFFSET %d", r.table, limit, offset)
	logger.Debug(query)
	rows, err := r.sql.DB().Query(query)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	list := make([]*model.Good, 0)

	for rows.Next() {
		i := &model.Good{}
		err := i.Scan(rows)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		list = append(list, i)
	}

	return list, nil
}

func (r good) GetById(id uint32) (*model.Good, error) {
	//TODO implement me
	panic("implement me")
}

func (r good) Create(model *model.Good) (*model.Good, error) {
	var cols []string
	var vals []any

	for c, v := range model.ToDbList() {
		cols = append(cols, c)
		vals = append(vals, v)
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		r.table,
		strings.Join(cols, ", "),
		strings.Join(make([]string, len(cols)), "?, ")+"?",
	)

	res, err := r.sql.DB().Exec(query, vals...)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	newModel, err := r.GetById(uint32(id))
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return newModel, nil
}

func (r good) Update(model *model.Good) (*model.Good, error) {
	//TODO implement me
	panic("implement me")
}

func (r good) Delete(model *model.Good) error {
	//TODO implement me
	panic("implement me")
}

func (r good) ListWithFilters(model *model.Good) ([]*model.Good, error) {
	//TODO implement me
	panic("implement me")
}
