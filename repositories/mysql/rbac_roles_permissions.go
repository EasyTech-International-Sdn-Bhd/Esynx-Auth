package mysql

import (
	"github.com/easytech-international-sdn-bhd/esynx-auth/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-auth/entities"
)

type RbacRolesPermissionsRepository struct {
	option *contracts.IRepository
}

func NewRbacRolesPermissionsRepository(option *contracts.IRepository) *RbacRolesPermissionsRepository {
	return &RbacRolesPermissionsRepository{option}
}

func (r *RbacRolesPermissionsRepository) Get(roleCode string) ([]*entities.RbacRolesPermissions, error) {
	var permissions []*entities.RbacRolesPermissions
	err := r.option.Db.Where("role_code = ?", roleCode).Find(&permissions)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *RbacRolesPermissionsRepository) GetMany(roleCodes []string) ([]*entities.RbacRolesPermissions, error) {
	var permissions []*entities.RbacRolesPermissions
	err := r.option.Db.In("role_code", roleCodes).Find(&permissions)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (r *RbacRolesPermissionsRepository) Assign(roleCode, permCode string) error {
	relation := &entities.RbacRolesPermissions{
		RoleCode:       roleCode,
		PermissionCode: permCode,
	}
	_, err := r.option.Db.InsertOne(relation)
	if err != nil {
		return err
	}
	return nil
}

func (r *RbacRolesPermissionsRepository) Delete(roleCode, permCode string) error {
	_, err := r.option.Db.Delete(&entities.RbacRolesPermissions{
		RoleCode:       roleCode,
		PermissionCode: permCode,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *RbacRolesPermissionsRepository) DeleteByPermission(permCode string) error {
	_, err := r.option.Db.Delete(&entities.RbacRolesPermissions{
		PermissionCode: permCode,
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *RbacRolesPermissionsRepository) DeleteByRoleCode(roleCode string) error {
	_, err := r.option.Db.Delete(&entities.RbacRolesPermissions{
		RoleCode: roleCode,
	})
	if err != nil {
		return err
	}
	return nil
}
