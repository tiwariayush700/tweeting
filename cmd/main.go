package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tiwariayush700/tweeting/config"
	"github.com/tiwariayush700/tweeting/controller"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	appConfig := config.GetAppConfiguration()

	router := gin.Default()

	db, err := gorm.Open(postgres.Open(postgresUri(appConfig.PostgresServer, appConfig.PostgresUser, appConfig.PostgresPassword, appConfig.PostgresPort, appConfig.DbName)),
		&gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatal("unable to connect db ", err)
	}

	app := controller.NewApp(appConfig, db, router)

	err = app.Migrate()
	if err != nil {
		logrus.Fatalf("err migrating db with err : %v", err)
	}

	app.Start()

}

func postgresUri(host, user string, password string, port string, dbname string) string {

	connectionUri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	return connectionUri
}
