package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Exercise2/Config"
	"github.com/Exercise2/Controllers"
	"github.com/Exercise2/Models"
	"github.com/Exercise2/Routes"
	"github.com/jinzhu/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestGet(t *testing.T) {
	//open sql connection by gorm
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	//setup router via gin
	router := Routes.SetupRouter()
	router.GET("/user-api/user/", Controllers.GetUsers)

	//send get request
	req, _ := http.NewRequest("GET", "/user-api/user/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	//print response body
	fmt.Println(resp.Body.String(),"some text")

	//check test case by response code
	if resp.Code != 200 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Code, 200)
	}
}

func TestPost(t *testing.T) {
	//Setup and open sql db
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	Config.DB.AutoMigrate(&Models.Student{})
	//defer Config.DB.Close()

	//setup router
	router := Routes.SetupRouter()
	router.POST("/user-api/user/",Controllers.CreateUser)

	//send request
	newStudent := Models.Student{
		FirstName: "Boy",
		LastName: "Random",
		Subject: "Physics",
		Marks: 55,
	}

	responseBody,_ := json.Marshal(newStudent)
	req, _ := http.NewRequest("POST", "/user-api/user/", bytes.NewBuffer([]byte(responseBody)))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	//read body of response
	fmt.Println(resp.Body.String(),"some text posted",resp.Code)

	//check for test case by status code
	if resp.Code != http.StatusOK {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Code, http.StatusOK)
	}
}

func TestPut(t *testing.T) {
	//Setup and open sql db
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	Config.DB.AutoMigrate(&Models.Student{})
	//defer Config.DB.Close()

	//setup router
	router := Routes.SetupRouter()
	router.PUT("/user-api/user/42/",Controllers.UpdateUser)

	//send request
	newStudent := Models.Student{
		Marks:     69,
	}
	student,_ := json.Marshal(newStudent)
	req, _ := http.NewRequest("PUT", "/user-api/user/42/", bytes.NewBuffer(student))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	//read body of response
	fmt.Println(resp.Body.String(),"some text updated")

	//check for test case by status code
	if resp.Code != http.StatusOK {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Code, http.StatusOK)
	}
}

func TestDelete(t *testing.T) {
	//open sql connection by gorm
	Config.DB,_ = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	defer Config.DB.Close()

	//setup router via gin
	router := Routes.SetupRouter()
	router.DELETE("/user-api/user/21/",Controllers.DeleteUser )

	//send get request
	req, _ := http.NewRequest("DELETE", "/user-api/user/21/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	//print response body
	fmt.Println(resp.Body.String(),"some text deleted")

	//check test case by response code
	if resp.Code != 200 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Code, 200)
	}
}
