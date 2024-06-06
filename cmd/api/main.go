package main

import (
	"fmt"
	"go_api/internal/config"
	"go_api/internal/handler"
	"go_api/internal/middleware"
	"go_api/internal/repository"
	"go_api/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
    cfg := config.LoadConfig()
    fmt.Println("🚀 ~ funcmain ~ cfg:", cfg.DBHost)

    dsn := "host=" + cfg.DBHost + " user=" + cfg.DBUsername + " password=" + cfg.DBPassword + " dbname=" + cfg.DBName + " port=" + cfg.DBPort + " sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }

    // Initialize repositories
    orgRepo := repository.NewOrganizationRepository(db)
    userRepo := repository.NewUserRepository(db)

    // Initialize services
    orgService := service.NewOrganizationService(orgRepo)
    userService := service.NewUserService(userRepo)
    authService := service.NewAuthService(cfg.JWTSecret)

    // Initialize handlers
    orgHandler := handler.NewOrganizationHandler(orgService)
    userHandler := handler.NewUserHandler(userService)

    r := gin.Default()

    // Public routes
    r.POST("/api/login", func(c *gin.Context) {
        var req struct {
            Email    string `json:"email"`
            Password string `json:"password"`
        }
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // Validate user credentials (dummy validation for example purposes)
        if req.Email == "test@example.com" && req.Password == "password" {
            token, err := authService.GenerateToken(req.Email)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }
            c.JSON(http.StatusOK, gin.H{"token": token})
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        }
    })

    // Protected routes
    api := r.Group("/api", middleware.JWTAuth())
    {
        // Organization routes
        api.POST("/organizations", orgHandler.CreateOrganization)
        api.GET("/organizations", orgHandler.GetOrganizations)
        api.GET("/organizations/:id", orgHandler.GetOrganizationByID)
        api.PUT("/organizations/:id", orgHandler.UpdateOrganization)
        api.DELETE("/organizations/:id", orgHandler.DeleteOrganization)

        // User routes
        api.POST("/users", userHandler.CreateUser)
        api.GET("/users", userHandler.GetUsers)
        api.GET("/users/:id", userHandler.GetUserByID)
        api.PUT("/users/:id", userHandler.UpdateUser)
        api.DELETE("/users/:id", userHandler.DeleteUser)

    }

    r.Run()
}
