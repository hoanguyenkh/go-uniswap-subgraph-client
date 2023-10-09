package converter

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToBigInt(t *testing.T) {
	t.Run("when valid", func(t *testing.T) {
		s := "123456789"
		expected := big.NewInt(123456789)

		b, err := StringToBigInt(s)
		assert.Nil(t, err)
		assert.Equal(t, expected, b)
	})

	t.Run("when invalid", func(t *testing.T) {
		s := "not a number"

		_, err := StringToBigInt(s)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "unable to convert")
	})
}

func TestStringToBigFloat(t *testing.T) {
	t.Run("when valid", func(t *testing.T) {
		s := "12.3456"
		expected := big.NewFloat(12.3456)

		_, err := StringToBigFloat(s)
		assert.Nil(t, err)
		assert.Equal(t, expected.String(), s)
	})

	t.Run("when invalid", func(t *testing.T) {
		s := "not a number"

		_, err := StringToBigFloat(s)
		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "unable to convert")
	})
}
