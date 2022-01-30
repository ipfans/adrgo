package adrgo

import (
	"bytes"
	"context"
	"os"
	"regexp"

	mermaid "github.com/abhinav/goldmark-mermaid"
	"github.com/yuin/goldmark"
	emoji "github.com/yuin/goldmark-emoji"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
)

var (
	md = goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			highlighting.Highlighting,
			emoji.Emoji,
			&mermaid.Extender{},
		),
		goldmark.WithParserOptions(
			parser.WithBlockParsers(),
			parser.WithInlineParsers(),
			parser.WithAutoHeadingID(),
			parser.WithAttribute(),
		),
	)

	statusRegexp = regexp.MustCompile(`(?m)^(.*): (.*)$`)
)

type Head struct {
	Head  string
	Start int
	End   int
}

type Status struct {
	Value string
	Date  string
}

type Body struct {
	FileName   string
	Index      int
	Title      string
	LastUpdate string
	Status     []Status
	Heads      []Head
}

func ParserFile(ctx context.Context, fp string) (body Body, err error) {
	body.FileName = fp

	var b []byte
	b, err = os.ReadFile(fp)
	if err != nil {
		return
	}
	dateRegexp := regexp.MustCompile(`(?m)^` + T("LastModified") + `: (.*)$`)
	loc := dateRegexp.FindIndex(b)
	if loc != nil {
		body.LastUpdate = string(b[loc[0]+13 : loc[1]])
	}
	re := regexp.MustCompile(`(?m)^##\x20+(.*)\x20*$`)
	locs := re.FindAllIndex(b, -1)
	for i := range locs {
		body.Heads = append(
			body.Heads,
			Head{
				Head:  string(bytes.TrimSpace(b[locs[i][0]+2 : locs[i][1]])),
				Start: locs[i][0],
				End:   locs[i][1],
			},
		)
	}

	for i := range body.Heads {
		if body.Heads[i].Head == T("Status") {
			tmp := b[body.Heads[i].End+1 : body.Heads[i+1].Start-1]
			matched := statusRegexp.FindAllStringSubmatch(string(tmp), -1)
			for i := range matched {
				body.Status = append(
					body.Status,
					Status{
						Value: matched[i][2],
						Date:  matched[i][1],
					},
				)
			}
		}
	}

	return
}
