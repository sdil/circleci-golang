package main

import (
	"testing"
	// "github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"github.com/labstack/echo"
	"io/ioutil"
)

func TestRootEndpoint(t *testing.T) {
	// Setup
	e := echo.New()
	e.GET("/", HelloWorld)
	r , _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	e.ServeHTTP(w,r)
	resp := w.Result()

	// Assert status code
	if resp.StatusCode != http.StatusOK {
		t.Error("Failed endpoint with unexpected status code: ", resp.Status)
	}

	// Assert body
	body, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	if bodyString != "Hello World!" {
		t.Error("Unexpected string returned: ", bodyString)
	}
}

func BenchmarkRootEndpoint( b *testing.B) {
	// Setup
	e := echo.New()
	e.GET("/", HelloWorld)
	r , _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	for i:= 0; i < b.N; i++ {
		e.ServeHTTP(w,r)
	}
}
