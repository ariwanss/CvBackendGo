package controller

import (
	"net/http"

	"github.com/ariwanss/CvBackendGo/entity"
	"github.com/ariwanss/CvBackendGo/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCvItem(c *gin.Context) {
	var newCvItem entity.CvItem
	err := c.ShouldBind(&newCvItem)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userId := c.Value("userId").(primitive.ObjectID)
	newCvItem.UserID = userId
	inserted, err := service.InsertCvItem(&newCvItem)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, inserted)
}

func GetCvItems(c *gin.Context) {
	userId := c.Value("userId").(primitive.ObjectID)
	cvItems, err := service.GetCvItemsByUserId(userId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, cvItems)
}

func UpdateCvItem(c *gin.Context) {
	var update entity.CvItem
	err := c.ShouldBind(&update)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	userId := c.Value("userId").(primitive.ObjectID)
	cvItemId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	updatedCvItem, err := service.UpdateCvItem(cvItemId, userId, &update)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, updatedCvItem)
}

func DeleteCvItem(c *gin.Context) {
	cvItemId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	deletedCount, err := service.DeleteCvItem(cvItemId)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": deletedCount})
}

func GetAllCvItems(c *gin.Context) {
	cvItems, err := service.GetAllCvItems()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, cvItems)
}

// func DropCvItemsCol(c *gin.Context) {
// 	err := service.DropCvItemCol()
// 	if err != nil {
// 		c.AbortWithError(http.StatusInternalServerError, err)
// 		return
// 	}
// 	c.Status(http.StatusOK)
// }
