package controllers

import (
	"bytes"
	"encoding/json"
	"go-user-api/models"
	"go-user-api/server"
	"net/http"
	"net/http/httptest"
	"strconv"

	"testing"
)

var uuid string

func TestGetUsers(t *testing.T) {
	response := httptest.NewRecorder()
	router := server.SetupRouter()
	req, err := http.NewRequest("GET", "http://localhost:8080/api/users", nil)
	if err != nil {
		t.Error(err)
	}
	router.ServeHTTP(response, req)
	if response.Code != http.StatusOK {
		t.Errorf("Got error for GET request to /api/users")
	}
}

func TestCreateUser(t *testing.T) {
	response := httptest.NewRecorder()
	router := server.SetupRouter()
	requestParams := `{
		"name": "NAME",
		"account_name": "ACCOUNTNAME",
		"email": "EMAIL"
	}`
	req, err := http.NewRequest("POST", "http://localhost:8080/api/users", bytes.NewBuffer([]byte(requestParams)))
	if err != nil {
		t.Error(err)
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, req)
	if response.Code != http.StatusCreated {
		t.Errorf("Got error for POST request to /api/users")
	}

	body := json.NewDecoder(response.Body)
	var user models.User
	body.Decode(&user)

	if user.Name != "NAME" ||
		user.AccountName != "ACCOUNTNAME" ||
		user.Email != "EMAIL" {
		t.Errorf("Create build failed.\nGot: %v", user)
	}

	uuid = strconv.Itoa(int(user.ID))
}

func TestGetUser(t *testing.T) {
	response := httptest.NewRecorder()
	router := server.SetupRouter()
	req, err := http.NewRequest("GET", "http://localhost:8080/api/users/"+uuid, nil)
	if err != nil {
		t.Error(err)
	}
	router.ServeHTTP(response, req)
	if response.Code != http.StatusOK {
		t.Errorf("Got error for GET request to /api/users/" + uuid)
	}
	body := json.NewDecoder(response.Body)
	var user models.User
	body.Decode(&user)

	if user.Name != "NAME" ||
		user.AccountName != "ACCOUNTNAME" ||
		user.Email != "EMAIL" {
		t.Errorf("Show build failed.\nGot: %v", user)
	}
}

func TestUpdateUser(t *testing.T) {
	response := httptest.NewRecorder()
	router := server.SetupRouter()
	requestParams := `{
		"name": "NAME_2",
		"account_name": "ACCOUNTNAME_2",
		"email": "EMAIL_2"
	}`
	req, err := http.NewRequest("PUT", "http://localhost:8080/api/users/"+uuid, bytes.NewBuffer([]byte(requestParams)))
	if err != nil {
		t.Error(err)
	}

	// Content-Type 設定
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, req)
	if response.Code != http.StatusOK {
		t.Errorf("Got error for PUT request to /api/users/" + uuid)
	}
	body := json.NewDecoder(response.Body)
	var user models.User
	body.Decode(&user)

	if user.Name != "NAME_2" ||
		user.AccountName != "ACCOUNTNAME_2" ||
		user.Email != "EMAIL_2" {
		t.Errorf("Update build failed.\nGot: %v", user)
	}
}

func TestDeleteUser(t *testing.T) {
	response := httptest.NewRecorder()
	router := server.SetupRouter()
	req, err := http.NewRequest("DELETE", "http://localhost:8080/api/users/"+uuid, nil)
	if err != nil {
		t.Error(err)
	}
	router.ServeHTTP(response, req)
	if response.Code != http.StatusNoContent {
		t.Errorf("Got error for DELETE request to /api/users/" + uuid)
	}
}
