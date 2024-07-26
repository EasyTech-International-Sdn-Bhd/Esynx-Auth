package entities

type RbacUserRoles struct {
	UserCode string `xorm:"unique(rbac_user_roles_unx) VARCHAR(80)" json:"userCode,omitempty" xml:"userCode"`
	RoleCode string `xorm:"unique(rbac_user_roles_unx) VARCHAR(80)" json:"roleCode,omitempty" xml:"roleCode"`
}

func (m *RbacUserRoles) TableName() string {
	return "rbac_user_roles"
}
