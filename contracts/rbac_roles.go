package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-auth/entities"
)

type IRbacRoles interface {
	Get(roleCode string) (*entities.RbacRoles, error)
	GetMany(roleCodes []string) ([]*entities.RbacRoles, error)
	CreateRole(info *entities.RbacRoles) error
	GetByName(roleName string) (*entities.RbacRoles, error)
	DeleteRole(role *entities.RbacRoles) error
}
