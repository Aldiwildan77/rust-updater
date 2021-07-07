package domains

import (
	"strconv"
	"strings"

	"github.com/Aldiwildan77/rust-notifier-api/api/gateways/presenters"
	"github.com/Aldiwildan77/rust-notifier-api/library/oxide"
	"github.com/Aldiwildan77/rust-notifier-api/library/rust"
	"github.com/Aldiwildan77/rust-notifier-api/library/steam"
)

type UsecaseProto interface {
	GetListPayment() (res presenters.UpdaterResponse, err error)
}

type usecase struct{}

func NewUsecase() UsecaseProto {
	return &usecase{}
}

func (uc *usecase) GetListPayment() (res presenters.UpdaterResponse, err error) {
	req := steam.SteamRequest{
		Appid:     "252490",
		Count:     "1",
		MaxLength: "300",
		Format:    "json",
	}

	snResp, err := steam.NewSteam().GetSteamNews(req)
	if err != nil {
		return presenters.UpdaterResponse{}, err
	}

	oResp, err := oxide.NewOxide().GetOxideRelease()
	if err != nil {
		return presenters.UpdaterResponse{}, err
	}

	rResp, err := rust.NewRust().GetWhenIsUpdate()
	if err != nil {
		return presenters.UpdaterResponse{}, err
	}

	res.Client = rResp.Updates[len(rResp.Updates)-1].BuildID
	res.Server = rResp.ServerUpdates[len(rResp.ServerUpdates)-1].BuildID
	res.Staging = rResp.StagingUpdates[len(rResp.StagingUpdates)-1].BuildID

	oName, err := strconv.Atoi(strings.ReplaceAll(oResp.Name, ".", ""))
	res.Oxide = int(oName)

	res.Devblog = snResp.Appnews.Newsitems[len(snResp.Appnews.Newsitems)-1].Date

	return
}
