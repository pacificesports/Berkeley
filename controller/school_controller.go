package controller

import (
	"berkeley/model"
	"berkeley/service"
	"berkeley/utils"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	oteltrace "go.opentelemetry.io/otel/trace"
	"net/http"
)

func GetAllSchools(c *gin.Context) {
	// Start tracing span
	span := utils.BuildSpan(c.Request.Context(), "GetAllSchools", oteltrace.WithAttributes(attribute.Key("Request-ID").String(c.GetHeader("Request-ID"))))
	defer span.End()

	result := service.GetAllSchools()
	c.JSON(http.StatusOK, result)
}

func GetSchoolByID(c *gin.Context) {
	// Start tracing span
	span := utils.BuildSpan(c.Request.Context(), "GetSchool", oteltrace.WithAttributes(attribute.Key("Request-ID").String(c.GetHeader("Request-ID"))))
	defer span.End()

	result := service.GetSchoolByID(c.Param("schoolID"))
	if result.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "No school found with given id: " + c.Param("schoolID")})
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func CreateSchool(c *gin.Context) {
	// Start tracing span
	span := utils.BuildSpan(c.Request.Context(), "CreateSchool", oteltrace.WithAttributes(attribute.Key("Request-ID").String(c.GetHeader("Request-ID"))))
	defer span.End()

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
	// Start tracing span
	span := utils.BuildSpan(c.Request.Context(), "VerifySchool", oteltrace.WithAttributes(attribute.Key("Request-ID").String(c.GetHeader("Request-ID"))))
	defer span.End()

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
	// Start tracing span
	span := utils.BuildSpan(c.Request.Context(), "UnverifySchool", oteltrace.WithAttributes(attribute.Key("Request-ID").String(c.GetHeader("Request-ID"))))
	defer span.End()

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
