package main

import (
	"log"
	"collegeAPI/models"
	"net/http"
	"github.com/gin-gonic/gin"
)



func Init(){
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	err := models.ConnectDb()
	CheckErr(err)
	r := gin.Default()


	// test connection
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})
	v1 := r.Group("/api/college_app/v1/")
	{
		v1.GET("all_diplomas/:dpID",getDiplomaByDepartment)
		v1.GET("all_modules/:id",get_DiplomaModule) // http://localhost:8080/api/college_app/v1/all_modules/32
		v1.GET("modules/notes/:diploma_id/:module_id",get_DiplomaModule_Notes) // http://localhost:8080/api/college_app/v1/sub_modules/notes
	}

	
	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run(":8080")
	// ginLambda = ginadapter.New(r)
}


func main() {
	// since we are deploying to lambda we should
	Init()
}

