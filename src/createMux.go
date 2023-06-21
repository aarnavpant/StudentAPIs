package src

import (
	"github.com/gorilla/mux"
)

func CreateNewMux() *mux.Router {
	muxNew := mux.NewRouter()
	muxNew.HandleFunc("/getStudent", getStudent).Methods("GET")
	muxNew.HandleFunc("/getStudents", getStudents).Methods("GET")
	muxNew.HandleFunc("/deleteStudent", deleteStudent).Methods("DELETE")
	muxNew.HandleFunc("/addStudent", addStudent).Methods("POST")
	return muxNew
}