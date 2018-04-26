package main

import (
	"encoding/json"
	"fmt"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"studentMobilityServer/dbManager"
)

func main() {
	//init dbManager with the db path 
	err := dbManager.OpenDBAccess("data/studentMobility.sql")
	checkErr(err)
	if err == nil{
		r := mux.NewRouter()
		r.HandleFunc("/getStudentData", getStudentData).Methods("GET")
		r.HandleFunc("/putStudentData", putStudentData).Methods("PUT")
		r.HandleFunc("/removeStudentData", removeStudentData).Methods("POST")

		http.ListenAndServe(":3000", r)
	
		err = dbManager.CloseDBAccess()
		checkErr(err)
	}

}
//GET 

//used to get data from student 
func getStudentData(w http.ResponseWriter, r *http.Request){
	//id := get id from get request  
	//student := dbManager.SelectStudent(id)
	//json.NewEncoder(w).Encode(student)
}

//PUT

//used to create and modify data of a student 
//if data.json had no id field is considered has a student creation 
func putStudentData(w http.ResponseWriter, r *http.Request){
	//add a parameter to know if its an update or other 
	var err error 
	var student dbManager.Student 
	json.NewDecoder(r.Body).Decode(&student)
	//update 
	//dbManager.UpdateStudent(student)
	fmt.Println(student.Id)
	if student.Id != 0{
		err = dbManager.UpdateStudent(student)
	}else{
		err = dbManager.CreateStudent(student)
	}
	if err == nil{
		err = errors.New("OK")
	}
	fmt.Fprintf(w, "%q",err)
}

func removeStudentData(w http.ResponseWriter, r *http.Request) {
	//add id inside get params 
	var err error 
	var student dbManager.Student 
	fmt.Println(r.Body)
	json.NewDecoder(r.Body).Decode(&student)
	if student.Id != 0{
		//test id in db 
		var idInDB bool
		idInDB, err = dbManager.FindID(student.Id)
		if idInDB{
			err = dbManager.RemoveStudent(student)
		}
	}
	if err == nil{
		err = errors.New("OK")
	}
	fmt.Fprintf(w, "%q",err)
}

//used to check and handle error (replace print by logPrint)
func checkErr(err error, args ...string){
	if err != nil {
		fmt.Println("Error")
		fmt.Print(err)
		fmt.Println(args)

	}
}