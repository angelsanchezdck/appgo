package common

import (
	"encoding/json"
	"fmt"
	"io"
)

func ParseBody(body io.ReadCloser) (*File, error) {
	file := File{}

	err := json.NewDecoder(body).Decode(&file)
	if err != nil {
		return &file, err
	}
	if file.FileName == "" {
		err = fmt.Errorf("missing FileName: %s", file.FileName)
		return &file, err
	}
	if file.Content == "" {
		err = fmt.Errorf("missing Content: %s", file.Content)
		return &file, err
	}

	return &file, nil

}
