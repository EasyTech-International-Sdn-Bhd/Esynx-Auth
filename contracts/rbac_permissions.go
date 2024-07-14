package contracts

import "github.com/easytech-international-sdn-bhd/esynx-auth/entities"

type IRbacPermissions interface {
	Get(permCode string) (*entities.RbacPermissions, error)
	GetMany(permCodes []string) ([]*entities.RbacPermissions, error)
	GetByName(permName string) (*entities.RbacPermissions, error)
	CreatePermission(perm *entities.RbacPermissions) error
	DeletePermission(perm *entities.RbacPermissions) error
}
