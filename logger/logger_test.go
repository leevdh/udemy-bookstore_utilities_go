package logger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLogger(t *testing.T) {
	logger := GetLogger()
	fmt.Printf("%+v", logger)
	assert.NotEmpty(t, logger)
}
