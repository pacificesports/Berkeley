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

func GetTagsForSchool(c *gin.Context) {
	// Start tracing span
	span := utils.BuildSpan(c.Request.Context(), "GetTagsForSchool", oteltrace.WithAttributes(attribute.Key("Request-ID").String(c.GetHeader("Request-ID"))))
	defer span.End()

	result := service.GetTagsForSchool(c.Param("schoolID"))
	c.JSON(http.StatusOK, result)
}

func AddTagForSchool(c *gin.Context) {
	// Start tracing span
	span := utils.BuildSpan(c.Request.Context(), "AddTagForSchool", oteltrace.WithAttributes(attribute.Key("Request-ID").String(c.GetHeader("Request-ID"))))
	defer span.End()

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
	// Start tracing span
	span := utils.BuildSpan(c.Request.Context(), "RemoveTagForSchool", oteltrace.WithAttributes(attribute.Key("Request-ID").String(c.GetHeader("Request-ID"))))
	defer span.End()

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
