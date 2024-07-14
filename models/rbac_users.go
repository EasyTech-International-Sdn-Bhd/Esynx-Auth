package models

type CreateRbacUser struct {
	Username        string
	Password        string
	ClientCompany   string
	Metadata        string
	Server          string
	BiDealer        string
	BiSubscriptions string
	BiState         string
	BiIndustry      string
	ShortCode       string
	CreatedBy       string
}

type UpdateRbacUser struct {
	UserCode        string
	Password        string
	Metadata        string
	Server          string
	BiDealer        string
	BiSubscriptions string
	BiState         string
	BiIndustry      string
	UpdatedBy       string
}

type DeleteRbacUser struct {
	UserCode  string
	DeletedBy string
}
