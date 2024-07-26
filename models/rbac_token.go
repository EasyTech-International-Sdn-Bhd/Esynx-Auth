package models

type RbacTokenClaim struct {
	UserCode      string `json:"userCode,omitempty" xml:"userCode"`
	ClientCompany string `json:"clientCompany,omitempty" xml:"clientCompany"`
	Server        string `json:"server,omitempty" xml:"server"`
	Metadata      string `json:"metadata,omitempty" xml:"metadata"`
}
