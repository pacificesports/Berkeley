package controller

import (
	"berkeley/model"
	"berkeley/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetEmailsForSchool(c *gin.Context) {
	result := service.GetEmailsForSchool(c.Param("schoolID"))
	c.JSON(http.StatusOK, result)
}

func AddEmailForSchool(c *gin.Context) {
	var input model.Email
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	input.SchoolID = c.Param("schoolID")
	if err := service.AddEmailForSchool(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service.GetEmailsForSchool(c.Param("schoolID")))
}

func RemoveEmailForSchool(c *gin.Context) {
	var input model.Email
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	input.SchoolID = c.Param("schoolID")
	if err := service.DeleteEmailForSchool(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service.GetEmailsForSchool(c.Param("schoolID")))
}
