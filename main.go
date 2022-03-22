package main

import (
	"fmt"
	"go-api-mahasiswa/handler"
	"go-api-mahasiswa/mahasiswa"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// dsn := "jennifer:jennifer38500@tcp(koperasi.crossnet.co.id:3306)/koperasi?parseTime=True"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbPort, dbName)
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
	api.GET("/all_mahasiswa", mahasiswaHandler.FindAllMahasiswa)
	api.POST("/delete_mahasiswa/:id", mahasiswaHandler.DeleteMahasiswa)

	router.Run(":9999")
}
