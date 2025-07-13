package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "abc-user-backend/models"
    "abc-user-backend/services"
)

type UserHandler struct {
    service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
    return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
    search := c.Query("search")

    pageStr := c.DefaultQuery("page", "1")
    limitStr := c.DefaultQuery("limit", "10")

    page, err := strconv.Atoi(pageStr)
    if err != nil || page < 1 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
        return
    }
    limit, err := strconv.Atoi(limitStr)
    if err != nil || limit < 1 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
        return
    }

    offset := (page - 1) * limit

    total, err := h.service.Count(search)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count users"})
        return
    }

    users, err := h.service.GetAll(search, limit, offset)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "data":  users,
        "page":  page,
        "limit": limit,
        "total": total,
    })
}


func (h *UserHandler) GetUserByID(c *gin.Context) {
    id := c.Param("id")

    user, err := h.service.GetByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.Create(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, input)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
    id := c.Param("id")

    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.service.Update(id, &input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := h.service.GetByID(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
    id := c.Param("id")

    if err := h.service.Delete(id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusNoContent, nil)
}

