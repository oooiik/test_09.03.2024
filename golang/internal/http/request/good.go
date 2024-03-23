package request

type GoodIndex struct {
	Limit  uint32 `json:"limit" form:"limit" binding:"numeric"`
	Offset uint32 `json:"offset" form:"offset" binding:"numeric"`
}

// INFO: why projectId and not project_id
type GoodCreate struct {
	ProjectId   uint32 `json:"projectId" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type GoodUpdate struct {
	Id          uint32 `json:"id" binding:"required"`
	ProjectId   uint32 `json:"projectId" binding:"required"`
	Name        string `json:"name"  binding:""`
	Description string `json:"description"`
}

type GoodDelete struct {
	Id uint32 `uri:"id" binding:"required"`
}

// INFO: why multi required projectId or id
type GoodRePrioritize struct {
	Id        uint32 `json:"id" binding:"required"`
	ProjectId uint32 `json:"projectId" binding:"required"`
	Priority  uint32 `json:"priority" binding:"required"`
}
