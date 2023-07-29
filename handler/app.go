package handler

import (
	"os"

	"github.com/fydhfzh/letter-notification/db"
	"github.com/fydhfzh/letter-notification/middleware"
	"github.com/fydhfzh/letter-notification/repository/letter_repository/letter_pg"
	"github.com/fydhfzh/letter-notification/repository/subdit_repository/subdit_pg"
	"github.com/fydhfzh/letter-notification/repository/user_letter_repository/user_letter_pg"
	"github.com/fydhfzh/letter-notification/repository/user_repository/user_pg"
	"github.com/fydhfzh/letter-notification/service"
	"github.com/gin-gonic/gin"
)

// Dont forget to implement CORS, Authentication, Authorization middleware
func App() {
	port := os.Getenv("PORT")

	db.InitializeDB()
	db := db.GetDBInstance()

	//User endpoint
	userRepository := user_pg.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := NewUserHandler(userService)

	//User letter endpoint
	userLetterRepository := user_letter_pg.NewUserLetterRepository(db)

	//Letter endpoint
	letterRepository := letter_pg.NewLetterRepository(db)
	letterService := service.NewLetterService(letterRepository, userRepository, userLetterRepository)
	letterHandler := NewLetterHandler(letterService)

	//Subdit endpoint
	subditRepository := subdit_pg.NewSubditRepository(db)
	subditService := service.NewSubditService(subditRepository)
	subditHandler := NewSubditHandler(subditService)

	r := gin.Default()
	api := r.Group("/api/v1")
	api.Use(middleware.CORSMiddleware())
	api.OPTIONS("/*any", middleware.CORSMiddleware())

	userRoute := api.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)
		userRoute.POST("/login", userHandler.Login)
		userRoute.POST("/logout", middleware.Authentication(), userHandler.Logout)
		userRoute.PATCH("/reset-password", middleware.Authentication(), userHandler.ResetPassword)
		userRoute.GET("/:userID", middleware.Authentication(), userHandler.GetUserByID)
		userRoute.GET("/", middleware.Authentication(), middleware.Authorization(), userHandler.GetUsersBySubditID)
	}

	letterRoute := api.Group("/letters")
	{
		letterRoute.POST("/", middleware.Authentication(), middleware.Authorization(), letterHandler.CreateLetter)
		letterRoute.GET("/:letterID", middleware.Authentication(), letterHandler.GetLetterByID)
		letterRoute.GET("/", middleware.Authentication(), letterHandler.GetLettersByToSubditID)
		letterRoute.PATCH("/:letterID", middleware.Authentication(), letterHandler.ArchiveLetter)
		letterRoute.DELETE("/:letterID", middleware.Authentication(), letterHandler.DeleteLetterByID)
		letterRoute.PATCH("/:letterID/archive", middleware.Authentication(), letterHandler.ArchiveLetter)
	}

	subditRoute := api.Group("/subdits")
	{
		subditRoute.POST("/", middleware.Authentication(), middleware.Authorization(), subditHandler.CreateSubdit)
		subditRoute.GET("/:subditID", middleware.Authentication(), subditHandler.GetSubditByID)
		subditRoute.GET("/", subditHandler.GetAllSubdit)
		subditRoute.DELETE("/:subditID", middleware.Authentication(), middleware.Authorization(), subditHandler.DeleteSubditByID)
		subditRoute.PATCH("/:subditID", middleware.Authentication(), middleware.Authorization(), subditHandler.UpdateSubditByID)
	}

	r.Run(":" + port)
}
