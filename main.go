package main

import (
	"fmt"
	"jejetickets/controllers"
	"jejetickets/initializers"
	"jejetickets/models"
	"jejetickets/routes"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	EventController      controllers.EventController
	EventRouteController routes.EventRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	EventController = controllers.NewEventController(initializers.DB)
	EventRouteController = routes.NewRouteEventController(EventController)

	server = gin.Default()
}

func migrateDatabase() {
	err := initializers.DB.AutoMigrate(&models.User{}, &models.Event{})
	if err != nil {
		log.Fatal("Migration Didnt Succeed", err)
	}

	fmt.Println("---Migration Done---")
}

func main() {

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	migrateDatabase()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/jejetickets", func(ctx *gin.Context) {
		message := "Welcome to JejeTickets"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	EventRouteController.EventRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
