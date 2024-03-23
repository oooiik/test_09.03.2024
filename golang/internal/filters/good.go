package filters

import (
	"reflect"
)

type Good struct {
	Id        *uint32 `db:"id"`
	ProjectId *uint32 `db:"project_id"`
	Name      *string `db:"name"`
	//Description *string     `db:"description"`
	Priority *uint32 `db:"priority"`
	Removed  *bool   `db:"removed"`
	//CreatedAt   *time.Time `db:"created_at"`

	Limit  *uint
	Offset *uint
}

func (m *Good) ToFilters() map[string]interface{} {
	filters := make(map[string]interface{})

	refTypeOf := reflect.TypeOf(*m)
	refValOf := reflect.ValueOf(m).Elem()
	for i := 0; i < refTypeOf.NumField(); i++ {
		field := refTypeOf.Field(i)
		tag := field.Tag.Get("db")
		if tag == "" {
			continue
		}
		if !refValOf.Field(i).IsNil() {
			filters[tag] = refValOf.Field(i).Elem().Interface()
		}
	}
	return filters
}
