package entities

import (
	"github.com/google/uuid"
	"time"
)

type RbacRoles struct {
	Id          uint64    `xorm:"not null pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	RoleCode    string    `xorm:"not null unique VARCHAR(80)" json:"roleCode,omitempty" xml:"roleCode"`
	RoleName    string    `xorm:"VARCHAR(50)" json:"roleName,omitempty" xml:"roleName"`
	Description string    `xorm:"VARCHAR(255)" json:"description,omitempty" xml:"description"`
	CreatedAt   time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP" json:"createdAt,omitempty" xml:"createdAt"`
}

func (m *RbacRoles) TableName() string {
	return "rbac_roles"
}

func (m *RbacRoles) BeforeInsert() {
	m.RoleCode = uuid.New().URN()
	m.CreatedAt = time.Now()
}
