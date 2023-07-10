package main

import (
	"jejetickets/initializers"
	"jejetickets/models"

	"fmt"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could Not Load Environment Variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Event{})
	fmt.Println("---Migration Complete---")
}
