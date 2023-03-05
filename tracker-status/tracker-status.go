package trackerstatus

import (
	"qbit-tools/qbit"
)

type TrackerstatusCmd struct{}

func (c *TrackerstatusCmd) Run() error {
	qbit, err := qbit.NewClient()

	if err != nil {
		return err
	}

	torrents, err := qbit.GetTorrentsInfo()

	if err != nil {
		return err
	}

	var hashes []string

	for _, torrent := range torrents {
		if torrent.Tracker == "" {
			hashes = append(hashes, torrent.Hash)
		}
	}

	err = qbit.PostTorrentsAddTags(hashes, "tracker-error")

	if err != nil {
		return err
	}

	return nil
}
