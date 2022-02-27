package main

import (
	"go-api-koperasi/auth"
	"go-api-koperasi/handler"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "jennifer:jennifer38500@tcp(koperasi.crossnet.co.id:3306)/koperasi?parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := auth.NewRepository(db)
	userService := auth.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")
	api.POST("/login", userHandler.Login)
	router.Run()
}
