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

func get_DiplomaSubModule(c *gin.Context) {
	diploma_id,_ := strconv.Atoi(c.Param("diploma_id")) // get id from param and convert to int
	courses, err := models.GetDiplomaSubModule(diploma_id)

	CheckErr(err)

	if courses == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		// return persons
		c.JSON(http.StatusOK, gin.H{"data": courses})
	}

}

func get_DiplomaSubModule_ModeID(c *gin.Context) {
	diploma_id,_ := strconv.Atoi(c.Param("diploma_id")) // get id from param and convert to int
	module_id,_ := strconv.Atoi(c.Param("module_id")) 
	courses, err := models.GetDiplomaSubModule_ModeID(diploma_id,module_id)

	CheckErr(err)

	if courses == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		// return persons
		c.JSON(http.StatusOK, gin.H{"data": courses})
	}

}

func get_DiplomaSubModule_Notes(c *gin.Context) {
	diploma_id,_ := strconv.Atoi(c.Param("diploma_id")) // get id from param and convert to int
	module_id,_ := strconv.Atoi(c.Param("module_id")) 
	sub_module_id,_ := strconv.Atoi(c.Param("sub_module_id")) 
	notes, err := models.GetDiplomaSubModuleNotes(diploma_id,module_id,sub_module_id)

	CheckErr(err)

	if notes == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		// return persons
		c.JSON(http.StatusOK, gin.H{"data": notes})
	}

}