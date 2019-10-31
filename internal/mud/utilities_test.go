package mud

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitString(t *testing.T) {
	str := "a dog jumped over the moon"
	newStr := formatToWidth(str, 12)

	assert.Equal(t, `a dog jumped 
over the moon`, newStr)
}
