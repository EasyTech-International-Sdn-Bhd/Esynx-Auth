package contracts

import "github.com/easytech-international-sdn-bhd/esynx-auth/models"

type IRbacAuthenticate interface {
	Authenticate(userName, password string) (*models.Authenticated, error)
	RefreshAuthentication(refreshToken string) (*models.RefreshAuthentication, error)
}
