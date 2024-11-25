package server

import (
	"AirBnbProject/controller"
	"AirBnbProject/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateRouter(handlers *controller.ControllerManager, mode string) *gin.Engine {
	var router *gin.Engine
	if mode == "test" {
		gin.SetMode(gin.ReleaseMode)
		router = gin.New()
	} else {
		router = gin.Default()
	}

	//router := gin.Default()
	//set a lower memory limit for multipart forms
	router.MaxMultipartMemory = 8 << 20 //8 Mib
	router.Static("/static", "./public")

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://google.com"}
	// config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	// config.AllowAllOrigins = true

	router.Use(cors.New(config))

	api := router.Group("/api")

	api.GET("/home", func(ctx *gin.Context) {
		ctx.String(200, "Hello Gin FB")
	})

	categoryRoute := api.Group("/category")
	{
		categoryRoute.Use(middleware.AuthMiddleware())
		categoryRoute.GET("/", handlers.GetListCategory)
		categoryRoute.POST("/", handlers.PostCategory)
		categoryRoute.GET("/:id", handlers.GetCategoryById)
		categoryRoute.PUT("/:id", handlers.UpdateCategory)

	}

	productRoute := api.Group("/product")
	{
		productRoute.Use(middleware.AuthMiddleware())
		productRoute.GET("", handlers.FindAllProduct)
		productRoute.GET("/", handlers.FindAllProduct)
		productRoute.POST("/", handlers.CreateProduct)
		productRoute.GET("/:id", handlers.FindProductById)
		productRoute.DELETE("/:id", handlers.DeleteProduct)
		productRoute.GET("/paging", handlers.FindAllProductPaging)
		productRoute.POST("/uploadProductImage", handlers.UploadMultipleProductImage)
		productRoute.PUT("/:id", handlers.UpdateProduct)
	}

	orderRoute := api.Group("/order")
	{
		productRoute.Use(middleware.AuthMiddleware())
		orderRoute.Use(middleware.AuthMiddleware())
		orderRoute.GET("/cart/:id", handlers.FindCartByCustomerId)
		orderRoute.POST("/cart/add", handlers.AddToCart)
		orderRoute.GET("/", handlers.FindAllOrder)
		orderRoute.GET("/:id", handlers.FindOrderById)
		orderRoute.POST("/", handlers.CreateOrder)
	}

	userRoute := api.Group("/user")
	{
		userRoute.POST("/signup", handlers.Signup)
		userRoute.POST("/signin", handlers.Sigin)
		userRoute.POST("/signout", handlers.Signout)
		userRoute.GET("/profile", handlers.GetProfile)
	}

	return router
}
