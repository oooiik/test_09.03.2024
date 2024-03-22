package model

import (
	"database/sql"
	"github.com/oooiik/test_09.03.2024/internal/logger"
	"reflect"
	"time"
)

type Good struct {
	Id          uint32     `db:"id"`
	ProjectId   uint32     `db:"project_id"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	Priority    uint32     `db:"priority"`
	Removed     bool       `db:"removed"`
	CreatedAt   *time.Time `db:"created_at"`
}

func (m *Good) Fill(f *Good) {
	if f.Id != 0 {
		m.Id = f.Id
	}
	if f.ProjectId != 0 {
		m.ProjectId = f.ProjectId
	}
	if f.Name != "" {
		m.Name = f.Name
	}
	if f.Description != "" {
		m.Description = f.Description
	}
	if f.Priority != 0 {
		m.Priority = f.Priority
	}
	if f.Removed != false {
		m.Removed = f.Removed
	}
	if f.CreatedAt != nil {
		m.CreatedAt = f.CreatedAt
	}
}

func (m *Good) Scan(rows *sql.Rows) error {
	cols, err := rows.Columns()
	if err != nil {
		logger.Error(err)
		return err
	}

	scanArgs := make([]interface{}, len(cols))
	for i := range cols {
		scanArgs[i] = new(interface{})
	}

	if err := rows.Scan(scanArgs...); err != nil {
		logger.Error(err)
		return err
	}

	refTypeOf := reflect.TypeOf(*m)
	refValueOf := reflect.ValueOf(m).Elem()

	for i := 0; i < refTypeOf.NumField(); i++ {
		field := refTypeOf.Field(i)
		tag := field.Tag.Get("db")
		for iCol, col := range cols {
			if tag == col {
				fieldValue := refValueOf.Field(i)
				if fieldValue.IsValid() && fieldValue.CanSet() {
					// Check is pointer
					if fieldValue.Kind() == reflect.Ptr {
						if scanArgs[iCol] != nil {
							// Create new pointed type
							val := reflect.New(fieldValue.Type().Elem())
							val.Elem().Set(reflect.ValueOf(*scanArgs[iCol].(*interface{})).Convert(fieldValue.Type().Elem()))
							fieldValue.Set(val)
						}
					} else {
						fieldValue.Set(reflect.ValueOf(*scanArgs[iCol].(*interface{})).Convert(fieldValue.Type()))
					}
				}
				break
			}
		}
	}

	return nil
}

func (m *Good) ToDbList() map[string]any {
	result := make(map[string]interface{})

	refTypeOf := reflect.TypeOf(*m)
	refValueOf := reflect.ValueOf(m)

	for i := 0; i < refTypeOf.NumField(); i++ {
		field := refTypeOf.Field(i)
		value := refValueOf.Field(i).Interface()

		key := field.Tag.Get("db")
		if key == "" {
			break
		}

		result[key] = value
	}

	return result
}
