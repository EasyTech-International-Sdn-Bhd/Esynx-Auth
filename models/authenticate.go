package models

import "github.com/easytech-international-sdn-bhd/esynx-auth/entities"

type Authenticated struct {
	AccessToken  string                      `json:"accessToken,omitempty" xml:"accessToken"`
	RefreshToken string                      `json:"refreshToken,omitempty" xml:"refreshToken"`
	Roles        []*entities.RbacRoles       `json:"roles,omitempty" xml:"roles"`
	Permissions  []*entities.RbacPermissions `json:"permissions,omitempty" xml:"permissions"`
}

type RefreshAuthentication struct {
	AccessToken string `json:"accessToken,omitempty" xml:"accessToken"`
}
