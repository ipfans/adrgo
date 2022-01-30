package adrgo

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	fnPattern    = regexp.MustCompile(`(?mi)^(\d+)-(.+)\.md`)
	titlePattern = regexp.MustCompile(`(?mi)^(\d+)\.(.+)$`)
)

// pathExists checks if the path exists.
func pathExists(p string) bool {
	_, err := os.Stat(p)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// validFileName validates the file name with id and name.
func validFileName(fn string) bool {
	return fnPattern.MatchString(fn)
}

// extractInfo extracts the id and name from the file name.
func extractInfo(fn string) (int, string) {
	s := fnPattern.FindAllStringSubmatch(fn, -1)
	id, _ := strconv.Atoi(s[0][1])
	return id, s[0][2]
}

// headerInfo formats the header information.
func HeaderInfo(title string) (id int, name string) {
	title = strings.TrimSpace(title)
	s := titlePattern.FindAllStringSubmatch(title, -1)
	if len(s) == 0 {
		return 0, ""
	}
	id, _ = strconv.Atoi(s[0][1])
	name = strings.TrimSpace(s[0][2])
	return
}

func today() string {
	return time.Now().Format(T("timeformat"))
}

func filename(r ADRecord, digits int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(digits)+"d-%s.md", r.ID, r.Title)
}
