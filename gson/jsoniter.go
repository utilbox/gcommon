// Package gson provides an alternative solution for json processing.
// It's faster than the standard json lib.
// For more information, see https://github.com/json-iterator/go
package gson

import jsoniter "github.com/json-iterator/go"

var json jsoniter.API

func init() {
	json = jsoniter.ConfigCompatibleWithStandardLibrary
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
