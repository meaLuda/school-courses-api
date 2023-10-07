package main

import (
	"testing"
	"fmt"
	"log"
)


func TestGetDiplomas(t *testing.T) {
	err := ConnectDb()
	CheckErr(err)
	// Test connection result
	courses, err_ := GetDiplomasByDepartment(1)

	if err_ != nil {
		log.Fatal(err_)
	}

	if courses == nil {
		fmt.Println("No records found")
		return
	} else {
		// return records
		fmt.Println("--------------------- Records found -------------------------")

		fmt.Println(courses)
	}
}


func TestGetDiplomaModule(t *testing.T) {
	err := ConnectDb()
	CheckErr(err)
	// Test connection result
	courses, err_ := GetDiplomaModule(32)

	if err_ != nil {
		log.Fatal(err_)
	}

	if courses == nil {
		fmt.Println("No records found")
		return
	} else {
		// return records
		fmt.Println("--------------------- Records found -------------------------")

		fmt.Println(courses)
	}
}