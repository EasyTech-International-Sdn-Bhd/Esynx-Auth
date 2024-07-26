package entities

import (
	"github.com/google/uuid"
	"time"
)

type RbacPermissions struct {
	Id             uint64    `xorm:"not null pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	PermissionCode string    `xorm:"not null unique VARCHAR(80)" json:"permissionCode,omitempty" xml:"permissionCode"`
	PermissionName string    `xorm:"not null unique VARCHAR(50)" json:"permissionName,omitempty" xml:"permissionName"`
	Description    string    `xorm:"VARCHAR(255)" json:"description,omitempty" xml:"description"`
	CreatedAt      time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP" json:"createdAt,omitempty" xml:"createdAt"`
}

func (m *RbacPermissions) TableName() string {
	return "rbac_permissions"
}

func (m *RbacPermissions) BeforeInsert() {
	m.PermissionCode = uuid.New().URN()
}
