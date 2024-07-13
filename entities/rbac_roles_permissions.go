package entities

type RbacRolesPermissions struct {
	Id             uint64 `xorm:"not null pk autoincr unique UNSIGNED BIGINT"`
	RoleCode       string `xorm:"unique(rbac_roles_permissions_unx) VARCHAR(80)"`
	PermissionCode string `xorm:"unique(rbac_roles_permissions_unx) VARCHAR(80)"`
}

func (m *RbacRolesPermissions) TableName() string {
	return "rbac_roles_permissions"
}
