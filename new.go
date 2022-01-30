package adrgo

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func New(ctx context.Context, conf Config, new ADRecord) (err error) {
	if pathExists(conf.Path) {
		err = os.MkdirAll(conf.Path, os.ModePerm)
		if err != nil {
			return
		}
	}

	var record ADRecord
	records := listRecords(ctx, conf.Path)
	if len(records) > 0 {
		sort.Slice(records, func(i, j int) bool {
			return records[i].ID > records[j].ID
		})
		record = records[0]
	} else {
		record = ADRecord{
			ID: 0,
		}
	}
	new.ID = record.ID + 1

	data := fmt.Sprintf(mdTemplate, new.ID, new.Title,
		T("LastModified"), today(),
		T("Status"), today(), T("proposed"),
		T("Context"), T("Decision"), T("Consequences"),
	)
	fn := filename(new, conf.Digits)

	return os.WriteFile(filepath.Join(conf.Path, fn), []byte(data), 0644)
}
