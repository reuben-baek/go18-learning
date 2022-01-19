package list_test

import (
	"github.com/stretchr/testify/assert"
	"go18-learning/generic"
	"go18-learning/generic/list"
	"testing"
)

func TestList(t *testing.T) {
	t.Run("factory", func(t *testing.T) {
		var l generic.List[int]
		l = list.New[int](1, 2, 3)
		assert.Equal(t, 3, l.Len())
		assert.Equal(t, 1, l.Head())
		assert.Equal(t, 2, l.Tail().Len())

		l = l.Prepend(4)
		assert.Equal(t, 4, l.Len())
		assert.Equal(t, 4, l.Head())
		assert.Equal(t, 3, l.Tail().Len())
	})
}
