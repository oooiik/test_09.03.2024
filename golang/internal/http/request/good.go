package request

type GoodIndex struct {
	Limit  uint32 `json:"limit" uri:"limit" binding:"numeric,min=1"`
	Offset uint32 `json:"offset" uri:"limit" binding:"numeric,min=1"`
}

type GoodCreate struct {
	ProjectId   uint32 `json:"project_id" uri:"project_id" binding:"required"`
	Name        string `json:"name" uri:"name" binding:"required"`
	Description string `db:"description"`
	Priority    uint32 `db:"priority"`
}

type GoodUpdate struct {
	Id          uint32 `json:"id" uri:"id"`
	ProjectId   uint32 `json:"project_id" uri:"project_id"`
	Name        string `json:"name" uri:"name"`
	Description string `db:"description"`
	Priority    uint32 `db:"priority"`
}

type GoodDelete struct {
	Id uint32 `json:"id" uri:"id" binding:"required"`
}

type GoodRePrioritize struct {
	Priority uint32 `db:"priority"`
}
