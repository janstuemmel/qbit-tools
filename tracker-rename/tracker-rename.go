package trackerrename

import (
	"net/url"
	"qbit-tools/qbit"
)

type TrackerrenameCmd struct {
	From *url.URL `required:"" name:"from" help:"Existing tracker url" type:"url"`
	To   *url.URL `required:"" name:"to" help:"New tracker url" type:"url"`
}

func (c *TrackerrenameCmd) Run() error {
	qbit, err := qbit.NewClient()

	if err != nil {
		return err
	}

	torrents, err := qbit.GetTorrentsInfo()

	if err != nil {
		return err
	}

	for _, torrent := range torrents {
		payload := map[string]string{
			"hash":    torrent.Hash,
			"origUrl": c.From.String(),
			"newUrl":  c.To.String(),
		}

		// ignore all errors here
		qbit.Cli.R().
			SetFormData(payload).
			Post("/api/v2/torrents/editTracker")
	}

	return nil
}
