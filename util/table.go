package util

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

func NewTable(header ...interface{}) table.Table {
	headerFmt := color.New(color.FgHiMagenta, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgHiYellow).SprintfFunc()
	tbl := table.New(header...)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	return tbl
}
