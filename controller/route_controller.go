package controller

import (
	"berkeley/service"
	"berkeley/utils"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/berkeley/ping", Ping)
	router.GET("/schools", GetAllSchools)
	router.GET("/schools/:schoolID", GetSchoolByID)
	router.POST("/schools/:schoolID", CreateSchool)
	router.POST("/schools/:schoolID/verify", VerifySchool)
	router.DELETE("/schools/:schoolID/verify", UnverifySchool)
	router.GET("/schools/:schoolID/tags", GetTagsForSchool)
	router.POST("/schools/:schoolID/tags", AddTagForSchool)
	router.DELETE("/schools/:schoolID/tags", RemoveTagForSchool)
	router.GET("/schools/:schoolID/emails", GetEmailsForSchool)
	router.POST("/schools/:schoolID/emails", AddEmailForSchool)
	router.DELETE("/schools/:schoolID/emails", RemoveEmailForSchool)
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.SugarLogger.Infoln("GATEWAY REQUEST ID: " + c.GetHeader("Request-ID"))
		c.Next()
	}
}

func AuthChecker() gin.HandlerFunc {
	return func(c *gin.Context) {

		var requestUserID string

		ctx := context.Background()
		client, err := service.FirebaseAdmin.Auth(ctx)
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}
		if c.GetHeader("Authorization") != "" {
			token, err := client.VerifyIDToken(ctx, strings.Split(c.GetHeader("Authorization"), "Bearer ")[1])
			if err != nil {
				utils.SugarLogger.Errorln("error verifying ID token")
				requestUserID = "null"
			} else {
				utils.SugarLogger.Infoln("Decoded User ID: " + token.UID)
				requestUserID = token.UID
			}
		} else {
			utils.SugarLogger.Infoln("No user token provided")
			requestUserID = "null"
		}
		c.Set("userID", requestUserID)
		// The main authentication gateway per request path
		// The requesting user's ID and roles are pulled and used below
		// Any path can also be quickly halted if not ready for prod
		c.Next()
	}
}

func contains(s []string, element string) bool {
	for _, i := range s {
		if i == element {
			return true
		}
	}
	return false
}
