package bb

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileTraitsWithRegularFile(t *testing.T) {
	path, err := NewTempFile(nil)
	assert.Nil(t, err, "Error should be nil")
	defer DeleteFile(path)

	traits, err := GetFileTraits(path)
	assert.Nil(t, err, "Error should be nil")

	expected := []string{"regular_file"}
	result := traits.ToList()
	assert.Equal(t, expected, result, "Traits should match")
}

func TestGetFileTraitsWithDirectory(t *testing.T) {
	path, err := NewTempDir()
	assert.Nil(t, err, "Error should be nil")
	defer DeleteFile(path)

	traits, err := GetFileTraits(path)
	assert.Nil(t, err, "Error should be nil")

	expected := []string{"directory"}
	result := traits.ToList()
	assert.Equal(t, expected, result, "Traits should match")
}

func TestGetFileTraitsWithHardLink(t *testing.T) {
	path, err := NewTempFile(nil)
	assert.Nil(t, err, "Error should be nil")
	defer DeleteFile(path)

	link := path + ".link"
	err = os.Link(path, link)
	defer DeleteFile(link)

	assert.Nil(t, err, "Error should be nil")
}

func TestGetFileTraitsWithNonExistentFile(t *testing.T) {
	path := NewUUID4()
	traits, err := GetFileTraits(path)
	assert.NotNil(t, err, "Error should not be nil")
	assert.Nil(t, traits, "Traits should be nil")
}
