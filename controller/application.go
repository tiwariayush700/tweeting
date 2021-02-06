package controller

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	authImpl "github.com/tiwariayush700/tweeting/auth/impl"
	"github.com/tiwariayush700/tweeting/config"
	"github.com/tiwariayush700/tweeting/models"
	repositoryImpl "github.com/tiwariayush700/tweeting/repository/impl"
	"github.com/tiwariayush700/tweeting/services"
	serviceImpl "github.com/tiwariayush700/tweeting/services/impl"
	"gorm.io/gorm"
)

// App structure for tenant microservice
type app struct {
	Config *config.Config
	DB     *gorm.DB //set from main.go
	Router *gin.Engine
}

func NewApp(config *config.Config, db *gorm.DB, router *gin.Engine) *app {
	return &app{
		Config: config,
		DB:     db,
		Router: router,
	}
}

func (app *app) Start() {

	app.Router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS", "HEAD", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	}))

	//repositories
	userRepositoryImpl := repositoryImpl.NewUserRepositoryImpl(app.DB)
	actionRepostiroyImpl := repositoryImpl.NewActionRepositoryImpl(app.DB)

	//services
	userService := serviceImpl.NewUserServiceImpl(userRepositoryImpl)
	actionService := serviceImpl.NewActionServiceImpl(actionRepostiroyImpl)
	authService := authImpl.NewAuthService(app.Config.AuthSecret)
	userApprovalService := serviceImpl.NewUserApprovalServiceImpl(userRepositoryImpl, actionRepostiroyImpl)
	approvalServiceProviders := make(map[string]services.ApprovalService)
	approvalServiceProviders["user"] = userApprovalService
	approvalService := serviceImpl.NewApprovalServiceImpl(approvalServiceProviders)

	//controllers
	userController := NewUserController(userService, actionService, authService, app)
	actionController := NewActionController(actionService, userService, approvalService, authService, app)

	//register routes
	userController.RegisterRoutes()
	actionController.RegisterRoutes()

	app.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	logrus.Info("Application loaded successfully ")
	logrus.Fatal(app.Router.Run(":" + app.Config.Port))

}

func (app *app) Migrate() error {
	if err := app.DB.AutoMigrate(&models.User{}); err != nil {
		return err
	}

	if err := app.DB.AutoMigrate(&models.Action{}); err != nil {
		return err
	}

	if err := app.DB.AutoMigrate(&models.Tweet{}); err != nil {
		return err
	}

	return nil
}
