package main

import (
	"net/http"
	"user-crud/infra/db"
	"user-crud/model"
	"user-crud/router"
	"user-crud/usecase"

	"github.com/gin-gonic/gin"
)
	
func initDB() {
	err := db.InitDB()
	if err != nil {
		panic(err)
	}

	err = db.DB.AutoMigrate(&model.Role{}, &model.User{}, &model.UserRole{})
	if err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()
	initDB()

	// ثبت‌نام
	r.POST("/register", func(c *gin.Context) {
		var req struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := usecase.RegisterUser(req.Username, req.Email, req.Password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "user registered"})
	})

	// ورود
	r.POST("/login", func(c *gin.Context) {
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := usecase.LoginUser(req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "login successful",
			"token":   token,
		})
	})

	// Profile محافظت‌شده
	r.GET("/profile", func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
			return
		}

		claims, err := usecase.ValidateToken(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "welcome to your profile",
			"username": claims.Username,
		})
	})

	// مسیرهای Role و UserRole
	router.RegisterRoleRoutes(r)

	r.Run(":8080")
}
