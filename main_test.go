package main

import (
	"os"
	"testing"
)

func init() {
	os.Setenv("PLATFORM_ENV", "test")
	os.Setenv("LOG_LEVEL", "DEBUG")
}

func TestMain(m *testing.M) {
	reloadData()
	BcryptCost = 4 // set to minimum for testing
	retCode := m.Run()
	// myTeardownFunction()
	os.Exit(retCode)
}
