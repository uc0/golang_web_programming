package slice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice(t *testing.T) {
	t.Run("nil slice에는 값을 넣을 수 있다.", func(t *testing.T) {
		var nilSlice []int
		assert.Equal(t, len(nilSlice), 0)
		assert.Nil(t, nilSlice)
		assert.NotPanics(t, func() {
			nilSlice = append(nilSlice, 1)
		})
	})

	t.Run("empty slice", func(t *testing.T) {
		emptySlice1 := make([]int, 1, 2)
		assert.Equal(t, 1, len(emptySlice1))
		assert.Equal(t, 2, cap(emptySlice1))
		assert.NotPanics(t, func() {
			emptySlice1[0] = 1
		})

		emptySlice2 := []int{}
		// var emptySlice2 []int 랑 똑같다..!라고 IDE에서 보여준다.
		assert.NotNil(t, emptySlice2)
		assert.Equal(t, 0, len(emptySlice2))
		assert.NotPanics(t, func() {
			emptySlice2 = append(emptySlice2, 1)
		})
	})
}

func TestCopy(t *testing.T) {
	t.Run("b = a 하면 값을 공유한다.", func(t *testing.T) {
		a := []int{1, 2, 3}
		var b []int
		b = a
		b[0] = 10
		assert.Equal(t, b[0], a[0])
	})

	t.Run("부분으로 전달하면 값을 공유한다.", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := a[:1] // b (ptr, len, cap) b.ptr == a.ptr[0]
		b[0] = 100
		assert.Equal(t, b[0], a[0])

		d := []string{"r", "o", "a", "d"}
		e := d[2:]
		// e == []byte{'a', 'd'}
		e[1] = "m"
		fmt.Println(d)
		fmt.Println(e)
		// e == []byte{'a', 'm'}
		// d == []byte{'r', 'o', 'a', 'm'}
	})
}

func TestInitialize(t *testing.T) {
	t.Run("초기화했을 때 zero value로 되어있음", func(t *testing.T) {
		var s []byte
		s = make([]byte, 5, 5)
		for i := 0; i < 5; i++ {
			assert.EqualValues(t, 0, s[i])
		}
	})
}
