package models

type TokenDetails struct {
	AccessToken      string
	RefreshToken     string
	AccessTokenUUID  string
	RefreshTokenUUID string
	UserCode         string
	AtExpires        int64
	RtExpires        int64
}
