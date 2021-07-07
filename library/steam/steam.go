package steam

import (
	"encoding/json"

	"github.com/monaco-io/request"
)

const URL = "http://api.steampowered.com/ISteamNews/GetNewsForApp/v0002"

type SteamProto interface {
	GetSteamNews(req SteamRequest) (response SteamResponse, err error)
}

type steam struct{}

func NewSteam() SteamProto {
	return &steam{}
}

func (rr *steam) GetSteamNews(req SteamRequest) (response SteamResponse, err error) {
	var reqInf map[string]string

	inrec, _ := json.Marshal(req)
	json.Unmarshal(inrec, &reqInf)

	client := request.Client{
		URL:    URL,
		Method: "GET",
		Query:  reqInf,
		Header: map[string]string{"Content-Type": "application/json"},
	}

	err = client.Send().Scan(&response).Error()

	return
}
