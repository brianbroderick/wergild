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

func TestToSlug(t *testing.T) {
	str := "!Billy The 3rd!"
	newStr := ToSlug(str)
	assert.Equal(t, `billy_the_3rd`, newStr)
}
