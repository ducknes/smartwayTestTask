package tools

import (
	"encoding/json"
	"io"
)

func ParseJson(reader io.Reader, data interface{}) (err error) {
	if b, err := io.ReadAll(reader); err == nil {
		if err = json.Unmarshal(b, data); err != nil {
			return err
		}
	}
	return
}
