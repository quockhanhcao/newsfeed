package http

import (
	"fmt"
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

	userRouter := router.Group("/user")
	userRouter.POST("/signup", server.Signup)
	userRouter.POST("/login", server.Login)

	server.router = router
	return server, nil
}

func (s *HttpServer) Start(addr string) error {
	return s.router.Run(addr)
}

type SignupRequest struct {
	Username    string `json:"user_name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
	Dob         string `json:"dob"` // Date of birth
}

type UserData struct {
	Username    string `json:"user_name"`
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Dob         string `json:"dob"`        // Date of birth
	CreatedAt   string `json:"created_at"` // Creation timestamp
	UpdatedAt   string `json:"updated_at"` // Last update timestamp
}

func (s *HttpServer) Signup(c *gin.Context) {
	// Handle user signup logic here
	req := &SignupRequest{}
	if err := c.ShouldBind(req); err != nil {
		fmt.Println("error parsing signup request:", err)
		c.JSON(http.StatusBadRequest, &CommonResponse{
			Code:    BadRequest,
			Message: "Invalid request data", // should be customized based on error
			Data:    nil})
		return
	}

	c.JSON(http.StatusOK, &CommonResponse{
		Code:    Success,
		Message: "User signed up successfully",
		Data: &UserData{
			Username:    req.Username,
			DisplayName: req.DisplayName,
			Email:       req.Email,
			Dob:         req.Dob,
			CreatedAt:   "2023-10-01T12:00:00Z", // Example timestamp
			UpdatedAt:   "2023-10-01T12:00:00Z", // Example timestamp
		}})
}

func (s *HttpServer) Login(c *gin.Context) {
	// Handle user login logic here
	c.JSON(200, gin.H{
		"message": "User logged in successfully",
	})
}
