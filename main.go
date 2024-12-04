package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GenerateRequest struct {
	Length    int  `json:"length"`
	Numbers   bool `json:"numbers"`
	Symbols   bool `json:"symbols"`
	Lowercase bool `json:"lowercase"`
	Uppercase bool `json:"uppercase"`
}

func generatePassword(length int, numbers, symbols, lowercase, uppercase bool) string {
	var charset string
	if numbers {
		charset += "0123456789"
	}
	if symbols {
		charset += "!@#$%^&*()_+-=[]{}|;:,.<>?"
	}
	if lowercase {
		charset += "abcdefghijklmnopqrstuvwxyz"
	}
	if uppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if charset == "" {
		charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}

	rand.Seed(time.Now().UnixNano())
	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}
	return string(password)
}

func main() {
	r := gin.Default()

	// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	r.Use(cors.New(config))

	r.POST("/generate", func(c *gin.Context) {
		var req GenerateRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if req.Length <= 0 {
			req.Length = 12 // Default length
		}

		password := generatePassword(req.Length, req.Numbers, req.Symbols, req.Lowercase, req.Uppercase)
		c.JSON(http.StatusOK, gin.H{
			"password": password,
		})
	})

	r.Run(":8080")
} 