package auth

import (
	"github.com/tiwariayush700/tweeting/models"
)

type Service interface {
	GenerateUserToken(userID uint, role string) (string, error)
	AuthenticateUser(jwtTokenString string) (*models.UserLoginJWTClaims, error)
}
