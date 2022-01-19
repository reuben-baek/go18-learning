package option_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go18-learning/generic"
	"go18-learning/generic/option"
	"testing"
)

func TestOption(t *testing.T) {
	t.Run("int Some", func(t *testing.T) {
		intOpt := option.Some[int](10)
		assert.False(t, intOpt.IsNone())
		assert.Equal(t, 10, intOpt.Get())
	})
	t.Run("int Some Foreach", func(t *testing.T) {
		intOpt := option.Some[int](10)
		called := 0
		intOpt.ForEach(func(v int) {
			assert.Equal(t, 10, v)
			called++
		})
		assert.Equal(t, 1, called)
	})

	t.Run("int None", func(t *testing.T) {
		none := option.None[int]()
		assert.True(t, none.IsNone())
		assert.Equal(t, 0, none.Get())
		assert.Equal(t, 1, none.GetOrElse(1))
	})
	t.Run("int None ForEach", func(t *testing.T) {
		none := option.None[int]()
		called := 0
		none.ForEach(func(v int) {
			called++
		})
		assert.Equal(t, 0, called)
	})

	t.Run("option Map 1", func(t *testing.T) {
		intOpt := option.Some[int](10)
		intOpt2 := option.Map(intOpt, func(v int) int {
			return v * v
		})
		assert.Equal(t, 100, intOpt2.Get())
	})

	t.Run("option Map 2", func(t *testing.T) {
		intOpt := option.Some[int](10)
		stringOpt := option.Map(intOpt, func(v int) string {
			return fmt.Sprintf("v-%d", v)
		})
		assert.Equal(t, "v-10", stringOpt.Get())
	})

	t.Run("option FlatMap", func(t *testing.T) {
		intOpt := option.Some[int](10)
		intOpt2 := option.FlatMap(intOpt, func(v int) generic.Option[int] {
			return option.Some[int](v * v)
		})
		assert.Equal(t, 100, intOpt2.Get())
	})
}
