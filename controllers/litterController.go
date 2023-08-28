package controllers

import (
	"github.com/bilguunerkh/unetei.mn/initializers"
	"github.com/bilguunerkh/unetei.mn/models"
	"github.com/gin-gonic/gin"
)

func LitterCreate(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	litter := models.Litter{Title: body.Title, Body: body.Body}

	hariu := initializers.DB.Create(&litter)

	if hariu.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"litters": litter,
	})
}

func GetLitters(c *gin.Context) {
	var litters []models.Litter
	initializers.DB.Find(&litters)

	c.JSON(200, gin.H{
		"litter": litters,
	})
}

func GetLitter(c *gin.Context) {
	id := c.Param("id")

	var litter models.Litter
	initializers.DB.First(&litter, id)

	c.JSON(200, gin.H{
		"litter": litter,
	})
}

func UpdateLitter(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	var litter models.Litter
	initializers.DB.First(&litter, id)

	initializers.DB.Model(&litter).Updates(models.Litter{
		Title: body.Title,
		Body:  body.Body,
	})

	c.JSON(200, gin.H{
		"litter": litter,
	})
}

func DeleteLitter(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Litter{}, id)

	c.JSON(200, gin.H{
		"message": "Successfully deleted",
	})
}
