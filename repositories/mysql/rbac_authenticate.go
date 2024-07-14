package mysql

import (
	"errors"
	"github.com/easytech-international-sdn-bhd/esynx-auth/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-auth/entities"
	"github.com/easytech-international-sdn-bhd/esynx-auth/models"
	iterator "github.com/ledongthuc/goterators"
)

type RbacAuthenticateRepository struct {
	option *contracts.IRepository
	u      *RbacUsersRepository
	ur     *RbacUserRolesRepository
	rl     *RbacRolesRepository
	rp     *RbacRolesPermissionsRepository
	p      *RbacPermissionsRepository
	t      *RbacTokenRepository
}

func NewRbacAuthenticateRepository(option *contracts.IRepository) *RbacAuthenticateRepository {
	return &RbacAuthenticateRepository{
		option,
		NewRbacUsersRepository(option),
		NewRbacUserRolesRepository(option),
		NewRbacRolesRepository(option),
		NewRbacRolesPermissionsRepository(option),
		NewRbacPermissionsRepository(option),
		NewRbacTokenRepository(option),
	}
}

func (r *RbacAuthenticateRepository) Authenticate(userName, password string) (*models.Authenticated, error) {
	user, err := r.u.Get(userName, password)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	tokenPair, err := r.t.CreateToken(&models.RbacTokenClaim{
		UserCode:      user.UserCode,
		ClientCompany: user.ClientCompany,
		Metadata:      user.Metadata,
		Server:        user.Server,
	})
	if err != nil {
		return nil, err
	}
	err = r.GetUserRolesPermission(user, tokenPair)
	if err != nil {
		return nil, err
	}
	return tokenPair, nil
}

func (r *RbacAuthenticateRepository) RefreshAuthentication(refreshToken string) (*models.Authenticated, error) {
	userCode, err := r.t.RefreshTokenClaims(refreshToken)
	if err != nil {
		return nil, err
	}
	if userCode == nil {
		return nil, errors.New("invalid refresh token")
	}
	user, err := r.u.GetByUserCode(*userCode)
	if err != nil {
		return nil, err
	}
	tokenPair, err := r.t.CreateToken(&models.RbacTokenClaim{
		UserCode:      user.UserCode,
		ClientCompany: user.ClientCompany,
		Metadata:      user.Metadata,
		Server:        user.Server,
	})
	if err != nil {
		return nil, err
	}
	err = r.GetUserRolesPermission(user, tokenPair)
	if err != nil {
		return nil, err
	}
	return tokenPair, nil
}

func (r *RbacAuthenticateRepository) GetUserRolesPermission(user *entities.RbacUsers, result *models.Authenticated) error {
	userRoles, err := r.ur.GetByUser(user.UserCode)
	if err != nil {
		return err
	}
	rolePerm, err := r.rp.GetMany(iterator.Map(userRoles, func(item *entities.RbacUserRoles) string {
		return item.RoleCode
	}))
	if err != nil {
		return err
	}
	permissions, err := r.p.GetMany(iterator.Map(rolePerm, func(item *entities.RbacRolesPermissions) string {
		return item.PermissionCode
	}))
	if err != nil {
		return err
	}
	roles, err := r.rl.GetMany(iterator.Map(userRoles, func(item *entities.RbacUserRoles) string {
		return item.RoleCode
	}))
	result.Permissions = permissions
	result.Roles = roles

	return nil
}
