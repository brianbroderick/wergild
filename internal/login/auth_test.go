package login

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Make sure you run the seed before running these tests

func init() {
	os.Setenv("PLATFORM_ENV", "test")
}

func TestMain(m *testing.M) {
	// // mud.ReloadData()
	// mud.BcryptCost = 4 // set to minimum for testing
	retCode := m.Run()
	// myTeardownFunction()
	os.Exit(retCode)
}

func TestAuth(t *testing.T) {
	auth, err := Auth("azkul", "password")
	assert.NoError(t, err)
	assert.Equal(t, true, auth)
}

func TestFailedAuth(t *testing.T) {
	auth, err := Auth("azkul", "1234567")
	assert.NoError(t, err)
	assert.Equal(t, false, auth)
}

func TestInvalidUser(t *testing.T) {
	auth, err := Auth("invalid", "1234567")
	assert.NoError(t, err)
	assert.Equal(t, false, auth)
}
