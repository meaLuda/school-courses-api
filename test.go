package main

import (
	"fmt"
	"collegeAPI/models"
)

func Test_db_conntection_getCourses() {
	// Test connection result
	courses, err_ := models.GetDiplomaModule(32)

	CheckErr(err_)

	if courses == nil {
		fmt.Println("No records found")
		return
	} else {
		// return records
		fmt.Println(courses)
	}
}

func Test_db_conntection_getNotes(){
	notes, err := models.GetDiplomaSubModuleNotes(32,4,35)
	CheckErr(err)

	if notes == nil {
		fmt.Println("No records found")
		return
	} else {
		// return records
		fmt.Println("-------------------- DATA -----------------------")
		fmt.Println(notes)
	}
}

