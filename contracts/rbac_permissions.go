package contracts

import "github.com/easytech-international-sdn-bhd/esynx-auth/entities"

type IRbacPermissions interface {
	CreatePermission(perm *entities.RbacPermissions) error
	Get(permCode string) (*entities.RbacPermissions, bool)
	GetByName(permName string) (*entities.RbacPermissions, bool)
}
