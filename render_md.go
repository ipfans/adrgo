package adrgo

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed templates/adr.md
var mdTemplate string

func CreateADRTemplate(body Body) (out string) {
	return fmt.Sprintf(
		mdTemplate,
		body.Index,
		body.Title,
		T("LastModified"),
		time.Now().Format(T("timeformat")),
		T("Status"),
		time.Now().Format(T("timeformat")),
		T("Proposal"),
		T("Context"),
		T("Decision"),
		T("Consequences"),
	)
}
