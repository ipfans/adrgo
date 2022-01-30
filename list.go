package adrgo

import (
	"context"
	"io"
	"io/fs"
	"path"
	"path/filepath"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func listRecords(ctx context.Context, root string) (records []ADRecord) {
	_ = filepath.Walk(root, func(fp string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		_, fn := path.Split(fp)
		if !validFileName(fn) {
			return nil
		}
		var record ADRecord
		record.ID, record.Title = extractInfo(fn)
		body, err := ParserFile(ctx, fp)
		if err != nil {
			return err
		}
		record.Status = body.Status[len(body.Status)-1].Value
		record.LastModified = body.LastUpdate
		records = append(records, record)
		return nil
	})
	return
}

func List(ctx context.Context, conf Config, w io.Writer) (err error) {
	records := listRecords(ctx, conf.Path)
	writer := tablewriter.NewWriter(w)
	writer.SetHeader([]string{
		T("ID"),
		T("Title"),
		T("Status"),
		T("LastModified"),
	})
	for _, record := range records {
		writer.Append([]string{
			strconv.Itoa(record.ID),
			record.Title,
			record.Status,
			record.LastModified,
		})
	}
	writer.Render()
	return
}
