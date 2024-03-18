package model

import "time"

type Good struct {
	Id          int32     `db:"id"`
	ProjectId   int32     `db:"project_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Priority    int32     `db:"priority"`
	Removed     bool      `db:"removed"`
	CreatedAt   time.Time `db:"created_at"`
}
