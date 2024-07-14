package entities

import (
	"github.com/google/uuid"
	"time"
)

type RbacPermissions struct {
	Id             uint64    `xorm:"not null pk autoincr unique UNSIGNED BIGINT"`
	PermissionCode string    `xorm:"not null unique VARCHAR(80)"`
	PermissionName string    `xorm:"not null unique VARCHAR(50)"`
	Description    string    `xorm:"VARCHAR(255)"`
	CreatedAt      time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (m *RbacPermissions) TableName() string {
	return "rbac_permissions"
}

func (m *RbacPermissions) ToCreate() {
	m.PermissionCode = uuid.New().URN()
}
