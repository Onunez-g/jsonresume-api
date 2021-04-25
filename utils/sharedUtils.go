package utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

func ReadFromBody(body io.ReadCloser, val interface{}) error {
	value, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(value, &val); err != nil {
		return err
	}
	return nil
}

func UpdateResume(resume interface{}, list interface{}) {

	copy(resume.([]interface{}), list.([]interface{}))
}
