package sql

import (
	"github.com/easytech-international-sdn-bhd/esynx-auth/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-auth/entities"
)

type RbacUserRolesRepository struct {
	option *contracts.IRepository
}

func NewRbacUserRolesRepository(option *contracts.IRepository) *RbacUserRolesRepository {
	return &RbacUserRolesRepository{option}
}

func (r *RbacUserRolesRepository) Get(roleCode string) ([]*entities.RbacUserRoles, error) {
	var roles []*entities.RbacUserRoles
	err := r.option.Db.Where("role_code = ?", roleCode).Find(&roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RbacUserRolesRepository) GetByUser(userCode string) ([]*entities.RbacUserRoles, error) {
	var roles []*entities.RbacUserRoles
	err := r.option.Db.Where("user_code = ?", userCode).Find(&roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (r *RbacUserRolesRepository) Assign(roleCode, userCode string) error {
	relation := entities.RbacUserRoles{
		UserCode: userCode,
		RoleCode: roleCode,
	}
	_, err := r.option.Db.InsertOne(relation)
	if err != nil {
		return err
	}
	return nil
}
func (r *RbacUserRolesRepository) Delete(roleCode, userCode string) error {
	_, err := r.option.Db.Delete(&entities.RbacUserRoles{
		RoleCode: roleCode,
		UserCode: userCode,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *RbacUserRolesRepository) DeleteByUserCode(userCode string) error {
	_, err := r.option.Db.Delete(&entities.RbacUserRoles{
		UserCode: userCode,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *RbacUserRolesRepository) DeleteByRoleCode(roleCode string) error {
	_, err := r.option.Db.Delete(&entities.RbacUserRoles{
		RoleCode: roleCode,
	})
	if err != nil {
		return err
	}
	return nil
}
