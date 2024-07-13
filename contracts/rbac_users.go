package contracts

import "github.com/easytech-international-sdn-bhd/esynx-auth/models"

type IRbacUsers interface {
	CreateUser(info models.CreateRbacUser) error
	UpdateUser(info models.UpdateRbacUser) error
	DeleteUser(info models.DeleteRbacUser) error
}
