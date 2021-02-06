package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tiwariayush700/tweeting/auth"
	"github.com/tiwariayush700/tweeting/constants"
	userError "github.com/tiwariayush700/tweeting/error"
	"github.com/tiwariayush700/tweeting/models"
	"github.com/tiwariayush700/tweeting/services"
	"gorm.io/datatypes"
	"net/http"
	"strconv"
)

type tweetController struct {
	service       services.TweetService
	userService   services.UserService
	actionService services.ActionService
	app           *app
	authService   auth.Service
}

func NewTweetController(service services.TweetService, userService services.UserService, actionService services.ActionService, authService auth.Service, app *app) *tweetController {
	return &tweetController{
		service:       service,
		userService:   userService,
		actionService: actionService,
		app:           app,
		authService:   authService,
	}
}

func (t *tweetController) RegisterRoutes() {
	router := t.app.Router

	tweetRouterGroup := router.Group("/tweets")
	{
		tweetRouterGroup.Use(VerifyUserAndServe(t.authService))
		tweetRouterGroup.POST("", t.CreateTweets())
		tweetRouterGroup.GET("", t.GetTweetsByUserID())

	}

	adminTweetRouterGroup := router.Group("/admin/tweets")
	{
		adminTweetRouterGroup.Use(VerifyUserAndServe(t.authService))
		adminTweetRouterGroup.Use(VerifyAdminAndServe(t.authService))
		adminTweetRouterGroup.GET("", t.GetTweets())
		adminTweetRouterGroup.PUT("/:tweet_id", t.UpdateTweet())

	}
}

func (t *tweetController) GetTweetsByUserID() gin.HandlerFunc {
	return func(c *gin.Context) {

		userID, _, err := getUserIdAndRoleFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		tweets, err := t.service.GetTweetsByUserID(c, userID)
		if err != nil {
			if err == userError.ErrorNotFound {
				c.JSON(http.StatusNotFound, userError.NewErrorNotFound(err, "tweets not found"))
				return
			}
			c.JSON(http.StatusInternalServerError, userError.NewErrorInternal(err, "something went wrong"))
			return
		}

		c.JSON(http.StatusOK, tweets)
	}
}

func (t *tweetController) CreateTweets() gin.HandlerFunc {
	return func(c *gin.Context) {

		userID, _, err := getUserIdAndRoleFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		params := &models.TweetRequest{}
		err = c.ShouldBind(params)
		if err != nil {
			errRes := userError.NewErrorBadRequest(err, "invalid input")
			c.JSON(http.StatusBadRequest, errRes)
			return
		}

		tweet := &models.Tweet{
			Message: params.Message,
			UserID:  userID,
		}

		err = t.service.CreateTweet(c, tweet)
		if err != nil {
			c.JSON(http.StatusInternalServerError, userError.NewErrorInternal(err, "something went wrong"))
			return
		}

		c.JSON(http.StatusOK, tweet)
	}
}

func (t *tweetController) UpdateTweet() gin.HandlerFunc {
	return func(c *gin.Context) {

		userIdFromToken, _, err := getUserIdAndRoleFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		params := &models.TweetRequest{}
		err = c.ShouldBind(params)
		if err != nil {
			errRes := userError.NewErrorBadRequest(err, "invalid input")
			c.JSON(http.StatusBadRequest, errRes)
			return
		}

		tweetIDFromParam, ok := c.Params.Get("tweet_id")
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid input",
			})
			return
		}

		tweetID, err := strconv.ParseUint(tweetIDFromParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid input",
				"error":   err.Error(),
			})
			return
		}

		tweet := &models.Tweet{
			Message: params.Message,
		}
		tweet.ID = uint(tweetID)
		tweetBytes, _ := json.Marshal(&tweet)
		action := &models.Action{
			Message: fmt.Sprintf("Approval for updating tweet message : %s for user ID : %d pending for approval", tweet.Message, userIdFromToken),
			Status:  constants.ActionStatusPending,
			Body:    datatypes.JSON(tweetBytes),
		}

		err = t.actionService.CreateAction(c, action)
		if err != nil {
			c.JSON(http.StatusNotFound, userError.NewErrorInternal(err, "something went wrong"))
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message":   fmt.Sprintf("Approval for updating tweet message : %s for user ID : %d pending for approval", tweet.Message, userIdFromToken),
			"action_id": action.ID,
			"status":    action.Status,
		})

	}
}

func (t *tweetController) GetTweets() gin.HandlerFunc {
	return func(c *gin.Context) {

		_, _, err := getUserIdAndRoleFromContext(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		tweets, err := t.service.FetchTweets(c)
		if err != nil {
			if err == userError.ErrorNotFound {
				c.JSON(http.StatusNotFound, userError.NewErrorNotFound(err, "tweets not found"))
				return
			}
			c.JSON(http.StatusInternalServerError, userError.NewErrorInternal(err, "something went wrong"))
			return
		}

		c.JSON(http.StatusOK, tweets)
	}
}