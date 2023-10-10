package converter

import (
	"encoding/json"
	"errors"
	"math/big"
)

func StringToBigInt(s string) (*big.Int, error) {
	b, ok := new(big.Int).SetString(s, 10)
	if !ok {
		return nil, errors.New("converter error: unable to convert string to big int")
	}
	return b, nil
}

func StringToBigFloat(s string) (*big.Float, error) {
	b, ok := new(big.Float).SetString(s)
	if !ok {
		return nil, errors.New("converter error: unable to convert string to big float")
	}
	return b, nil
}

func ModelToJsonBytes(model any) ([]byte, error) {
	return json.Marshal(model)
}

func ModelToJsonString(model any) (string, error) {
	bytes, err := ModelToJsonBytes(model)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
