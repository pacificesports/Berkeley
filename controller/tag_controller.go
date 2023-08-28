package controller

import (
	"berkeley/model"
	"berkeley/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTagsForSchool(c *gin.Context) {
	result := service.GetTagsForSchool(c.Param("schoolID"))
	c.JSON(http.StatusOK, result)
}

func AddTagForSchool(c *gin.Context) {
	var input model.Tag
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	input.SchoolID = c.Param("schoolID")
	if err := service.AddTagForSchool(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service.GetTagsForSchool(c.Param("schoolID")))
}

func RemoveTagForSchool(c *gin.Context) {
	var input model.Tag
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	input.SchoolID = c.Param("schoolID")
	if err := service.DeleteTagForSchool(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service.GetTagsForSchool(c.Param("schoolID")))
}
