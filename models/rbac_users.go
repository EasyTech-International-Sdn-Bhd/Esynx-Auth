package models

type CreateRbacUser struct {
	Username        string `json:"username,omitempty" xml:"username"`
	Password        string `json:"password,omitempty" xml:"password"`
	ClientCompany   string `json:"clientCompany,omitempty" xml:"clientCompany"`
	Metadata        string `json:"metadata,omitempty" xml:"metadata"`
	Server          string `json:"server,omitempty" xml:"server"`
	BiDealer        string `json:"biDealer,omitempty" xml:"biDealer"`
	BiSubscriptions string `json:"biSubscriptions,omitempty" xml:"biSubscriptions"`
	BiState         string `json:"biState,omitempty" xml:"biState"`
	BiIndustry      string `json:"biIndustry,omitempty" xml:"biIndustry"`
	CreatedBy       string `json:"createdBy,omitempty" xml:"createdBy"`
}

type UpdateRbacUser struct {
	UserCode        string `json:"userCode,omitempty" xml:"userCode"`
	Password        string `json:"password,omitempty" xml:"password"`
	Metadata        string `json:"metadata,omitempty" xml:"metadata"`
	Server          string `json:"server,omitempty" xml:"server"`
	BiDealer        string `json:"biDealer,omitempty" xml:"biDealer"`
	BiSubscriptions string `json:"biSubscriptions,omitempty" xml:"biSubscriptions"`
	BiState         string `json:"biState,omitempty" xml:"biState"`
	BiIndustry      string `json:"biIndustry,omitempty" xml:"biIndustry"`
	UpdatedBy       string `json:"updatedBy,omitempty" xml:"updatedBy"`
}

type DeleteRbacUser struct {
	UserCode  string `json:"userCode,omitempty" xml:"userCode"`
	DeletedBy string `json:"deletedBy,omitempty" xml:"deletedBy"`
}
