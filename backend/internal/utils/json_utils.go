package utils

import (
	"encoding/json"
)

// MarshalUintOmitZero helps omit uint fields if they are zero.
func MarshalUintOmitZero(src interface{}, field string, value uint) ([]byte, error) {
	type Alias struct {
		*json.RawMessage
	}

	aux := make(map[string]interface{})

	data, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return nil, err
	}

	if value == 0 {
		delete(aux, field)
	}

	return json.Marshal(aux)
}
