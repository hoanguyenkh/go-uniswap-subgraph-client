package converter

import (
	"math"
	"math/big"
	"testing"

	unigraphclient "github.com/emersonmacro/go-uniswap-subgraph-client"

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

func TestModelToJsonBytes(t *testing.T) {
	t.Run("when valid", func(t *testing.T) {
		model := unigraphclient.Transaction{
			ID:          "test",
			BlockNumber: "1",
			Timestamp:   "1",
			GasUsed:     "1",
			GasPrice:    "1",
		}

		expected := []byte("{\"id\":\"test\",\"blockNumber\":\"1\",\"timestamp\":\"1\",\"gasUsed\":\"1\",\"gasPrice\":\"1\"}")

		bytes, err := ModelToJsonBytes(model)
		assert.Nil(t, err)
		assert.Equal(t, expected, bytes)
	})

	t.Run("when invalid", func(t *testing.T) {
		_, err := ModelToJsonBytes(math.Inf(1))
		assert.NotNil(t, err)
	})
}

func TestModelToJsonString(t *testing.T) {
	t.Run("when valid", func(t *testing.T) {
		model := unigraphclient.Transaction{
			ID:          "test",
			BlockNumber: "1",
			Timestamp:   "1",
			GasUsed:     "1",
			GasPrice:    "1",
		}

		expected := "{\"id\":\"test\",\"blockNumber\":\"1\",\"timestamp\":\"1\",\"gasUsed\":\"1\",\"gasPrice\":\"1\"}"

		bytes, err := ModelToJsonString(model)
		assert.Nil(t, err)
		assert.Equal(t, expected, bytes)
	})

	t.Run("when invalid", func(t *testing.T) {
		_, err := ModelToJsonString(math.Inf(1))
		assert.NotNil(t, err)
	})
}
