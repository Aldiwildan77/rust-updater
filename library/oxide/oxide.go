package oxide

import (
	"github.com/monaco-io/request"
)

type OxideProto interface {
	GetOxideRelease() (response OxideResponse, err error)
}

type oxide struct{}

func NewOxide() OxideProto {
	return &oxide{}
}

func (rr *oxide) GetOxideRelease() (response OxideResponse, err error) {
	client := request.Client{
		URL:    "https://api.github.com/repos/OxideMod/Oxide.Rust/releases/latest",
		Method: "GET",
	}

	err = client.Send().Scan(&response).Error()

	return
}
