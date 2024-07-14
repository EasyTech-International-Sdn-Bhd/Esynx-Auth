package models

import "github.com/easytech-international-sdn-bhd/esynx-auth/entities"

type Authenticated struct {
	AccessToken  string
	RefreshToken string
	Roles        []*entities.RbacRoles
	Permissions  []*entities.RbacPermissions
}
