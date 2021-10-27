package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)
type person struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
	Email string `json:"email"`
	Company string `json:"company"`
}

var personInfo []person

func listPersonInfo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(personInfo)
}

func creatPersonInfo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	var newperson person
	_=json.NewDecoder(r.Body).Decode(&newperson)
	newperson.Id=strconv.Itoa(rand.Intn(1000))
	personInfo=append(personInfo,newperson)
	json.NewEncoder(w).Encode(newperson)
}

func updatePersonInfo(w http.ResponseWriter,r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	param:=mux.Vars(r)
	for index,item:=range personInfo{
		if item.Id == param["id"] {
			personInfo=append(personInfo[:index],personInfo[index+1:]...)
			var newperson person
			_=json.NewDecoder(r.Body).Decode(&newperson)
			newperson.Id=strconv.Itoa(rand.Intn(1000))
			personInfo=append(personInfo,newperson)
			json.NewEncoder(w).Encode(newperson)
			return
		}
	}
	json.NewEncoder(w).Encode(personInfo)
}

func deletePersonInfo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	param:=mux.Vars(r)
	for index,item:=range personInfo{
		if item.Id==param["id"]{
			personInfo=append(personInfo[:index],personInfo[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(personInfo)
}

func deleteAllPersonInfo(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	personInfo=nil
	json.NewEncoder(w).Encode(personInfo)
}

func main()  {
	r:=mux.NewRouter()
	personInfo=append(personInfo,person{Id: "1", Name: "vikas", Age: 22, Gender: "male",Email: "vikas@gmail.com",Company: "appointy"})
	personInfo=append(personInfo,person{Id: "2", Name: "shivam", Age: 20, Gender: "male",Email: "shivam@gmail.com",Company: "mcc"})

	r.HandleFunc("/",listPersonInfo).Methods("GET")
	r.HandleFunc("/listpersoninfo",listPersonInfo).Methods("GET")
	r.HandleFunc("/createpersoninfo",creatPersonInfo).Methods("POST")
	r.HandleFunc("/updatepersoninfo/{id}", updatePersonInfo).Methods("PUT")
	r.HandleFunc("/deletepersoninfo/{id}",deletePersonInfo).Methods("DELETE")
	r.HandleFunc("/deleteallpersoninfo",deleteAllPersonInfo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000",r))
}