package backup

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"qbit-tools/qbit"
)

type BackupCmd struct {
	Path string `optional:"" name:"path" help:"Path to backup" type:"path" default:"."`
	File string `optional:"" name:"file" help:"File name for zip archive" type:"string" default:"backup.zip"`
}

func (c *BackupCmd) Run() error {
	archive, err := os.Create(path.Join(c.Path, c.File))
	defer archive.Close()

	if err != nil {
		return err
	}

	zipWriter := zip.NewWriter(archive)

	qbit, err := qbit.NewClient()

	if err != nil {
		return err
	}

	torrents, err := qbit.GetTorrentsInfo()

	for _, torrent := range torrents {

		res, err := qbit.Cli.R().
			SetQueryParam("hash", torrent.Hash).
			Get("api/v2/torrents/export")

		if err != nil {
			fmt.Println(err)
			continue
		}

		file := path.Join(torrent.SavePath, torrent.Name+".torrent")
		reader := bytes.NewReader(res.Bytes())
		zipFile, err := zipWriter.Create(file)

		if _, err := io.Copy(zipFile, reader); err != nil {
			fmt.Println(err)
			continue
		}
	}

	zipWriter.Close()
	return nil
}
