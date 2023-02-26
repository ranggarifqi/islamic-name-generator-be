package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RemoveSliceDuplicate(t *testing.T) {
	t.Run("Should remove duplicates correctly", func(t *testing.T) {
		input := []int{0, 0, 1, 2, 2, 2, 3, 3}
		result := RemoveSliceDuplicate(input)

		assert.Len(t, result, 4)
		assert.ElementsMatch(t, result, []int{0, 1, 2, 3})
	})

	t.Run("Should return the same set if no duplicates found", func(t *testing.T) {
		input := []int{0, 1, 2, 3}
		result := RemoveSliceDuplicate(input)

		assert.Len(t, result, 4)
		assert.ElementsMatch(t, result, []int{0, 1, 2, 3})
	})

	t.Run("String should be ordered asc", func(t *testing.T) {
		input := []string{"c", "a", "b", "a"}
		result := RemoveSliceDuplicate(input)

		assert.Len(t, result, 3)
		assert.Equal(t, []string{"a", "b", "c"}, result)
	})

	t.Run("Int should be ordered asc", func(t *testing.T) {
		input := []int{3, 3, 2, 2, 2, 0, 0, 1}
		result := RemoveSliceDuplicate(input)

		assert.Len(t, result, 4)
		assert.Equal(t, []int{0, 1, 2, 3}, result)
	})
}
