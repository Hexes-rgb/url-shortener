package random

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRandomString(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{
			name: "size = 1",
			size: 1,
		},
		{
			name: "size= 10",
			size: 10,
		},
		{
			name: "size = 30",
			size: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res1 := NewRandomString(tt.size)
			res2 := NewRandomString(tt.size)

			assert.Len(t, res1, tt.size)
			assert.Len(t, res2, tt.size)

			assert.NotEqual(t, res1, res2)
		})
	}
}
