package model

import (
	"database/sql"
	"github.com/oooiik/test_09.03.2024/internal/logger"
	"reflect"
	"strings"
	"time"
)

// db tag - column name in table
// allow: r - read, c - create, u - update
type Good struct {
	Id          uint32     `db:"id" allow:"r"`
	ProjectId   uint32     `db:"project_id" allow:"c,r"`
	Name        string     `db:"name" allow:"c,u,r"`
	Description string     `db:"description" allow:"u"` // INFO: why description not readonly
	Priority    uint32     `db:"priority" allow:"u,r"`
	Removed     bool       `db:"removed" allow:"u,r"`
	CreatedAt   *time.Time `db:"created_at" allow:"r"`
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

	toReads := m.ToRead()
	refTypeOf := reflect.TypeOf(*m)
	refValueOf := reflect.ValueOf(m).Elem()

	for i := 0; i < refTypeOf.NumField(); i++ {
		field := refTypeOf.Field(i)
		tag := field.Tag.Get("db")
		isRead := false
		for _, toRead := range toReads {
			if tag == toRead {
				isRead = true
			}
		}

		if !isRead {
			continue
		}

		for iCol, col := range cols {
			if tag == col {
				fieldValue := refValueOf.Field(i)
				refValueRes := reflect.ValueOf(*scanArgs[iCol].(*interface{}))
				if fieldValue.IsValid() && fieldValue.CanSet() && refValueRes.IsValid() {
					// Check is pointer
					if fieldValue.Kind() == reflect.Ptr {
						if scanArgs[iCol] != nil {
							// Create new pointed type
							val := reflect.New(fieldValue.Type().Elem())
							val.Elem().Set(refValueRes.Convert(fieldValue.Type().Elem()))
							fieldValue.Set(val)
						}
					} else {
						fieldValue.Set(refValueRes.Convert(fieldValue.Type()))
					}
				}
				break
			}
		}
	}

	return nil
}

func (m *Good) toDbList() map[string]any {
	result := make(map[string]interface{})

	refTypeOf := reflect.TypeOf(*m)
	refValueOf := reflect.ValueOf(m).Elem()

	for i := 0; i < refTypeOf.NumField(); i++ {
		field := refTypeOf.Field(i)

		key := field.Tag.Get("db")
		if key == "" {
			continue
		}

		fieldValue := refValueOf.Field(i).Interface()
		result[key] = fieldValue
	}
	return result
}

func (m *Good) toOrm(tagKey, key string) map[string]interface{} {
	values := m.toDbList()

	refTypeOf := reflect.TypeOf(*m)
	for i := 0; i < refTypeOf.NumField(); i++ {
		field := refTypeOf.Field(i)
		tag := field.Tag.Get(tagKey)
		keys := strings.Split(tag, ",")

		has := false
		for _, k := range keys {
			if k == key {
				has = true
				break
			}
		}

		if !has {
			delete(values, field.Tag.Get("db"))
		}
	}

	return values
}

func (m *Good) ToCreate() map[string]interface{} {
	return m.toOrm("allow", "c")
}

func (m *Good) ToUpdate() map[string]interface{} {
	return m.toOrm("allow", "u")
}

func (m *Good) ToRead() []string {
	mapR := m.toOrm("allow", "r")
	keys := make([]string, 0, len(mapR))

	for k := range m.toOrm("allow", "r") {
		keys = append(keys, k)
	}

	return keys
}
