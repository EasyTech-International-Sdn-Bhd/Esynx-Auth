package models

type TokenDetails struct {
	AccessToken      string
	RefreshToken     string
	AccessTokenUUID  string
	RefreshTokenUUID string
	UserCode         string
	Server           string
	AtExpires        int64
	RtExpires        int64
}
