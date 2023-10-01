package models

import (
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)

// ----------------------------------------------  structs for college --------------------------------
type CollegeDepartmentFields struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type CollegeDiplomaFields struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	DepertmentID int    `json:"depertment_id"`
}

type CollegeModuleFields struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	DiplomaID int    `json:"diploma_id"`
}

type CollegeCourseContentNotesFields struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	NotesPdf    string `json:"notes_pdf"`
	DiplomaID   int    `json:"diploma_id"`
	ModuleID    int    `json:"module_id"`
	// SubModuleID int    `json:"sub_module_id"`
}

// ---------------------------------------------- queries for college apps  --------------------------------

// get diplomas by depertment ID
func GetDiplomasByDepartment(departmentID int) ([]CollegeDiplomaFields, error) {
	query := "SELECT * FROM CollegeDiploma WHERE depertment_id = ?"
	rows, err := DB.Query(query, departmentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close() //good habit to close

	diplomas := make([]CollegeDiplomaFields, 0)

	for rows.Next() {
		singlDiploma := CollegeDiplomaFields{}

		err := rows.Scan(&singlDiploma.ID, &singlDiploma.Title, &singlDiploma.DepertmentID)
		if err != nil {
			return nil, err
		}

		diplomas = append(diplomas, singlDiploma)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return diplomas, err
}

func GetDiplomaModule(diploma_id int) ([]CollegeModuleFields, error) {
	//query db
	// 	SELECT * FROM "CollegeDiploma_Module"
	// WHERE   diploma_id=32;
	rows, err := DB.Query("SELECT * FROM CollegeDiploma_Module WHERE diploma_id="+strconv.Itoa(diploma_id)+";")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	modules := make([]CollegeModuleFields, 0)

	for rows.Next() {
		// create instances to be appended
		singleModule := CollegeModuleFields{}
		err = rows.Scan(&singleModule.ID, &singleModule.Title, &singleModule.DiplomaID)

		if err != nil {
			return nil, err
		}

		modules = append(modules, singleModule)
	}
	
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return modules, err
}


func GetDiplomaModuleNotes(diploma_id int,module_id int)([]CollegeCourseContentNotesFields, error){

	//query db
	rows, err := DB.Query("SELECT * FROM CollegeDiploma_SubmoduleContent_notes WHERE diploma_id="+strconv.Itoa(diploma_id)+" AND module_id="+strconv.Itoa(module_id)+";")
	if err != nil {
		return nil, err
	}

	defer DB.Close()

	SubModulenotes := make([]CollegeCourseContentNotesFields, 0)

	for rows.Next() {
		// create instances to be appended
		singleSubModuleNote := CollegeCourseContentNotesFields{}
		err = rows.Scan(&singleSubModuleNote.ID, &singleSubModuleNote.Title, &singleSubModuleNote.NotesPdf,&singleSubModuleNote.DiplomaID,&singleSubModuleNote.ModuleID)

		if err != nil {
			return nil, err
		}

		SubModulenotes = append(SubModulenotes, singleSubModuleNote)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	rows.Close() //good habit to close
	return SubModulenotes, err
}