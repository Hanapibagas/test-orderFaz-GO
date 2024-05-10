package routes

import (
	"orderfaz-1/app/middlewares"
	"orderfaz-1/features/logistic/data"
	"orderfaz-1/features/logistic/handler"
	"orderfaz-1/features/logistic/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	// hashService := encrypts.NewHashService()

	logisticData := data.NewLogistic(db)
	logisticServide := service.NewLogistic(logisticData)
	logisticHandler := handler.NewLogistic(logisticServide)

	e.POST("/logistic", logisticHandler.CreateLogistic, middlewares.JWTMiddleware())
	e.GET("/logistic", logisticHandler.SerachByQuery, middlewares.JWTMiddleware())
	// e.GET("/user", userHandler.GetById, middlewares.JWTMiddleware())
}
