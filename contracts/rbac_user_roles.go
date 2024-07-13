package contracts

type IRbacUserRoles interface {
	Assign(roleCode, userCode string) (bool, error)
	Delete(roleCode, userCode string) (bool, error)
}
