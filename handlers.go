package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	// "collegeAPI/models"
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
	dpmID,_ := strconv.Atoi(c.Param("dpmID"))
	fmt.Printf("id: %d",dpmID)
	// query db for diplomas by department id
	diplomas, err := GetDiplomasByDepartment(dpmID)

	// Check error
	CheckErr(err)

	if diplomas != nil{
		// return a json
		c.JSON(http.StatusOK, gin.H{"data": diplomas})
		return
	}else{
		c.JSON(http.StatusBadRequest, gin.H{"ERROR":"No records found!"})
	}
}

func get_DiplomaModule(c *gin.Context) {

	dplid,_ := strconv.Atoi(c.Param("dplid")) // get id from param and convert to int
	courses, err := GetDiplomaModule(dplid)

	CheckErr(err)

	if courses == nil || len(courses) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else if courses != nil && len(courses) >= 1 {
		c.JSON(http.StatusOK, gin.H{"data": courses})
	}

}

func get_DiplomaModule_Notes(c *gin.Context) {
	diploma_id,_ := strconv.Atoi(c.Param("diploma_id")) // get id from param and convert to int
	module_id,_ := strconv.Atoi(c.Param("module_id")) 
	notes, err := GetDiplomaModuleNotes(diploma_id,module_id)

	CheckErr(err)

	if notes == nil || len(notes) < 1 {
		fmt.Println(len(notes))
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
	} else if notes != nil && len(notes) > 1 {
		fmt.Println(len(notes))

		// return persons
		c.JSON(http.StatusOK, gin.H{"data": notes})
		return
	}

}