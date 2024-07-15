package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-auth/models"
	"github.com/easytech-international-sdn-bhd/esynx-auth/options"
)

type IUserSession interface {
	GetUser() string
	GetApp() string
	GetStore() options.DatabaseStore
	GetConnection() string
	GetJwtSecret() string
	GetRedisConfig() models.RedisConfig
	GetLogger() *interface{}
}

type IDatabase interface {
	Open(conn string, logger *interface{}) error
	DefineSchema() error
	Close() error
}
