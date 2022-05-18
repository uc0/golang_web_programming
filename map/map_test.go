package _map

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	t.Run("nil map", func(t *testing.T) {
		var nilMap map[string]int
		assert.Equal(t, len(nilMap), 0)
		assert.Panics(t, func() {
			nilMap["a"] = 1
		})
		assert.Nil(t, nilMap)
	})

	t.Run("empty map", func(t *testing.T) {
		emptyMap1 := make(map[string]int)
		assert.Equal(t, 0, len(emptyMap1))
		assert.NotPanics(t, func() {
			emptyMap1["a"] = 1
		})

		emptyMap2 := map[string]int{}
		assert.NotNil(t, emptyMap2)
		assert.Equal(t, 0, len(emptyMap2))
		assert.NotPanics(t, func() {
			emptyMap2["a"] = 1
		})
	})
}
