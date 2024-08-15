package esynx_auth

import (
	"context"
	"github.com/easytech-international-sdn-bhd/esynx-auth/contracts"
	migrate "github.com/easytech-international-sdn-bhd/esynx-auth/migrate/sql"
	"github.com/easytech-international-sdn-bhd/esynx-auth/options"
	"github.com/easytech-international-sdn-bhd/esynx-auth/repositories/redis"
	"github.com/easytech-international-sdn-bhd/esynx-auth/repositories/sql"
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

func NewEsynxTokenProvider(session contracts.IUserSession) (*EsynxAuth, error) {
	if session.GetStore() == options.SqlDb {
		ctx := context.Background()
		memDb := redis.NewRedis(ctx, session.GetRedisConfig())
		userOptions := contracts.IRepository{
			Db:          nil,
			User:        session.GetUser(),
			AppName:     session.GetApp(),
			JwtSecret:   session.GetJwtSecret(),
			RedisClient: memDb,
		}
		return &EsynxAuth{
			engine:               nil,
			Auth:                 nil,
			RbacPermissions:      nil,
			RbacRoles:            nil,
			RbacRolesPermissions: nil,
			RbacTokens:           sql.NewRbacTokenRepository(&userOptions),
			RbacUserRoles:        nil,
			RbacUsers:            nil,
		}, nil
	} else {

	}
	return nil, nil
}

func NewEsynxAuthProvider(session contracts.IUserSession) (*EsynxAuth, error) {
	if session.GetStore() == options.SqlDb {
		db := sql.NewSqlDb()
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
			Auth:                 sql.NewRbacAuthenticateRepository(&userOptions),
			RbacPermissions:      sql.NewRbacPermissionsRepository(&userOptions),
			RbacRoles:            sql.NewRbacRolesRepository(&userOptions),
			RbacRolesPermissions: sql.NewRbacRolesPermissionsRepository(&userOptions),
			RbacTokens:           sql.NewRbacTokenRepository(&userOptions),
			RbacUserRoles:        sql.NewRbacUserRolesRepository(&userOptions),
			RbacUsers:            sql.NewRbacUsersRepository(&userOptions),
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
