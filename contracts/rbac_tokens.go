package contracts

import "github.com/easytech-international-sdn-bhd/esynx-auth/entities"

type IRbacTokens interface {
	CreateToken(userCode string, claimData map[string]interface{}) error
	DeleteToken(userCode string) error
	GetToken(userCode string) (string, error)
	GetUserByToken(token string) (*entities.RbacUsers, error)
}
