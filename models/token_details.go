package models

type TokenDetails struct {
	AccessToken      string `json:"accessToken,omitempty" xml:"accessToken"`
	RefreshToken     string `json:"refreshToken,omitempty" xml:"refreshToken"`
	AccessTokenUUID  string `json:"accessTokenUUID,omitempty" xml:"accessTokenUUID"`
	RefreshTokenUUID string `json:"refreshTokenUUID,omitempty" xml:"refreshTokenUUID"`
	UserCode         string `json:"userCode,omitempty" xml:"userCode"`
	Server           string `json:"server,omitempty" xml:"server"`
	AtExpires        int64  `json:"atExpires,omitempty" xml:"atExpires"`
	RtExpires        int64  `json:"rtExpires,omitempty" xml:"rtExpires"`
}
