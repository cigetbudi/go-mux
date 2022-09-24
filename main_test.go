package main_test

import (
	"os"
	"testing"
	"github.com/cigetbudi/go-mux"
	_"github.com/cigetbudi/go-mux"

)

var a main.App

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("App_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	
}

func ensureTableExists() {
	if _, err := a.DB.Exec()
}

