package oxide

import (
	"github.com/monaco-io/request"
)

const URL = "https://api.github.com/repos/OxideMod/Oxide.Rust/releases/latest"

type OxideProto interface {
	GetOxideRelease() (response OxideResponse, err error)
}

type oxide struct{}

func NewOxide() OxideProto {
	return &oxide{}
}

func (rr *oxide) GetOxideRelease() (response OxideResponse, err error) {
	client := request.Client{
		URL:    URL,
		Method: "GET",
		Header: map[string]string{"Content-Type": "application/json"},
	}

	err = client.Send().Scan(&response).Error()

	return
}
