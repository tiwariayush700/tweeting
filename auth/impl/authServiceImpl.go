package authImpl

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"github.com/tiwariayush700/tweeting/auth"
	userError "github.com/tiwariayush700/tweeting/error"
	"github.com/tiwariayush700/tweeting/models"
)

type authImpl struct {
	SecretKey string
}

func NewAuthService(secret string) auth.Service {
	return &authImpl{SecretKey: secret}
}

func (impl *authImpl) GenerateUserToken(userID uint, role string) (string, error) {
	claims := models.UserLoginJWTClaims{
		Authorized: true,
		Id:         userID,
		Role:       role,
		StandardClaims: jwt.StandardClaims{
			Issuer: "admin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(impl.SecretKey))

	return tokenString, err
}

func (impl *authImpl) AuthenticateUser(jwtTokenString string) (*models.UserLoginJWTClaims, error) {

	if len(jwtTokenString) == 0 {
		return nil, userError.ErrorTokenExpected
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(jwtTokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(impl.SecretKey), nil
	})

	if token != nil && !token.Valid {
		return nil, userError.ErrorTokenInvalid
	}

	userLoginJWTClaims := &models.UserLoginJWTClaims{}
	err = auth.GetDataFromTokenClaims(claims, &userLoginJWTClaims)
	if err != nil {
		logrus.Errorf("err decoding payload err: %v", err)
		return nil, userError.ErrorTokenExpected
	}

	return userLoginJWTClaims, nil
}
