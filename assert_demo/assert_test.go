package assert_demo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssert(t *testing.T) {
	err := errors.New("123")
	assert.Error(t, err)
}
