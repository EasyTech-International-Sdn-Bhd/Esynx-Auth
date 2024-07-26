package mysql

import (
	"errors"
	"github.com/easytech-international-sdn-bhd/esynx-auth/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-auth/entities"
)

type RbacRolesRepository struct {
	option *contracts.IRepository
	ur     *RbacUserRolesRepository
	rp     *RbacRolesPermissionsRepository
}

func NewRbacRolesRepository(option *contracts.IRepository) *RbacRolesRepository {
	return &RbacRolesRepository{
		option,
		NewRbacUserRolesRepository(option),
		NewRbacRolesPermissionsRepository(option),
	}
}

func (r *RbacRolesRepository) CreateRole(info *entities.RbacRoles) error {
	res, err := r.Get(info.RoleName)
	if err != nil {
		return err
	}
	if res != nil {
		return errors.New("role already exists")
	}
	info.BeforeInsert()
	_, err = r.option.Db.InsertOne(info)
	if err != nil {
		return err
	}
	return nil
}

func (r *RbacRolesRepository) Get(roleCode string) (*entities.RbacRoles, error) {
	var res *entities.RbacRoles
	has, err := r.option.Db.Where("role_code = ?", roleCode).Get(&res)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return res, nil
}

func (r *RbacRolesRepository) GetMany(roleCodes []string) ([]*entities.RbacRoles, error) {
	var res []*entities.RbacRoles
	err := r.option.Db.In("role_code", roleCodes).Find(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RbacRolesRepository) GetByName(roleName string) (*entities.RbacRoles, error) {
	var res *entities.RbacRoles
	has, err := r.option.Db.Where("role_name = ?", roleName).Get(&res)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return res, nil
}

func (r *RbacRolesRepository) DeleteRole(role *entities.RbacRoles) error {
	_, err := r.option.Db.Delete(role)
	if err != nil {
		return err
	}
	err = r.ur.DeleteByRoleCode(role.RoleCode)
	if err != nil {
		return err
	}
	err = r.rp.DeleteByRoleCode(role.RoleCode)
	if err != nil {
		return err
	}
	return nil
}
