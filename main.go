package main

import (
	"go-api-mahasiswa/handler"
	"go-api-mahasiswa/mahasiswa"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "cvhqcltochehwv:ab23f153eb3b8ca6b21bc2b5534fe2e43233217081eb78303eb95aa1931153b2@tcp(ec2-3-231-254-204.compute-1.amazonaws.com:5432)/d5iv6jkaephs52?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

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

	router.Run()
}
