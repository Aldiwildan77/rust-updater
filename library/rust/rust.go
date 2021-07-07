package rust

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type RustProto interface {
	GetWhenIsUpdate() (response RustResponse, err error)
}

type rust struct{}

func NewRust() RustProto {
	return &rust{}
}

func (rr *rust) GetWhenIsUpdate() (response RustResponse, err error) {
	var resp *http.Response

	resp, err = http.Get("https://whenisupdate.com/api.json")
	if err != nil {
		return
	}

	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	respByte := buf.Bytes()

	err = json.Unmarshal(respByte, &response)

	return
}
