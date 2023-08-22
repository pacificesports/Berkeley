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

func GetEmailsForSchool(c *gin.Context) {
	// Start tracing span
	span := utils.BuildSpan(c.Request.Context(), "GetEmailsForSchool", oteltrace.WithAttributes(attribute.Key("Request-ID").String(c.GetHeader("Request-ID"))))
	defer span.End()

	result := service.GetEmailsForSchool(c.Param("schoolID"))
	c.JSON(http.StatusOK, result)
}

func AddEmailForSchool(c *gin.Context) {
	// Start tracing span
	span := utils.BuildSpan(c.Request.Context(), "AddEmailForSchool", oteltrace.WithAttributes(attribute.Key("Request-ID").String(c.GetHeader("Request-ID"))))
	defer span.End()

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
	// Start tracing span
	span := utils.BuildSpan(c.Request.Context(), "RemoveEmailForSchool", oteltrace.WithAttributes(attribute.Key("Request-ID").String(c.GetHeader("Request-ID"))))
	defer span.End()

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
