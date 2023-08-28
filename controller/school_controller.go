package controller

import (
	"berkeley/model"
	"berkeley/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllSchools(c *gin.Context) {
	result := service.GetAllSchools()
	c.JSON(http.StatusOK, result)
}

func GetSchoolByID(c *gin.Context) {
	result := service.GetSchoolByID(c.Param("schoolID"))
	if result.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "No school found with given id: " + c.Param("schoolID")})
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func CreateSchool(c *gin.Context) {
	var input model.School
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	input.ID = c.Param("schoolID")
	if err := service.CreateSchool(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, service.GetSchoolByID(input.ID))
}

func VerifySchool(c *gin.Context) {
	if err := service.VerifySchool(c.Param("schoolID")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	result := service.GetSchoolByID(c.Param("schoolID"))
	if result.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "No school found with given id: " + c.Param("schoolID")})
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func UnverifySchool(c *gin.Context) {
	if err := service.UnverifySchool(c.Param("schoolID")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	result := service.GetSchoolByID(c.Param("schoolID"))
	if result.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "No school found with given id: " + c.Param("schoolID")})
	} else {
		c.JSON(http.StatusOK, result)
	}
}
