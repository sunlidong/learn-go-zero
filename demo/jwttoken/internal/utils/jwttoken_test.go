package utils

import (
	"testing"
)

func TestDecodeJwtToken(t *testing.T) {
	data := map[string]interface{}{
		"role": "1",
		"id":   "2",
	}
	token, err := EncodeJwtToken(data)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(token)
	result, err := DecodeJwtToken(token)
	if err != nil {
		t.Error(err)
		return
	}
	if result["id"] != "2" {
		t.Error(err)
		return
	}
	if result["role"] != "1" {
		t.Error(err)
		return
	}
}
