package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-auth/entities"
	"github.com/easytech-international-sdn-bhd/esynx-auth/models"
)

type IRbacRoles interface {
	CreateRole(info models.CreateRbacRole) error
	Get(roleCode string) (*entities.RbacRoles, error)
	GetByName(roleName string) (*entities.RbacRoles, error)
}
