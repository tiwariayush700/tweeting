package controller

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tiwariayush700/tweeting/auth"
	"github.com/tiwariayush700/tweeting/constants"
	"github.com/tiwariayush700/tweeting/models"
	"net/http"
	"os"
	"time"
)

func VerifyUserAndServe(authSvc auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		claims := &models.UserLoginJWTClaims{}

		if os.Getenv("environment") == "local" {
			claims = &models.UserLoginJWTClaims{
				Authorized:     true,
				Id:             9999,
				StandardClaims: jwt.StandardClaims{},
			}
		} else {
			claims, err = authSvc.AuthenticateUser(c.GetHeader("Authorization"))
			if err != nil {
				c.JSON(http.StatusUnauthorized, err)
				return
			}
		}

		ctx := context.WithValue(c.Request.Context(), constants.ContextKeyUserId, claims.Id)
		c.Request = c.Request.WithContext(ctx)
		c.Set(constants.ContextKeyUserId, claims.Id)

		ctx = context.WithValue(c.Request.Context(), constants.ContextKeyRole, claims.Role)
		c.Request = c.Request.WithContext(ctx)
		c.Set(constants.ContextKeyRole, claims.Role)

		c.Next()
	}
}

func VerifyAdminAndServe(authSvc auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {

		_, role, err := getUserIdAndRoleFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		if role != string(models.UserRoleAdmin) && role != string(models.UserRoleSuperAdmin) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not authorised to use this service",
			})
			return
		}

		c.Next()

	}
}

func VerifySuperAdminAndServe() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, role, err := getUserIdAndRoleFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		if role != string(models.UserRoleSuperAdmin) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "You are not authorised to use this service",
			})
			return
		}

		c.Next()

	}
}

func getUserIdAndRoleFromContext(c *gin.Context) (uint, string, error) {

	userIdInterface, ok := c.Get(constants.ContextKeyUserId)
	if !ok {
		return 0, "", fmt.Errorf("unable to authorize user")
	}

	userID, ok := userIdInterface.(uint)
	if !ok {
		return 0, "", fmt.Errorf("unable to convert userid type")
	}

	roleInterface, ok := c.Get(constants.ContextKeyRole)
	if !ok {
		return 0, "", fmt.Errorf("unable to authorize role")
	}

	role, ok := roleInterface.(string)
	if !ok {
		return 0, "", fmt.Errorf("unable to convert role type")
	}

	logrus.Infof("User with user id : %d and role : %s currently logged in at %s", userID, role, time.Now().String())

	return userID, role, nil
}
