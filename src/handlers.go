package src

import (
	"net/http"
	"fmt"
	"strconv"
	"encoding/json"
)

var studs = []Student{
	{
		Name:   "Aarnav",
		RollNo: 1,
	},
	{
		Name:   "Pranav",
		RollNo: 2,
	},
	{
		Name:   "Sarthak",
		RollNo: 3,
	},
	{
		Name:   "Avinash",
		RollNo: 4,
	},
}

// getStudents fetches all students
//	@Summary		Get all students
//	@Description	Get all students
//	@ID				get-students
//	@Tags			students
//	@Produce		json
//	@Success		200	{array}	Student
//	@Router			/getStudents [get]
func getStudent(w http.ResponseWriter, r *http.Request) {
	roll, error := strconv.Atoi(r.URL.Query().Get("id"))
	if error != nil {
		http.Error(w,error.Error(),http.StatusBadRequest)
	}
	flag := false
	for _, stud := range studs {
		if stud.RollNo == roll {
			jData, err := json.Marshal(stud)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.Write(jData)
			flag = true
			break
		}
	}
	if !flag {
		fmt.Fprintf(w, "No Record found")
	}
}

// getStudent fetches the student with the given id
//	@Summary		Get the student with given id
//	@Description	Get the student with given id
//	@ID				get-student
//	@Param			id	path	int	true	"id"
//	@Produce		json
//	@Success		200	{object}	Student
//	@Failure		404	{object}	string
//	@Router			/getStudent [get]
func getStudents(w http.ResponseWriter, r *http.Request) {
	jData, err := json.Marshal(studs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

// deleteStudent deletes a student with given id
//	@Summary		Deletes a student with given id
//	@Description	Deletes a student with given id
//	@ID				delete-student
//	@Param			id	path	int	true	"id"
//	@Produce		json
//	@Success		200	{object}	Student
//	@Failure		404	{object}	string
//	@Router			/deleteStudent [delete]
func deleteStudent(w http.ResponseWriter, r *http.Request) {
	roll, error := strconv.Atoi(r.URL.Query().Get("id"))
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}
	newStud := []Student{}
	flag := false
	for _, stud := range studs {
		if stud.RollNo != roll {
			newStud = append(newStud, stud)
		}
		if stud.RollNo == roll {
			flag = true
			fmt.Fprintf(w, "Record Deleted")
		}
	}
	studs = newStud
	if !flag {
		fmt.Fprintf(w, "No record to delete")
	}
}

// addStudent creates a student with given data
//	@Summary		Creates a student with given data
//	@Description	Creates a student with given data
//	@ID				create-student
//	@Param			stud	body	Student	true  "student" 
//	@Produce		json
//	@Success		200	{object}	string
//	@Failure		404	{object}	string
//	@Router			/addStudent [post]
func addStudent(w http.ResponseWriter, r *http.Request) {
	var stud Student
	err := json.NewDecoder(r.Body).Decode(&stud)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	studs = append(studs, stud)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Record Created")
}