package entities

type RbacUserRoles struct {
	UserCode string `xorm:"unique(rbac_user_roles_unx) VARCHAR(80)"`
	RoleCode string `xorm:"unique(rbac_user_roles_unx) VARCHAR(80)"`
}

func (m *RbacUserRoles) TableName() string {
	return "rbac_user_roles"
}
