package main

import (
	export "qbit-tools/backup"
	trackerstatus "qbit-tools/tracker-status"

	"github.com/alecthomas/kong"
)

var cli struct {
	Trackerstatus trackerstatus.TrackerstatusCmd `cmd:"" name:"tracker-status" help:"Label tracker errors"`
	Export        export.BackupCmd               `cmd:"" name:"backup" help:"Backup torrents as zip"`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
