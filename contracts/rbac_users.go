package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-auth/entities"
	"github.com/easytech-international-sdn-bhd/esynx-auth/models"
)

type IRbacUsers interface {
	Get(userName, password string) (*entities.RbacUsers, error)
	GetByUserCode(userCode string) (*entities.RbacUsers, error)
	CreateUser(info models.CreateRbacUser) error
	UpdateUser(info models.UpdateRbacUser) error
	DeleteUser(info models.DeleteRbacUser) error
	GetServiceAccounts() ([]*entities.RbacUsers, error)
}
