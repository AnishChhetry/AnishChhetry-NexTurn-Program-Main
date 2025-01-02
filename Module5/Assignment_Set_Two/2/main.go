package main

import (
	"ecommerce-inventory/config"
	"ecommerce-inventory/controller"
	"ecommerce-inventory/middleware"
	"ecommerce-inventory/repository"
	"ecommerce-inventory/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.InitializeDatabase()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	router := gin.Default()

	router.Use(middleware.LoggingMiddleware())

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	authorized := router.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/product", middleware.ValidationMiddleware(), productController.AddProduct)
		authorized.GET("/product/:id", productController.GetProduct)
		authorized.PUT("/product/:id", productController.UpdateProduct)
		authorized.DELETE("/product/:id", productController.DeleteProduct)
		authorized.GET("/products", productController.GetAllProducts)
	}

	router.Run(":8080")
}
