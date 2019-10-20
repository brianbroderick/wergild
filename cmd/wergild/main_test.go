package main

import (
	"os"
	"testing"

	"github.com/brianbroderick/wergild/internal/mud"
)

func init() {
	os.Setenv("PLATFORM_ENV", "test")
	os.Setenv("LOG_LEVEL", "DEBUG")
}

func TestMain(m *testing.M) {
	mud.ReloadData()
	mud.BcryptCost = 4 // set to minimum for testing
	retCode := m.Run()
	// myTeardownFunction()
	os.Exit(retCode)
}

// Show
// {
// 	q(func: type(Room)) {
// 		uid
// 		roomName
// 		exits {
// 			dest {
// 				uid
// 				roomName
// 			}
// 		}
// 	}
// }
