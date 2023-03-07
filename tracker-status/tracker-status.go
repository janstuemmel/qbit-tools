package trackerstatus

import (
	"qbit-tools/qbit"
	"qbit-tools/util"
	"strings"
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

	var hashesErr []string
	var hashesOk []string
	tblOk := util.NewTable("Torrents with tracker working again", "Save path")
	tblErr := util.NewTable("Torrents with errored tracker", "Save path")

	for _, torrent := range torrents {
		if torrent.Tracker == "" {
			tblErr.AddRow(torrent.Name, torrent.SavePath)
			hashesErr = append(hashesErr, torrent.Hash)
		}

		if torrent.Tracker != "" && strings.Contains(torrent.Tags, "tracker-error") {
			tblOk.AddRow(torrent.Name, torrent.SavePath)
			hashesOk = append(hashesOk, torrent.Hash)
		}
	}

	if len(hashesErr) > 0 {
		tblErr.Print()
	}

	if len(hashesOk) > 0 {
		tblOk.Print()
	}

	err = qbit.PostTorrentsAddTags(hashesErr, "tracker-error")
	err = qbit.PostTorrentsRemoveTags(hashesOk, "tracker-error")

	if err != nil {
		return err
	}

	return nil
}
