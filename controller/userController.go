package controller

import (
	"net/http"

	"github.com/ariwanss/CvBackendGo/entity"
	"github.com/ariwanss/CvBackendGo/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginData struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var newUser entity.User

	err := c.ShouldBind(&newUser)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := service.CreateUser(&newUser)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Set("user", user)
	c.Next()
}

func Login(c *gin.Context) {
	var loginData LoginData
	err := c.ShouldBind(&loginData)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	user, err := service.LoginUser(loginData.Username, loginData.Password)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Set("user", user)
	c.Next()
}

func UpdateUser(c *gin.Context) {
	var update entity.User
	err := c.ShouldBind(&update)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userId := c.Value("userId").(primitive.ObjectID)
	updatedUser, err := service.UpdateUser(userId, &update)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

func DeleteUser(c *gin.Context) {
	userId := c.Value("userId").(primitive.ObjectID)
	deletedCount, err := service.DeleteUser(userId)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"deletedCount": deletedCount})
}

func GetAllUser(c *gin.Context) {
	users, err := service.GetAllUser()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, users)
}
