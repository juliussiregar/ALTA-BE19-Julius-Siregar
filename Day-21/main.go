package main

import (
	"Day-21/user"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	dsn := "root:admin123@tcp(localhost:3306)/tugasday21?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("cannot connect to DB", err.Error())
	}

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(*userRepository)
	userHandler := user.NewUserHandler(*userService)

	// Routing
	e.POST("/users", userHandler.CreateUserHandler)
	e.GET("/users", userHandler.GetAllUsersHandler)
	e.GET("/users/:id", userHandler.GetUserByIDHandler)
	e.PUT("/users/:id", userHandler.UpdateUserHandler)
	e.DELETE("/users/:id", userHandler.DeleteUserHandler)

	// Start the server
	e.Logger.Fatal(e.Start(":8000"))
}
