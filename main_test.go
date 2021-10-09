package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testuser_handle(t *testing.T) {
	request, _ := http.NewRequest("GET", "/users/2", nil)
	response := httptest.NewRecorder()
	user_handle(response, request)
	if response.Code != 200 {
		t.Errorf("Response code is %v", response.Code)
	}
	var post Post
	json.Unmarshal(response.Body.Bytes(), &post)
	if post.ID != "1" {
		t.Error("Cannot retrieve")
	}
}
func testpost_handle(t *testing.T) {
	request, _ := http.NewRequest("GET", "/posts/9", nil)
	response := httptest.NewRecorder()
	post_handle(response, request)
	if response.Code != 200 {
		t.Errorf("Response code is %v", response.Code)
	}
	var post Post
	json.Unmarshal(response.Body.Bytes(), &post)
	if post.ID != "1" {
		t.Error("Cannot retrieve")
	}
}
