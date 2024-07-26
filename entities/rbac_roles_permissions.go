package entities

type RbacRolesPermissions struct {
	Id             uint64 `xorm:"not null pk autoincr unique UNSIGNED BIGINT" json:"id,omitempty" xml:"id"`
	RoleCode       string `xorm:"unique(rbac_roles_permissions_unx) VARCHAR(80)" json:"roleCode,omitempty" xml:"roleCode"`
	PermissionCode string `xorm:"unique(rbac_roles_permissions_unx) VARCHAR(80)" json:"permissionCode,omitempty" xml:"permissionCode"`
}

func (m *RbacRolesPermissions) TableName() string {
	return "rbac_roles_permissions"
}
