package contracts

import "github.com/easytech-international-sdn-bhd/esynx-auth/entities"

type IRbacRolesPermissions interface {
	Get(roleCode string) ([]*entities.RbacRolesPermissions, error)
	GetMany(roleCodes []string) ([]*entities.RbacRolesPermissions, error)
	Assign(roleCode, permCode string) error
	Delete(roleCode, permCode string) error
	DeleteByPermission(permCode string) error
	DeleteByRoleCode(roleCode string) error
}
