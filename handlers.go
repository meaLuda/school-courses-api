package main

import (
	_ "fmt"
	"github.com/gin-gonic/gin"
	"log"
	"collegeAPI/models"
	"net/http"
	"strconv"
)


// re-usable function to check errors
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


// -------------------------- courses functions ---------------------------

// Get diplomas by department dpID
func getDiplomaByDepartment(c *gin.Context){
	
	// take department id & vonvert into integer
	dpID,_ := strconv.Atoi(c.Param("dpID"))

	// query db for diplomas by department id
	diplomas, err := models.GetDiplomasByDepartment(dpID)

	// Check error
	CheckErr(err)

	if diplomas != nil{
		// return a json
		c.JSON(http.StatusBadRequest, gin.H{"ERROR":"No records found!"})
		return
	}else{
		c.JSON(http.StatusOK, gin.H{"data": diplomas})
	}
}

func get_DiplomaModule(c *gin.Context) {

	id,_ := strconv.Atoi(c.Param("id")) // get id from param and convert to int
	courses, err := models.GetDiplomaModule(id)

	CheckErr(err)

	if courses == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		// return persons
		c.JSON(http.StatusOK, gin.H{"data": courses})
	}

}

func get_DiplomaModule_Notes(c *gin.Context) {
	diploma_id,_ := strconv.Atoi(c.Param("diploma_id")) // get id from param and convert to int
	module_id,_ := strconv.Atoi(c.Param("module_id")) 
	notes, err := models.GetDiplomaModuleNotes(diploma_id,module_id)

	CheckErr(err)

	if notes == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		// return persons
		c.JSON(http.StatusOK, gin.H{"data": notes})
	}

}