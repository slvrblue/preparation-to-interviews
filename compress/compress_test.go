package compress

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompress1(t *testing.T) {
	res, err := Compress("")
	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestCompress2(t *testing.T) {
	res, err := Compress("A/")
	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestCompress3(t *testing.T) {
	res, err := Compress("A1")
	assert.Error(t, err)
	assert.Empty(t, res)
}

func TestCompress4(t *testing.T) {
	res, err := Compress("A")
	assert.NoError(t, err)
	assert.Equal(t, res, "A1")
}

func TestCompress5(t *testing.T) {
	res, err := Compress("AAAABBCDD")
	assert.NoError(t, err)
	assert.Equal(t, res, "A4B2CD2")
}
