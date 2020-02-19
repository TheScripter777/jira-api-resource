package configuration

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var metaParams MetaParameters

func setup(t *testing.T) func(t *testing.T) {
	t.Log("setup test cases...")
	metaParams = MetaParameters{}

	return func(t *testing.T) {
		t.Log("teardown test cases...")
	}
}

// Test Mandatory Values Are Present From Valid Meta Parameters
func TestMetaParameters_MandatoryValues(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	t.Run("mandatory parameters PRESENT from VALID and PARSED params", func(t *testing.T) {
		// Arrange
		metaParams.parsed = true
		metaParams.valid = true
		metaParams.mandatoryPresent = true
		require.NotNil(t, metaParams)

		// Act
		result := metaParams.AllMandatoryValuesPresent()

		// Assert
		assert.Equal(t, true, result)
	})

	t.Run("mandatory parameters NOT PRESENT from INVALID and PARSED params", func(t *testing.T) {
		// Arrange
		metaParams.parsed = true
		metaParams.valid = false
		metaParams.mandatoryPresent = false
		require.NotNil(t, metaParams)

		// Act
		result := metaParams.AllMandatoryValuesPresent()

		// Assert
		assert.Equal(t, false, result)
	})

	t.Run("mandatory parameters NOT PRESENT from VALID and PARSED params", func(t *testing.T) {
		// Arrange
		metaParams.parsed = true
		metaParams.valid = true
		metaParams.mandatoryPresent = false
		require.NotNil(t, metaParams)

		// Act
		result := metaParams.AllMandatoryValuesPresent()

		// Assert
		assert.Equal(t, false, result)
	})
}

func TestMetaParameters_ReadyState(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	t.Run("is READY from VALID and PARSED params", func(t *testing.T) {
		// Arrange
		metaParams.parsed = true
		metaParams.valid = true
		metaParams.mandatoryPresent = false
		require.NotNil(t, metaParams)

		// Act
		result := metaParams.Ready()

		// Assert
		assert.Equal(t, true, result)
	})

	t.Run("is NOT READY from INVALID and PARSED params", func(t *testing.T) {
		// Arrange
		metaParams.parsed = true
		metaParams.valid = false
		metaParams.mandatoryPresent = false
		require.NotNil(t, metaParams)

		// Act
		result := metaParams.Ready()

		// Assert
		assert.Equal(t, false, result)
	})

	t.Run("is NOT READY from VALID and UNPARSED params", func(t *testing.T) {
		// Arrange
		metaParams.parsed = false
		metaParams.valid = true
		metaParams.mandatoryPresent = false
		require.NotNil(t, metaParams)

		// Act
		result := metaParams.Ready()

		// Assert
		assert.Equal(t, false, result)
	})
}
