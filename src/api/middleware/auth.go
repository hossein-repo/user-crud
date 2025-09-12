// api/middleware/auth.go
package middleware

import (
    "net/http"
    "strings"
    "user-crud/usecase"

    "github.com/gin-gonic/gin"
)

func Authentication(tokenUsecase *usecase.TokenUsecase) gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token required"})
            return
        }

        token := strings.TrimPrefix(authHeader, "Bearer ")
        claims, err := tokenUsecase.ValidateToken(token)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }

        c.Set("userId", claims.UserID)
        c.Set("username", claims.Username)
        c.Set("roles", claims.Roles)
        
        c.Next()
    }
}