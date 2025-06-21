package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	router *gin.Engine
}

func New() (*HttpServer, error) {
	router := gin.Default()
	server := &HttpServer{}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// userRouter := router.Group("/user")
	// userRouter.POST("/signup", server.Signup)
	// userRouter.POST("/login", server.Login)

	server.router = router
	return server, nil
}

func (s *HttpServer) Start(addr string) error {
	return s.router.Run(addr)
}
