package domains

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Aldiwildan77/rust-notifier-api/api/gateways/presenters"
	"github.com/Aldiwildan77/rust-notifier-api/library"
	"github.com/Aldiwildan77/rust-notifier-api/library/oxide"
	"github.com/Aldiwildan77/rust-notifier-api/library/rust"
	"github.com/Aldiwildan77/rust-notifier-api/library/steam"
)

type UsecaseProto interface {
	GetListUpdates() (res presenters.UpdaterResponse, err error)
}

type usecase struct {
	Ctx context.Context
}

func NewUsecase(ctx context.Context) UsecaseProto {
	return &usecase{
		Ctx: ctx,
	}
}

func (uc *usecase) GetListUpdates() (res presenters.UpdaterResponse, err error) {
	// get update result async
	if err := <-getResult(uc.Ctx, &res); err != nil {
		log.Println("result error: ", err)
		return presenters.UpdaterResponse{}, err
	}

	return
}

func getResult(ctx context.Context, res *presenters.UpdaterResponse) <-chan error {
	timeout := 2 * time.Second
	out := make(chan error)

	go func() {
		defer close(out)

		ctx, cancel := context.WithTimeout(ctx, timeout)

		defer cancel()

		rResp := new(rust.RustResponse)
		oResp := new(oxide.OxideResponse)
		stResp := new(steam.SteamResponse)

		for err := range library.Merge(
			fetch(ctx, "https://whenisupdate.com/api.json", rResp),
			fetch(ctx, "https://api.github.com/repos/OxideMod/Oxide.Rust/releases/latest", oResp),
			fetch(ctx, "https://api.steampowered.com/ISteamNews/GetNewsForApp/v0002/?appid=252490&count=1&maxlength=300&format=json", stResp),
		) {
			if err != nil {
				out <- err
				return
			}
		}

		// merge the result
		oName, err := strconv.Atoi(strings.ReplaceAll(oResp.Name, ".", ""))
		if err != nil {
			out <- err
			return
		}

		res.Client = rResp.Updates[len(rResp.Updates)-1].BuildID
		res.Server = rResp.ServerUpdates[len(rResp.ServerUpdates)-1].BuildID
		res.Staging = rResp.StagingUpdates[len(rResp.StagingUpdates)-1].BuildID
		res.Devblog = stResp.Appnews.Newsitems[len(stResp.Appnews.Newsitems)-1].Date
		res.Oxide = int(oName)
	}()
	return out
}

func fetch(ctx context.Context, url string, result interface{}) <-chan error {
	out := make(chan error)
	go func() {
		defer close(out)

		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			out <- err
			return
		}

		rsp, err := http.DefaultClient.Do(req)
		if err != nil {
			out <- err
			return
		} else if rsp.StatusCode != http.StatusOK {
			out <- fmt.Errorf("%d: %s", rsp.StatusCode, rsp.Status)
			return
		}

		defer rsp.Body.Close()
		if err := json.NewDecoder(rsp.Body).Decode(result); err != nil {
			out <- err
			return
		}
	}()
	return out
}
