package entities

import (
	"github.com/google/uuid"
	"time"
)

type RbacRoles struct {
	Id          uint64    `xorm:"not null pk autoincr unique UNSIGNED BIGINT"`
	RoleCode    string    `xorm:"not null unique VARCHAR(80)"`
	RoleName    string    `xorm:"VARCHAR(50)"`
	Description string    `xorm:"VARCHAR(255)"`
	CreatedAt   time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
}

func (m *RbacRoles) TableName() string {
	return "rbac_roles"
}

func (m *RbacRoles) ToCreate() {
	m.RoleCode = uuid.New().URN()
	m.CreatedAt = time.Now()
}
