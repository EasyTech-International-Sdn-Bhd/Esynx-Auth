package entities

import (
	"encoding/base64"
	"github.com/google/uuid"
	"strings"
	"time"
)

type RbacUsers struct {
	Id             uint64    `xorm:"not null pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	UserCode       string    `xorm:"not null unique VARCHAR(80)" json:"userCode,omitempty" xml:"userCode"`
	Username       string    `xorm:"VARCHAR(70)" json:"username,omitempty" xml:"username"`
	Password       string    `xorm:"VARCHAR(100)" json:"password,omitempty" xml:"password"`
	ClientCompany  string    `xorm:"VARCHAR(50)" json:"clientCompany,omitempty" xml:"clientCompany"`
	Metadata       string    `xorm:"JSON" json:"metadata,omitempty" xml:"metadata"`
	ShortCode      string    `xorm:"VARCHAR(50)" json:"shortCode,omitempty" xml:"shortCode"`
	Server         string    `xorm:"VARCHAR(50)" json:"server,omitempty" xml:"server"`
	Deleted        int       `xorm:"default 0 TINYINT(1)" json:"deleted,omitempty" xml:"deleted"`
	DeletedBy      string    `xorm:"VARCHAR(80)" json:"deletedBy,omitempty" xml:"deletedBy"`
	DeletedAt      time.Time `xorm:"DATETIME" json:"deletedAt,omitempty" xml:"deletedAt"`
	CreatedAt      time.Time `xorm:"DATETIME" json:"createdAt,omitempty" xml:"createdAt"`
	CreatedBy      string    `xorm:"VARCHAR(80)" json:"createdBy,omitempty" xml:"createdBy"`
	UpdatedBy      string    `xorm:"VARCHAR(80)" json:"updatedBy,omitempty" xml:"updatedBy"`
	UpdatedAt      time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME" json:"updatedAt,omitempty" xml:"updatedAt"`
	BiDealer       string    `xorm:"VARCHAR(100)" json:"biDealer,omitempty" xml:"biDealer"`
	BiSubscription string    `xorm:"VARCHAR(20)" json:"biSubscription,omitempty" xml:"biSubscription"`
	BiState        string    `xorm:"VARCHAR(20)" json:"biState,omitempty" xml:"biState"`
	BiIndustry     string    `xorm:"VARCHAR(60)" json:"biIndustry,omitempty" xml:"biIndustry"`
}

func (m *RbacUsers) TableName() string {
	return "rbac_users"
}

func (m *RbacUsers) BeforeInsert(createdBy string) {
	m.UserCode = uuid.New().URN()
	m.CreatedAt = time.Now()
	m.CreatedBy = createdBy
	m.UpdatedBy = createdBy
	m.UpdatedAt = time.Now()
	m.ShortCode = strings.ToUpper(base64.URLEncoding.EncodeToString([]byte(m.UserCode[:]))[:8])
}

func (m *RbacUsers) BeforeUpdate(updatedBy string) {
	m.UpdatedBy = updatedBy
	m.UpdatedAt = time.Now()
}

func (m *RbacUsers) ToDelete(deletedBy string) {
	m.DeletedBy = deletedBy
	m.DeletedAt = time.Now()
	m.Deleted = 1
}
