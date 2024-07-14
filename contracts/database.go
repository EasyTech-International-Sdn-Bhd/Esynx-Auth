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
}

type IDatabase interface {
	Open(conn string) error
	DefineSchema() error
	Close() error
}
