package entities

import (
	"encoding/base64"
	"github.com/google/uuid"
	"strings"
	"time"
)

type RbacUsers struct {
	Id             uint64    `xorm:"not null pk autoincr unique UNSIGNED BIGINT"`
	UserCode       string    `xorm:"not null unique VARCHAR(80)"`
	Username       string    `xorm:"VARCHAR(70)"`
	Password       string    `xorm:"VARCHAR(100)"`
	ClientCompany  string    `xorm:"VARCHAR(50)"`
	Metadata       string    `xorm:"JSON"`
	ShortCode      string    `xorm:"VARCHAR(50)"`
	Server         string    `xorm:"VARCHAR(50)"`
	Deleted        int       `xorm:"default 0 TINYINT(1)"`
	DeletedBy      string    `xorm:"VARCHAR(80)"`
	DeletedAt      time.Time `xorm:"DATETIME"`
	CreatedAt      time.Time `xorm:"DATETIME"`
	CreatedBy      string    `xorm:"VARCHAR(80)"`
	UpdatedBy      string    `xorm:"VARCHAR(80)"`
	UpdatedAt      time.Time `xorm:"default CURRENT_TIMESTAMP DATETIME"`
	BiDealer       string    `xorm:"VARCHAR(100)"`
	BiSubscription string    `xorm:"VARCHAR(20)"`
	BiState        string    `xorm:"VARCHAR(20)"`
	BiIndustry     string    `xorm:"VARCHAR(60)"`
}

func (m *RbacUsers) TableName() string {
	return "rbac_users"
}

func (m *RbacUsers) ToCreate(createdBy string) {
	m.UserCode = uuid.New().URN()
	m.CreatedAt = time.Now()
	m.CreatedBy = createdBy
	m.UpdatedBy = createdBy
	m.UpdatedAt = time.Now()
	m.ShortCode = strings.ToUpper(base64.URLEncoding.EncodeToString([]byte(m.UserCode[:]))[:8])
}

func (m *RbacUsers) ToUpdate(updatedBy string) {
	m.UpdatedBy = updatedBy
	m.UpdatedAt = time.Now()
}

func (m *RbacUsers) ToDelete(deletedBy string) {
	m.DeletedBy = deletedBy
	m.DeletedAt = time.Now()
	m.Deleted = 1
}
