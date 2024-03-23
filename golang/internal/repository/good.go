package repository

import (
	"fmt"
	"github.com/oooiik/test_09.03.2024/internal/database"
	"github.com/oooiik/test_09.03.2024/internal/filters"
	"github.com/oooiik/test_09.03.2024/internal/logger"
	"github.com/oooiik/test_09.03.2024/internal/model"
	"strings"
)

type Good interface {
	GetById(id uint32) (*model.Good, error)
	Create(model *model.Good) (*model.Good, error)
	Update(model *model.Good) (*model.Good, error)
	Delete(model *model.Good) (*model.Good, error)

	ListWithFilters(model *filters.Good) ([]*model.Good, error)
	CountWithFilters(model *filters.Good) (uint, error)
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
func (r *good) GetById(id uint32) (*model.Good, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE removed = false AND id = $1 LIMIT 1", r.table)
	logger.Debug(query, "; args:", id)
	rows, err := r.sql.DB().Query(query, id)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		findModel := &model.Good{}
		err := findModel.Scan(rows)
		if err != nil {
			return nil, err
		}
		return findModel, nil
	}
	return nil, nil
}

func (r *good) Create(m *model.Good) (*model.Good, error) {
	var cols []string
	var vals []any
	var plhs string

	i := 1
	for c, v := range m.ToCreate() {
		cols = append(cols, c)
		vals = append(vals, v)
		plhs += fmt.Sprintf("$%d, ", i)
		i++
	}
	plhs = plhs[:len(plhs)-2]

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) RETURNING id",
		r.table,
		strings.Join(cols, ", "),
		plhs,
	)

	logger.Debug(query, vals)

	var id uint32
	err := r.sql.DB().QueryRow(query, vals...).Scan(&id)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	logger.Debug(id)
	return r.GetById(id)
}

func (r *good) Update(m *model.Good) (*model.Good, error) {
	var sets string
	var vals []any

	i := 1
	for c, v := range m.ToUpdate() {
		sets += fmt.Sprintf("%s = $%d, ", c, i)
		vals = append(vals, v)
		i++
	}

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE id = %d",
		r.table,
		sets[:len(sets)-2],
		m.Id,
	)
	logger.Debug(query, vals)

	res, err := r.sql.DB().Exec(query, vals...)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return r.GetById(m.Id)
}

func (r *good) Delete(m *model.Good) (*model.Good, error) {
	query := fmt.Sprintf("UPDATE %s SET removed = true WHERE id = %d", r.table, m.Id)
	logger.Debug(query, m.Id)

	res, err := r.sql.DB().Exec(query)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	m.Fill(&model.Good{
		Removed: true,
	})

	return m, nil
}

func (r *good) queryFromFilters(f *filters.Good, isCount bool) (string, []any) {
	if f == nil {
		f = &filters.Good{}
	}
	filterList := f.ToFilters()

	var fils []string
	var vals []any

	i := 1
	for k, v := range filterList {
		fils = append(fils, fmt.Sprintf("%s = $%d", k, i))
		vals = append(vals, v)
		i++
	}

	fils = append(fils, "true") // protection for sql WHERE

	sel := "*"
	if isCount {
		sel = "COUNT(*)"
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", sel, r.table, strings.Join(fils, " AND "))

	logger.Debug(f)

	if f.Limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *f.Limit)
	}
	if f.Offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *f.Offset)
	}

	logger.Debug(query, "; args: ", vals)
	return query, vals
}

func (r good) ListWithFilters(f *filters.Good) ([]*model.Good, error) {
	query, vals := r.queryFromFilters(f, false)
	rows, err := r.sql.DB().Query(query, vals...)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var list []*model.Good
	for rows.Next() {
		l := &model.Good{}
		err := l.Scan(rows)
		if err != nil {
			return nil, err
		}
		list = append(list, l)
	}

	return list, nil
}

func (r good) CountWithFilters(f *filters.Good) (uint, error) {
	query, vals := r.queryFromFilters(f, true)
	var count uint
	err := r.sql.DB().QueryRow(query, vals...).Scan(&count)
	if err != nil {
		logger.Error(err)
		return 0, err
	}

	return count, nil
}
