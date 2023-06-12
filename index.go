package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/aarnavpant/StudentAPIs.git/students"
	"github.com/gorilla/mux"
)

func main() {
	studs := []students.Student{
		{
			Name:   "Aarnav",
			RollNo: 1,
			// subjects: []string{"DSA", "ML", "STM", "DevOps"},
		},
		{
			Name:   "Pranav",
			RollNo: 2,
			// subjects: []string{"DSA", "ML", "STM", "DevOps"},
		},
		{
			Name:   "Sarthak",
			RollNo: 3,
			// subjects: []string{"DSA", "ML", "STM", "DevOps"},
		},
		{
			Name:   "Avinash",
			RollNo: 4,
			// subjects: []string{"DSA", "OOPS", "Maths", "4y"},
		},
	}

	//handler for getting student by id
	handler1 := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			fmt.Fprintf(w, "Method not allowed")
			return
		}
		roll, error := strconv.Atoi(r.URL.Query().Get("id"))
		if error != nil {
			fmt.Fprintf(w, "has to be number")
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

	//handler for getting all students
	handler2 := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			fmt.Fprintf(w, "Method not allowed")
			return
		}
		jData, err := json.Marshal(studs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jData)
	}

	//handler to delete a student
	handler3 := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			fmt.Fprintf(w, "Method not allowed")
			return
		}
		roll, error := strconv.Atoi(r.URL.Query().Get("id"))
		if error != nil {
			http.Error(w, error.Error(), http.StatusBadRequest)
			return
		}
		newStud := []students.Student{}
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

	//handler for POST call for adding a student
	handler4 := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			fmt.Fprintf(w, "Wrong method")
			return
		}
		var stud students.Student
		err := json.NewDecoder(r.Body).Decode(&stud)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		studs = append(studs, stud)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Record Created")
	}

	// making endpoints for handlers
	muxNew := mux.NewRouter()
	muxNew.HandleFunc("/getStudent", handler1)
	muxNew.HandleFunc("/getStudents", handler2)
	muxNew.HandleFunc("/deleteStudent", handler3)
	muxNew.HandleFunc("/addStudent", handler4)

	fmt.Println("server started")
	log.Fatal(http.ListenAndServe(":8081", muxNew))
}
