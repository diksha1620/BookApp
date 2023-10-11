package routes

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/books/controllers"
	"github.com/books/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/swaggo/swag/example/celler/controller"
)

func PublicEndPoints(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	r.POST("/register", controllers.Register)
	r.POST("/login", authMiddleware.LoginHandler)

	// RefreshHandler can be used to refresh a token. The token still needs to be valid on refresh.
	// Shall be put under an endpoint that is using the GinJWTMiddleware.
	// Reply will be of the form {"token": "TOKEN"}.S
	r.GET("/refresh", authMiddleware.RefreshHandler)
}

func AuthenticatedEndpoints(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	r.POST("/blog/create", controllers.CreateBlog)
}

func GetRouter(router chan *gin.Engine) {
	gin.ForceConsoleColor()
	r := gin.Default()

	r.Use(cors.Default())

	authMiddleware, _ := middlewares.GetAuthMiddleware()

	// Create a BASE_URL - /api/v1
	v1 := r.Group("/api/v1/")
	PublicEndPoints(v1, authMiddleware)
	// AuthenticatedEndpoints(v1.Group("auth"), authMiddleware)
	router <- r
}
