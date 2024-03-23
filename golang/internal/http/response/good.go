package response

import (
	"github.com/oooiik/test_09.03.2024/internal/logger"
	"github.com/oooiik/test_09.03.2024/internal/model"
)

type Good struct {
	Id          uint32 `json:"id"`
	ProjectId   uint32 `json:"projectId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    uint32 `json:"priority"`
	Removed     bool   `json:"removed"`
	CreatedAt   string `json:"createdAt"`
}

func GoodResponse(m *model.Good) *Good {
	logger.Debug(m)
	return &Good{
		Id:          m.Id,
		ProjectId:   m.ProjectId,
		Name:        m.Name,
		Description: m.Description,
		Priority:    m.Priority,
		Removed:     m.Removed,
		CreatedAt:   (*m.CreatedAt).String(),
	}
}

type GoodList struct {
	Goods []*Good `json:"goods"`
}

func GoodListResponse(list []*model.Good) *GoodList {
	l := GoodList{
		Goods: make([]*Good, len(list)),
	}
	for i, v := range list {
		l.Goods[i] = GoodResponse(v)
	}
	return &l
}

type GoodDeleted struct {
	Id        uint32 `json:"id"`
	ProjectId uint32 `json:"projectId"`
	Removed   bool   `json:"removed"`
}

func GoodDeletedResponse(m *model.Good) *GoodDeleted {
	return &GoodDeleted{
		Id:        m.Id,
		ProjectId: m.ProjectId,
		Removed:   m.Removed,
	}
}

type GoodLRePriority struct {
	Id        uint32 `json:"id"`
	ProjectId uint32 `json:"projectId"`
	Priority  uint32 `json:"priority"`
}

func GoodLRePriorityResponse(m *model.Good) *GoodLRePriority {
	return &GoodLRePriority{
		Id:        m.Id,
		ProjectId: m.ProjectId,
		Priority:  m.Priority,
	}
}

type GoodListRePriorities struct {
	Prioritize []*GoodLRePriority `json:"priorities"`
}

func GoodListRePrioritiesResponse(list []*model.Good) *GoodListRePriorities {
	l := GoodListRePriorities{
		Prioritize: make([]*GoodLRePriority, len(list)),
	}
	for i, v := range list {
		l.Prioritize[i] = GoodLRePriorityResponse(v)
	}
	return &l
}
