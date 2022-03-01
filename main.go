package main

import (
	"go-api-koperasi/auth"
	"go-api-koperasi/handler"
	"go-api-koperasi/pengajuan"
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

	pengajuanRepository := pengajuan.NewRepository(db)
	pengajuanService := pengajuan.NewService(pengajuanRepository)
	pengajuanHandler := handler.NewPengajuanHandler(pengajuanService)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")
	api.POST("/login", userHandler.Login)
	api.GET("/pengajuan", pengajuanHandler.FindAll)
	api.GET("/pengajuan/:id", pengajuanHandler.FindByID)
	api.PUT("/pengajuan/jaminanbarang", pengajuanHandler.AddJaminan)
	api.POST("/pengajuan/jaminan/bukti", pengajuanHandler.AddBuktiJaminan)
	api.GET("/pengajuan/jaminan/statusdraf", pengajuanHandler.FindAllByStatusDraf)
	router.Run()
}
