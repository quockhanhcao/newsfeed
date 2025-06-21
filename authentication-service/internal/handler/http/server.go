package http

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"net/http"
	"os"
	"time"
)

type HttpServer struct {
	router *gin.Engine
	DB     *sql.DB
}

var counts int

func New() (*HttpServer, error) {
	router := gin.Default()
	server := &HttpServer{}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	userRouter := router.Group("/user")
	userRouter.POST("/signup", server.Signup)
	userRouter.POST("/login", server.Login)
	server.router = router

	// connect to the database
	conn := connectToDB()
	if conn == nil {
        fmt.Println("Failed to connect to the database after multiple attempts")
        server.DB = nil
	} else {
        server.DB = conn
    }

	return server, nil
}

func (s *HttpServer) Start(addr string) error {
	return s.router.Run(addr)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return db, nil
}

func connectToDB() *sql.DB {
	connStr := os.Getenv("DSN")
    fmt.Println("Connecting to PostgreSQL with DSN:", connStr)
	if connStr == "" {
		return nil
	}

	for {
		connection, err := openDB(connStr)
		if err != nil {
            log.Println("Error connecting to PostgreSQL:", err)
			// log.Println("PostgreSQL not yet ready")
			counts++
		} else {
			log.Println("PostgreSQL connected successfully")
			return connection
		}
		if counts > 10 {
			log.Println(err)
			return nil
		}
		log.Println("Retrying in 2 seconds...")
		time.Sleep(2 * time.Second)
		continue
	}
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
