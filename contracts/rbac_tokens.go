package contracts

import (
	"github.com/easytech-international-sdn-bhd/esynx-auth/models"
	"github.com/golang-jwt/jwt/v5"
)

type IRbacTokens interface {
	CreateToken(claimData *models.RbacTokenClaim) (*models.Authenticated, error)
	TokenClaims(accessToken string) (*models.RbacTokenClaim, error)
	DeleteToken(userCode string) error
	VerifyToken(token string) (*jwt.Token, error)
	IsTokenValid(token string) error
	GetAccessToken(refreshToken string) (string, error)
}
