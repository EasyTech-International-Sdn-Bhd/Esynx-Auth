package esynx_auth

import (
	"context"
	"github.com/easytech-international-sdn-bhd/esynx-auth/contracts"
	migrate "github.com/easytech-international-sdn-bhd/esynx-auth/migrate/sql"
	"github.com/easytech-international-sdn-bhd/esynx-auth/options"
	"github.com/easytech-international-sdn-bhd/esynx-auth/repositories/mysql"
	"github.com/easytech-international-sdn-bhd/esynx-auth/repositories/redis"
)

type EsynxAuth struct {
	engine               contracts.IDatabase
	Auth                 contracts.IRbacAuthenticate
	RbacPermissions      contracts.IRbacPermissions
	RbacRoles            contracts.IRbacRoles
	RbacRolesPermissions contracts.IRbacRolesPermissions
	RbacTokens           contracts.IRbacTokens
	RbacUserRoles        contracts.IRbacUserRoles
	RbacUsers            contracts.IRbacUsers
}

func NewEsynxTokenProvider(session contracts.IUserSession) contracts.IRbacTokens {
	ctx := context.Background()
	memDb := redis.NewRedis(ctx, session.GetRedisConfig())
	userOptions := contracts.IRepository{
		Db:          nil,
		User:        session.GetUser(),
		AppName:     session.GetApp(),
		JwtSecret:   session.GetJwtSecret(),
		RedisClient: memDb,
	}
	return mysql.NewRbacTokenRepository(&userOptions)
}

func NewEsynxAuthProvider(session contracts.IUserSession) (*EsynxAuth, error) {
	if session.GetStore() == options.MySQL {
		db := mysql.NewMySqlDb()
		err := db.Open(session.GetConnection(), session.GetLogger())
		if err != nil {
			return nil, err
		}
		err = migrate.DefineSchema(db.Engine)
		if err != nil {
			return nil, err
		}
		ctx := context.Background()
		memDb := redis.NewRedis(ctx, session.GetRedisConfig())
		userOptions := contracts.IRepository{
			Db:          db.Engine,
			User:        session.GetUser(),
			AppName:     session.GetApp(),
			JwtSecret:   session.GetJwtSecret(),
			RedisClient: memDb,
		}
		return &EsynxAuth{
			engine:               db,
			Auth:                 mysql.NewRbacAuthenticateRepository(&userOptions),
			RbacPermissions:      mysql.NewRbacPermissionsRepository(&userOptions),
			RbacRoles:            mysql.NewRbacRolesRepository(&userOptions),
			RbacRolesPermissions: mysql.NewRbacRolesPermissionsRepository(&userOptions),
			RbacTokens:           mysql.NewRbacTokenRepository(&userOptions),
			RbacUserRoles:        mysql.NewRbacUserRolesRepository(&userOptions),
			RbacUsers:            mysql.NewRbacUsersRepository(&userOptions),
		}, nil
	}
	if session.GetStore() == options.Firestore {

	}
	return nil, nil
}

func (e *EsynxAuth) Destroy() error {
	if e.engine == nil {
		return nil
	}
	return e.engine.Close()
}
