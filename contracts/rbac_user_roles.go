package contracts

import "github.com/easytech-international-sdn-bhd/esynx-auth/entities"

type IRbacUserRoles interface {
	Get(roleCode string) ([]*entities.RbacUserRoles, error)
	GetByUser(userCode string) ([]*entities.RbacUserRoles, error)
	Assign(roleCode, userCode string) error
	Delete(roleCode, userCode string) error
	DeleteByUserCode(userCode string) error
	DeleteByRoleCode(roleCode string) error
}
