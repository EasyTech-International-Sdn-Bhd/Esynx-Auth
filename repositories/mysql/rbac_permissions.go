package mysql

import (
	"errors"
	"github.com/easytech-international-sdn-bhd/esynx-auth/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-auth/entities"
)

type RbacPermissionsRepository struct {
	option *contracts.IRepository
	rp     *RbacRolesPermissionsRepository
}

func NewRbacPermissionsRepository(option *contracts.IRepository) *RbacPermissionsRepository {
	return &RbacPermissionsRepository{
		option,
		NewRbacRolesPermissionsRepository(option),
	}
}

func (r *RbacPermissionsRepository) Get(permCode string) (*entities.RbacPermissions, error) {
	var res *entities.RbacPermissions
	has, err := r.option.Db.Where("permission_code = ?", permCode).Get(&res)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return res, nil
}

func (r *RbacPermissionsRepository) GetMany(permCodes []string) ([]*entities.RbacPermissions, error) {
	var res []*entities.RbacPermissions
	err := r.option.Db.In("permission_codes", permCodes).Find(&res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *RbacPermissionsRepository) GetByName(permName string) (*entities.RbacPermissions, error) {
	var res *entities.RbacPermissions
	has, err := r.option.Db.Where("permission_name = ?", permName).Get(&res)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return res, nil
}

func (r *RbacPermissionsRepository) CreatePermission(perm *entities.RbacPermissions) error {
	res, err := r.Get(perm.PermissionName)
	if err != nil {
		return err
	}
	if res != nil {
		return errors.New("permission already exists")
	}
	perm.ToCreate()
	_, err = r.option.Db.InsertOne(perm)
	if err != nil {
		return err
	}
	return nil
}

func (r *RbacPermissionsRepository) DeletePermission(perm *entities.RbacPermissions) error {
	_, err := r.option.Db.Delete(perm)
	if err != nil {
		return err
	}
	err = r.rp.DeleteByPermission(perm.PermissionCode)
	if err != nil {
		return err
	}
	return nil
}
