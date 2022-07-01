package router

import (
	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go_api_boilerplate/config"
	"go_api_boilerplate/di"
	"time"
)

type Router struct {
	API *gin.Engine
}

func NewRouter() *Router {
	r := gin.Default()
	con := cors.DefaultConfig()
	con.AllowAllOrigins = true
	con.AllowCredentials = true
	con.AddAllowHeaders("authorization")
	r.Use(cors.New(con))
	r.Use(easyAuthMiddleware())

	logger, _ := zap.NewDevelopment()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	router := r.Group("/api/v1")
	{
		hd := di.InitUserApp()
		router.GET("/users", hd.Users())
		router.POST("/users", hd.CreateUser())
	}

	return &Router{
		API: r,
	}
}

func easyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		con := config.NewConfig()
		ah := c.Request.Header.Get("Authorization")
		if ah != con.Token {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("user", "1")
		c.Next()
	}
}
