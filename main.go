package main

import (
	"go-api-mahasiswa/handler"
	"go-api-mahasiswa/mahasiswa"
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
	db.AutoMigrate(&mahasiswa.Mahasiswa{})

	mahasiswaRepository := mahasiswa.NewRepository(db)
	mahasiswaService := mahasiswa.NewService(mahasiswaRepository)
	mahasiswaHandler := handler.NewMahasiswaHandler(mahasiswaService)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")
	api.POST("/add_mahasiswa", mahasiswaHandler.AddMahasiswa)

	router.Run(":9999")
}
