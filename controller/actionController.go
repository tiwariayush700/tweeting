package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tiwariayush700/tweeting/auth"
	"github.com/tiwariayush700/tweeting/constants"
	userError "github.com/tiwariayush700/tweeting/error"
	"github.com/tiwariayush700/tweeting/models"
	"github.com/tiwariayush700/tweeting/services"
	"net/http"
)

type actionController struct {
	service                  services.ActionService
	userService              services.UserService
	approvalServiceProviders services.ApprovalServiceProviders
	app                      *app
	authService              auth.Service
}

func NewActionController(service services.ActionService, userService services.UserService, approvalServiceProviders services.ApprovalServiceProviders, authService auth.Service, app *app) *actionController {
	return &actionController{service: service, userService: userService, approvalServiceProviders: approvalServiceProviders, authService: authService, app: app}
}

func (a *actionController) RegisterRoutes() {
	router := a.app.Router
	superAdminRouterGroup := router.Group("/super-admin")
	{
		superAdminRouterGroup.Use(VerifyUserAndServe(a.authService))
		superAdminRouterGroup.Use(VerifySuperAdminAndServe())
		adminUserRouterGroup := superAdminRouterGroup.Group("/actions", a.GetActions())
		{
			//list all actions
			adminUserRouterGroup.GET("")
			adminUserRouterGroup.PUT("/approve", a.ApproveAction())
		}
	}
}

func (a *actionController) GetActions() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, _, err := getUserIdAndRoleFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		actions, err := a.service.FetchActions(c)
		if err != nil {
			if err == userError.ErrorNotFound {
				c.JSON(http.StatusNotFound, userError.NewErrorNotFound(err, "action not found"))
				return
			}
			c.JSON(http.StatusInternalServerError, userError.NewErrorInternal(err, "something went wrong"))
			return
		}

		c.JSON(http.StatusOK, actions)
	}
}

func (a *actionController) ApproveAction() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, _, err := getUserIdAndRoleFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		params := &models.ActionRequest{}
		err = c.ShouldBind(params)
		if err != nil {
			c.JSON(http.StatusBadRequest, userError.NewErrorBadRequest(err, "invalid input"))
			return
		}

		err = a.approvalServiceProviders.Update(c, params.Provider, params.ActionID, constants.ActionStatusApproved)
		if err != nil {
			c.JSON(http.StatusInternalServerError, userError.NewErrorInternal(err, "something went wrong"))
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":       "Action approved successfully",
			"action_status": constants.ActionStatusApproved,
		})
	}
}
