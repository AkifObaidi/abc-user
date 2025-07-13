package routes

import (
    "github.com/gin-gonic/gin"
    "abc-user-backend/handlers"
    "abc-user-backend/services"
)

func RegisterUserRoutes(r *gin.Engine, userService *services.UserService) {
    userHandler := handlers.NewUserHandler(userService)

    users := r.Group("/users")
    {
        users.GET("", userHandler.GetUsers)          
        users.GET("/:id", userHandler.GetUserByID)   
        users.POST("", userHandler.CreateUser)       
        users.PUT("/:id", userHandler.UpdateUser)    
        users.DELETE("/:id", userHandler.DeleteUser) 
    }
}

