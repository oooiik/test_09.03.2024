package model

import (
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
