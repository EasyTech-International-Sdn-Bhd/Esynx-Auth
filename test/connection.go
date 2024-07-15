package test

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-auth/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-auth/models"
	"github.com/easytech-international-sdn-bhd/esynx-auth/options"
)

func NewTestAuthProvider() (contracts.IRepository, *AuthSession) {
	session := NewTestAuthSession("UnitTest")
	return contracts.IRepository{
		Db:          nil,
		User:        session.GetUser(),
		AppName:     session.GetApp(),
		JwtSecret:   session.GetJwtSecret(),
		RedisClient: nil,
	}, session
}

type AuthSession struct {
	user string
}

func NewTestAuthSession(user string) *AuthSession {
	return &AuthSession{
		user,
	}
}

func (a *AuthSession) GetUser() string {
	return a.user
}

func (a *AuthSession) GetApp() string {
	return "UnitTest"
}

func (a *AuthSession) GetStore() options.DatabaseStore {
	return options.MySQL
}

func (a *AuthSession) GetConnection() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=2s",
		"root",
		"mysql",
		"localhost",
		3306,
		"esynx",
	)
}

func (a *AuthSession) GetLogger() contracts.IDatabaseLogger {
	return nil
}

func (a *AuthSession) GetJwtSecret() string {
	return ""
}

func (a *AuthSession) GetRedisConfig() models.RedisConfig {
	return models.RedisConfig{}
}
