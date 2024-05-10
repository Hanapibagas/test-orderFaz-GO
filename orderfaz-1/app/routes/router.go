package routes

import (
	"orderfaz-1/app/middlewares"
	"orderfaz-1/features/user/data"
	"orderfaz-1/features/user/handler"
	"orderfaz-1/features/user/service"
	"orderfaz-1/utils/encrypts"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	hashService := encrypts.NewHashService()

	userData := data.NewUser(db)
	userServide := service.NewUser(userData, hashService)
	userHandler := handler.NewUser(userServide)

	e.POST("/register", userHandler.RegisterUser)
	e.POST("/login", userHandler.LoginUser)
	e.GET("/user", userHandler.GetById, middlewares.JWTMiddleware())
}
