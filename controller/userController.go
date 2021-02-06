package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tiwariayush700/tweeting/auth"
	userError "github.com/tiwariayush700/tweeting/error"
	"github.com/tiwariayush700/tweeting/models"
	"github.com/tiwariayush700/tweeting/services"
	"net/http"
	"strconv"
)

type userController struct {
	service     services.UserService
	app         *app
	authService auth.Service
}

func NewUserController(service services.UserService, authService auth.Service, app *app) *userController {
	return &userController{service: service, authService: authService, app: app}
}

func (u *userController) RegisterRoutes() {
	router := u.app.Router
	userRouterGroup := router.Group("/user")
	{
		userRouterGroup.POST("/register", u.Register())
		userRouterGroup.POST("/login", u.Login())

		routerGroupVerified := userRouterGroup.Use(VerifyUserAndServe(u.authService))
		routerGroupVerified.GET("/me", u.GetUserProfile())

		routerGroupVerified.Use(VerifyAdminAndServe(u.authService))
		routerGroupVerified.GET("/fetch/:user_id")
		routerGroupVerified.PUT("/role/:user_id", u.UpdateUserRole())
	}
}

func (u *userController) Register() gin.HandlerFunc {
	return func(c *gin.Context) {

		user := &models.User{}
		err := c.ShouldBind(user)
		if err != nil {
			errRes := userError.NewErrorBadRequest(err, "invalid input")
			c.JSON(http.StatusBadRequest, errRes)
			return
		}

		err = u.service.CreateUser(c, user)
		if err != nil {
			errRes := userError.NewErrorInternal(err, "something went wrong")
			c.JSON(http.StatusInternalServerError, errRes)
			return
		}

		//todo return jwt token
		token, err := u.authService.GenerateUserToken(user.ID, user.Role)
		if err != nil {
			errRes := userError.NewErrorForbidden(err, "unable to generate token")
			c.JSON(http.StatusUnauthorized, errRes)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user":  user,
			"token": token,
		})
	}
}

func (u *userController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {

		params := models.LoginRequest{}
		err := c.ShouldBind(&params)
		if err != nil {
			errRes := userError.NewErrorBadRequest(err, "invalid input")
			c.JSON(http.StatusBadRequest, errRes)
			return
		}

		user, err := u.service.LoginUser(c, params.Email, params.Password)
		if err != nil {
			if err == userError.ErrorNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"message": "Invalid email or password",
				})
				return
			}
			errRes := userError.NewErrorInternal(err, "something went wrong")
			c.JSON(http.StatusInternalServerError, errRes)
			return
		}

		token, err := u.authService.GenerateUserToken(user.ID, user.Role)
		if err != nil {
			errRes := userError.NewErrorForbidden(err, "unable to generate token")
			c.JSON(http.StatusUnauthorized, errRes)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"user":  user,
			"token": token,
		})
	}
}

func (u *userController) GetUserProfile() gin.HandlerFunc {
	return func(c *gin.Context) {

		userID, _, err := getUserIdAndRoleFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		user, err := u.service.GetUserByID(c, userID)
		if err != nil {
			errRes := userError.NewErrorInternal(err, "something went wrong")
			c.JSON(http.StatusInternalServerError, errRes)
			return
		}

		c.JSON(http.StatusOK, user)

	}
}

type userRoleInput struct {
	Role string `json:"role" binding:"required,oneof=ADMIN USER"`
}

func (u *userController) UpdateUserRole() gin.HandlerFunc {
	return func(c *gin.Context) {

		userIdFromToken, _, err := getUserIdAndRoleFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		params := &userRoleInput{}
		err = c.ShouldBind(params)
		if err != nil {
			errRes := userError.NewErrorBadRequest(err, "invalid input")
			c.JSON(http.StatusBadRequest, errRes)
			return
		}

		userIdParam, ok := c.Params.Get("user_id")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid input",
			})
			return
		}

		userId, err := strconv.ParseUint(userIdParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid input",
				"error":   err.Error(),
			})
			return
		}

		if userIdFromToken == uint(userId) {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "cannot update yourself",
			})
			return
		}

		err = u.service.UpdateUserRole(c, uint(userId), params.Role)
		if err != nil {
			if err == userError.ErrorNotFound {
				errRes := userError.NewErrorNotFound(err, "user not found")
				c.JSON(http.StatusNotFound, errRes)
				return
			}
			errRes := userError.NewErrorInternal(err, "something went wrong")
			c.JSON(http.StatusNotFound, errRes)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "User update successful",
		})

	}
}
