package router

import (
    "net/http"
    "strconv"
    "user-crud/middleware"
    "user-crud/usecase"

    "github.com/gin-gonic/gin"
)

func RegisterRoleRoutes(r *gin.Engine) {
    roleGroup := r.Group("/roles")
    roleGroup.Use(middleware.JWTMiddleware())

    roleGroup.POST("/", func(c *gin.Context) {
        var req struct{ Name string `json:"name"` }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        role, err := usecase.CreateRole(req.Name)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, role)
    })

    roleGroup.POST("/assign", func(c *gin.Context) {
        var req struct {
            UserID uint `json:"user_id"`
            RoleID uint `json:"role_id"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := usecase.AssignRoleToUser(req.UserID, req.RoleID); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "role assigned"})
    })

    roleGroup.GET("/user/:id", func(c *gin.Context) {
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
            return
        }
        
        roles, err := usecase.GetUserRoles(uint(id))
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        
        // ایجاد یک ساختار ساده برای پاسخ بدون فیلدهای اضافی
        type simpleRole struct {
            ID   uint   `json:"id"`
            Name string `json:"name"`
        }
        
        var response []simpleRole
        for _, role := range roles {
            response = append(response, simpleRole{
                ID:   role.ID,
                Name: role.Name,
            })
        }
        
        c.JSON(http.StatusOK, response)
    })
}