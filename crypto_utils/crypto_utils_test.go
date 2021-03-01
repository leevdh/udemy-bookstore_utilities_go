package crypto_utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMD5(t *testing.T) {
	hash := GetMd5("testing")
	assert.NotEmpty(t, hash)
	assert.Equal(t, "ae2b1fca515949e5d54fb22b8ed95575", hash)
}
