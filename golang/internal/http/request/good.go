package request

type GoodIndex struct {
	Limit  uint `json:"limit" form:"limit" binding:"numeric"`
	Offset uint `json:"offset" form:"offset" binding:"numeric"` // INFO: starts from zero in the technical specification it is written from 1 needs to be clarified
}

type GoodCreate struct {
	ProjectId   uint32 `json:"projectId" binding:"required"` // INFO: why projectId and not project_id
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type GoodUpdate struct {
	Id          uint32 `json:"id" binding:"required"`
	ProjectId   uint32 `json:"projectId" binding:"required"`
	Name        string `json:"name" binding:"min=1"`
	Description string `json:"description"`
}

type GoodDelete struct {
	Id uint32 `uri:"id" binding:"required"` // INFO: need ask TS
}

// INFO: why multi required projectId or id
// INFO: need ask TS
type GoodRePrioritize struct {
	Id        uint32 `json:"id" binding:"required"`
	ProjectId uint32 `json:"projectId" binding:"required"`
	Priority  uint32 `json:"priority" binding:"required"`
}
