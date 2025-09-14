// cmd/server/main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"user-crud/src/api/validation"
	"user-crud/src/domain/model"
	"user-crud/src/infra/db"

	_ "user-crud/docs"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/crypto/bcrypt"
)

// HealthResponse for health check
type HealthResponse struct {
	Status   string `json:"status" example:"OK"`
	Message  string `json:"message" example:"Server is running"`
	Database string `json:"database" example:"car_sale_db"`
}

// @title User Management API
// @version 1.0
// @description A complete user management system with authentication and validation
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

var jwtSecret = []byte("supersecretkey")

// RegisterRequest payload
type RegisterRequest struct {
	Username     string `json:"username" example:"johndoe" binding:"required,min=5,max=20"`
	Email        string `json:"email" example:"johndoe@example.com" binding:"required,email"`
	Password     string `json:"password" example:"Secret123!" binding:"required,password,min=8"`
	FirstName    string `json:"firstName" example:"John" binding:"required,min=2,max=50"`
	LastName     string `json:"lastName" example:"Doe" binding:"required,min=2,max=50"`
	MobileNumber string `json:"mobileNumber" example:"09123456789" binding:"required,mobile"`
}

// LoginRequest payload
type LoginRequest struct {
	Username string `json:"username" example:"johndoe" binding:"required,min=5,max=20"`
	Password string `json:"password" example:"Secret123!" binding:"required,min=8"`
}

// JWTClaims payload
type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// JWTMiddleware middleware
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if !strings.HasPrefix(tokenStr, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}
		tokenStr = strings.TrimPrefix(tokenStr, "Bearer ") // prefix رو بردار
		token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// RegisterHandler handles user registration
// @Summary Register a new user
// @Description Create a new user account with validation
// @Tags users
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "User registration data"
// @Success 201 {object} model.User
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /register [post]
func RegisterHandler(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": err.Error(),
		})
		return
	}

	// Check if username exists
	var existingUser model.User
	if result := db.DB.Where("username = ?", req.Username).First(&existingUser); result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Check if email exists
	if result := db.DB.Where("email = ?", req.Email).First(&existingUser); result.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user
	user := model.User{
		Username:     req.Username,
		Email:        req.Email,
		Password:     string(hashedPassword),
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		MobileNumber: req.MobileNumber,
		Enabled:      true,
	}

	// Save to database
	result := db.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// LoginHandler handles user login
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Login credentials"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /login [post]
func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Validation failed",
			"details": err.Error(),
		})
		return
	}

	// Find user
	var user model.User
	result := db.DB.Where("username = ?", req.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	claims := JWTClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenStr,
		"user": gin.H{
			"id":           user.ID,
			"username":     user.Username,
			"email":        user.Email,
			"firstName":    user.FirstName,
			"lastName":     user.LastName,
			"mobileNumber": user.MobileNumber,
		},
	})
}

// GetAllUsersHandler retrieves all users
// @Summary Get all users
// @Description Retrieve list of all registered users
// @Tags users
// @Produce json
// @Success 200 {array} model.User
// @Failure 500 {object} map[string]interface{}
// @Security BearerAuth
// @Router /users [get]
func GetAllUsersHandler(c *gin.Context) {
	var users []model.User
	result := db.DB.Find(&users)
	
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByIDHandler retrieves a user by ID
// @Summary Get user by ID
// @Description Retrieve a single user by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 404 {object} map[string]interface{}
// @Security BearerAuth
// @Router /users/{id} [get]
func GetUserByIDHandler(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	result := db.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// HealthCheckHandler health check endpoint
// @Summary Health check
// @Description Check if server is running
// @Tags system
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func HealthCheckHandler(c *gin.Context) {
	// به جای gin.H، از struct استفاده کن (اختیاری، اما بهتره)
	response := HealthResponse{ // <-- اینجا response رو با نام response تعریف کردم (نه goresponse – اشتباه تایپی بود!)
		Status:   "OK",
		Message:  "Server is running",
		Database: "car_sale_db",
	}
	c.JSON(http.StatusOK, response) // <-- و اینجا response رو پاس می‌دی
}

func main() {
	// Initialize database
	if err := db.InitDB(); err != nil {
		log.Fatal("❌ Failed to init DB:", err)
	}

	// Migrate existing data
	if err := db.MigrateExistingData(); err != nil {
		log.Fatal("❌ Failed to migrate existing data:", err)
	}

	// Setup Gin
	r := gin.Default()

	// Register custom validators
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validation.RegisterCustomValidators(v)
	}

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Public routes
	r.GET("/health", HealthCheckHandler)
	r.POST("/register", RegisterHandler)
	r.POST("/login", LoginHandler)

	// Protected routes group
	auth := r.Group("/", JWTMiddleware())
	auth.GET("/users", GetAllUsersHandler)
	auth.GET("/users/:id", GetUserByIDHandler)

	// Start server
	fmt.Println("🚀 Server starting on :8080")
	fmt.Println("📊 Database: car_sale_db")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
