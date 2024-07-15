package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-auth/models"
	"github.com/easytech-international-sdn-bhd/esynx-auth/options"
	"xorm.io/xorm/log"
)

type IUserSession interface {
	GetUser() string
	GetApp() string
	GetStore() options.DatabaseStore
	GetConnection() string
	GetJwtSecret() string
	GetRedisConfig() models.RedisConfig
	GetLogger() *log.Logger
}

type IDatabase interface {
	Open(conn string, logger *log.Logger) error
	DefineSchema() error
	Close() error
}
