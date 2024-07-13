package contracts

type IRbacRolesPermissions interface {
	Assign(roleCode, permCode string) (bool, error)
	Delete(roleCode, permCode string) (bool, error)
}
