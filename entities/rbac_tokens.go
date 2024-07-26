package entities

import (
	"time"
)

type RbacTokens struct {
	Id           uint64    `xorm:"not null pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	UserCode     string    `xorm:"not null index unique(rbac_tokens_unx) unique VARCHAR(80)" json:"userCode,omitempty" xml:"userCode"`
	RefreshToken string    `xorm:"not null index unique(rbac_tokens_unx) VARCHAR(255)" json:"refreshToken,omitempty" xml:"refreshToken"`
	IssuedAt     time.Time `xorm:"default CURRENT_TIMESTAMP TIMESTAMP" json:"issuedAt,omitempty" xml:"issuedAt"`
	ExpiresAt    time.Time `xorm:"not null TIMESTAMP" json:"expiresAt,omitempty" xml:"expiresAt"`
}

func (m *RbacTokens) TableName() string {
	return "rbac_tokens"
}
