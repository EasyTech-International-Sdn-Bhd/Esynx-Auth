package mysql

import (
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-auth/contracts"
	"github.com/easytech-international-sdn-bhd/esynx-auth/entities"
	"github.com/easytech-international-sdn-bhd/esynx-auth/models"
	"github.com/easytech-international-sdn-bhd/esynx-auth/utils"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type RbacTokenRepository struct {
	option *contracts.IRepository
}

func NewRbacTokenRepository(option *contracts.IRepository) *RbacTokenRepository {
	return &RbacTokenRepository{option}
}

func (r *RbacTokenRepository) CreateToken(claimData *models.RbacTokenClaim) (*models.Authenticated, error) {
	atExpUnix := time.Now().Add(time.Minute * 15).Unix()
	rtExpUnix := time.Now().Add(time.Hour * 24 * 7).Unix()
	td := &models.TokenDetails{
		AccessToken:      "",
		RefreshToken:     "",
		AccessTokenUUID:  utils.AccessTokenKey(claimData.UserCode),
		RefreshTokenUUID: utils.RefreshTokenKey(claimData.UserCode),
		UserCode:         claimData.UserCode,
		AtExpires:        atExpUnix,
		RtExpires:        rtExpUnix,
	}

	var err error
	claim := jwt.MapClaims{}
	claim["authorized"] = true
	claim["user_code"] = claimData.UserCode
	claim["metadata"] = claimData.Metadata
	claim["client_company"] = claimData.ClientCompany
	claim["exp"] = td.AtExpires

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	td.AccessToken, err = at.SignedString([]byte(r.option.JwtSecret))
	if err != nil {
		return nil, err
	}
	rtClaims := jwt.MapClaims{}
	rtClaims["user_code"] = claimData.UserCode
	rtClaims["exp"] = td.RtExpires

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(r.option.JwtSecret))
	if err != nil {
		return nil, err
	}

	atUtc := time.Unix(atExpUnix, 0)
	rtUtc := time.Unix(rtExpUnix, 0)
	now := time.Now()

	err = r.option.RedisClient.SetToken(td.AccessTokenUUID, td.AccessToken, atUtc.Sub(now))
	if err != nil {
		return nil, err
	}
	err = r.option.RedisClient.SetToken(td.RefreshTokenUUID, td.RefreshToken, rtUtc.Sub(now))
	if err != nil {
		return nil, err
	}
	_, err = r.option.Db.InsertOne(&entities.RbacTokens{
		UserCode:     claimData.UserCode,
		RefreshToken: td.RefreshToken,
		IssuedAt:     time.Now(),
		ExpiresAt:    time.Unix(rtExpUnix, 0),
	})
	if err != nil {
		return nil, err
	}
	return &models.Authenticated{
		AccessToken:  td.AccessToken,
		RefreshToken: td.RefreshToken,
		Roles:        nil,
		Permissions:  nil,
	}, nil
}

func (r *RbacTokenRepository) RefreshTokenClaims(refreshToken string) (*string, error) {
	token, err := r.VerifyToken(refreshToken)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userCode, ok := claims["user_code"].(string)
		if !ok {
			return nil, fmt.Errorf("user_code not found in claims")
		}
		return &userCode, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func (r *RbacTokenRepository) TokenClaims(accessToken string) (*models.RbacTokenClaim, error) {
	token, err := r.VerifyToken(accessToken)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userCode, ok := claims["user_code"].(string)
		if !ok {
			return nil, fmt.Errorf("user_code not found in claims")
		}
		clientCompany, ok := claims["client_company"].(string)
		if !ok {
			return nil, fmt.Errorf("client_company not found in claims")
		}
		metadata, ok := claims["metadata"].(string)
		if !ok {
			return nil, fmt.Errorf("metadata not found in claims")
		}
		return &models.RbacTokenClaim{
			UserCode:      userCode,
			ClientCompany: clientCompany,
			Metadata:      metadata,
		}, nil
	}
	return nil, err
}

func (r *RbacTokenRepository) VerifyToken(uToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(uToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return r.option.JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (r *RbacTokenRepository) IsTokenValid(uToken string) error {
	token, err := r.VerifyToken(uToken)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func (r *RbacTokenRepository) DeleteToken(userCode string) error {
	accessTokenKey := utils.AccessTokenKey(userCode)
	return r.option.RedisClient.DelToken(accessTokenKey)
}

func (r *RbacTokenRepository) GetAccessToken(refreshToken string) (*string, error) {
	token, err := r.VerifyToken(refreshToken)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, fmt.Errorf("token claim corrupted")
	}
	userCode, ok := claims["user_code"].(string)
	if !ok {
		return nil, fmt.Errorf("user_code not found in claims")
	}
	accessTokenKey := utils.AccessTokenKey(userCode)
	res, err := r.option.RedisClient.GetToken(accessTokenKey)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
