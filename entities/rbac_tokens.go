package entities

import (
	"time"
)

type RbacTokens struct {
	Id           uint64    `xorm:"not null pk autoincr unique UNSIGNED BIGINT"`
	UserCode     string    `xorm:"not null index unique(rbac_tokens_unx) unique VARCHAR(80)"`
	RefreshToken string    `xorm:"not null index unique(rbac_tokens_unx) VARCHAR(255)"`
	IssuedAt     time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP"`
	ExpiresAt    time.Time `xorm:"not null TIMESTAMP"`
}

func (m *RbacTokens) TableName() string {
	return "rbac_tokens"
}
