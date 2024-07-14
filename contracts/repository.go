package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-auth/repositories/redis"
	"xorm.io/xorm"
)

type IRepository struct {
	Db          *xorm.Engine
	User        string
	AppName     string
	JwtSecret   string
	RedisClient *redis.RedisInstance
}
