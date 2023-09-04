package main

import (
	"log"
	"context"
	"collegeAPI/models"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda


func Init(){
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")

	r := gin.Default()
	err := models.ConnectDb()
	CheckErr(err)
	// Test_db_conntection_getNotes()
	// //API v1

	// test connection
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{
			"message":"pong",
		})
	})
	v1 := r.Group("/api/college_app/v1")
	{
		v1.GET("all_modules/:id",get_DiplomaModule) // http://localhost:8080/api/college_app/v1/all_modules/32
		v1.GET("sub_modules/:diploma_id",get_DiplomaSubModule) // http://localhost:8080/api/college_app/v1/sub_modules/32/4
		v1.GET("sub_modules/with_module/:diploma_id/:module_id",get_DiplomaSubModule_ModeID) // http://localhost:8080/api/college_app/v1/sub_modules/32/4
		v1.GET("sub_modules/notes/:diploma_id/:module_id/:sub_module_id",get_DiplomaSubModule_Notes) // http://localhost:8080/api/college_app/v1/sub_modules/notes
		//   v1.POST("person", addPerson)
		//   v1.PUT("person/:id", updatePerson)
		//   v1.DELETE("person/:id", deletePerson)
		//   v1.OPTIONS("person", options)
	}

	
	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run(":8080")
	// ginLambda = ginadapter.New(r)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	// since we are deploying to lambda we should
	Init()
	lambda.Start(Handler)
}

