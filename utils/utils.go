package utils

import "fmt"

func AccessTokenKey(key string) string {
	return fmt.Sprintf("AccessToken:%s", key)
}

func RefreshTokenKey(key string) string {
	return fmt.Sprintf("RefreshToken:%s", key)
}
