package main

import (
	"github.com/gin-gonic/gin"
	"collegeAPI/models"
)


func main() {
	r := gin.Default()
	err := models.ConnectDb()
	CheckErr(err)
	// Test_db_conntection_getNotes()
	// //API v1
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
}

