package controllers

import (
	"github.com/bilguunerkh/unetei.mn/initializers"
	"github.com/bilguunerkh/unetei.mn/models"
	"github.com/gin-gonic/gin"
)

func GiftCreate(c *gin.Context) {
	var body struct {
		Image string
		Body  string
	}
	c.Bind(&body)

	gift := models.Gift{Image: body.Image, Body: body.Body}

	result := initializers.DB.Create(&gift)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"gifts": gift,
	})
}

func GetGifts(c *gin.Context) {
	var gifts []models.Gift
	initializers.DB.Find(&gifts)

	c.JSON(200, gin.H{
		"gifts": gifts,
	})
}

func GetGift(c *gin.Context) {
	id := c.Param("id")

	var gift models.Gift
	initializers.DB.First(&gift, id)

	c.JSON(200, gin.H{
		"gift": gift,
	})
}

func UpdateGift(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Image string
		Body  string
	}
	c.Bind(&body)

	var gift models.Gift
	initializers.DB.First(&gift, id)

	initializers.DB.Model(&gift).Updates(models.Gift{
		Image: body.Image,
		Body:  body.Body,
	})
	c.JSON(200, gin.H{
		"gift": gift,
	})
}

func DeleteGift(c *gin.Context) {
	id := c.Param("id")
	initializers.DB.Delete(&models.Gift{}, id)

	c.JSON(200, gin.H{
		"message": "Successfully deleted",
	})
}
