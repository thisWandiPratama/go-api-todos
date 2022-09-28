package main

import (
	"go-api-mahasiswa/auth"
	"go-api-mahasiswa/handler"
	"go-api-mahasiswa/helper"
	"go-api-mahasiswa/todos"
	"go-api-mahasiswa/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "gurujagoan:Gurujagoan@123@tcp(localhost:3306)/todosku?parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&todos.Todo{}, &user.User{})

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	mahasiswaRepository := todos.NewRepository(db)
	mahasiswaService := todos.NewService(mahasiswaRepository)
	mahasiswaHandler := handler.NewTodoHandler(mahasiswaService)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/images/", "./assets")

	//User
	router.POST("/api/v1/users/register/", userHandler.RegisterUser)
	router.POST("/api/v1/users/login/", userHandler.Login)
	router.POST("/api/v1/users/email_checkers/", userHandler.CheckEmailAvailability)
	router.POST("/api/v1/users/avatars/", authMiddleware(authService, userService), userHandler.UploadAvatar)
	router.GET("/api/v1/users/users_fetch/", authMiddleware(authService, userService), userHandler.FetchUser)

	// Todo
	router.POST("/api/v1/add_todo/", mahasiswaHandler.AddTodos)
	router.POST("/api/v1/all_todo/", mahasiswaHandler.FindAllTodo)
	router.POST("/api/v1/delete_todo/:id/", mahasiswaHandler.DeleteMahasiswa)
	router.POST("/api/v1/update_todo/", mahasiswaHandler.UpdateTodo)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
