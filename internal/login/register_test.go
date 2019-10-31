package login

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/brianbroderick/agora"
	"github.com/stretchr/testify/assert"
)

// Make sure you run the seed before running these tests

func init() {
	os.Setenv("PLATFORM_ENV", "test")
}

// TestUserExists: This user is created with the seed.
func TestUserExists(t *testing.T) {
	c := agora.NewDgraphTxn()
	defer c.DiscardTxn()

	u := User{Name: "azkul"}

	exists, user, err := UserExists(c, u)
	assert.NoError(t, err)
	assert.NotEmpty(t, user.UID)
	assert.True(t, exists)
}

func TestUserDoesNotExist(t *testing.T) {
	c := agora.NewDgraphTxn()
	defer c.DiscardTxn()

	u := User{Name: "not_a_user"}

	exists, user, err := UserExists(c, u)
	assert.NoError(t, err)
	assert.Equal(t, "", user.UID)
	assert.False(t, exists)
}

func TestCreateUser(t *testing.T) {
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(10000)

	created, err := CreateUser(fmt.Sprintf("new_user_%d", rnd), "123456")
	assert.NoError(t, err)
	assert.NotNil(t, created.UID)
}
