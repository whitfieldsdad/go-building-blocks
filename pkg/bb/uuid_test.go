package bb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUUID4(t *testing.T) {
	assert.NotEmpty(t, NewUUID4())
}

func TestNewUUID5(t *testing.T) {
	namespace := "cb3311f0-15f8-5a82-9f32-6b27fb088bb1"
	b := []byte("Hello world")
	u := NewUUID5(namespace, b)
	assert.Equal(t, "f958f841-478e-5319-8fdc-3641700c391d", u)
}

func TestNewUUID5FromMap(t *testing.T) {
	m := map[string]interface{}{
		"pid":  123,
		"ppid": 456,
	}
	namespace := "cb3311f0-15f8-5a82-9f32-6b27fb088bb1"
	u, err := NewUUID5FromMap(namespace, m)
	assert.Nil(t, err)
	assert.Equal(t, "658cadf9-708d-5c1a-a3f0-9b6e850df6ba", u)
}
