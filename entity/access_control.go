package entity

type AccessControl struct {
	ID           uint
	ActorID      uint
	ActorType    ActorType
	PermissionID uint
}

type ActorType string

const (
	RoleActoreType = "role"
	UserActoreType = "user"
)
