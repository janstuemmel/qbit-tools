package main

import (
	backup "qbit-tools/backup"
	trackerrename "qbit-tools/tracker-rename"
	trackerstatus "qbit-tools/tracker-status"

	"github.com/alecthomas/kong"
)

var cli struct {
	Trackerstatus trackerstatus.TrackerstatusCmd `cmd:"" name:"tracker-status" help:"Label tracker errors"`
	Backup        backup.BackupCmd               `cmd:"" name:"backup" help:"Backup torrents as zip"`
	Trackerrename trackerrename.TrackerrenameCmd `cmd:"" name:"tracker-rename" help:"Batch rename tracker urls"`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
