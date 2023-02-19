package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Sample struct {
	Email string
}

func Test_GetStructValue(t *testing.T) {
	t.Run("Should get the struct's attribute value correctly", func(t *testing.T) {
		sample := Sample{Email: "rangga@test.com"}

		result, err := GetStructValue(sample, "Email", "notrangga@test.com")

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, *result, "rangga@test.com")
	})
	t.Run("Should use the fallback value if selected attribute has zero value, or not found", func(t *testing.T) {
		sample := Sample{Email: "rangga@test.com"}

		result, err := GetStructValue(sample, "emailzzz", "notrangga@test.com")

		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, *result, "notrangga@test.com")
	})
	t.Run("Should throw an error if the provided obj is not a struct", func(t *testing.T) {
		result, err := GetStructValue("notastruct", "emailzzz", "notrangga@test.com")

		assert.Nil(t, result)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "provided obj is not a struct")
	})
}
