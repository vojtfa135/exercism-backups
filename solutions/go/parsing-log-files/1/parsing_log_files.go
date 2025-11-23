package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[TRC\])|^(\[DBG\])|^(\[INF\])|^(\[WRN\])|^(\[ERR\])|^(\[FTL\])`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]{0,}>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`\".*(p|P)(a|A)(s|S)(s|S)(w|W)(o|O)(r|R)(d|D)\"`)
	matches := []bool{}

	for _, val := range lines {
		if isMatch := re.MatchString(val); isMatch {
			matches = append(matches, isMatch)
		}
	}

	return len(matches)
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d{1,}`)

	if re.MatchString(text) {
		return re.ReplaceAllString(text, "")
	}

	return text
}

func TagWithUserName(lines []string) []string {
	var (
		taggedLogs      []string
		re              = regexp.MustCompile(`User {1,}\w*`)
		tag             = "[USR]"
		name, taggedLog string
	)

	for _, val := range lines {
		taggedLog = val

		if re.MatchString(val) {
			name = re.FindString(val)
			name = strings.ReplaceAll(name, "User", "")
			name = strings.ReplaceAll(name, " ", "")
			taggedLog = fmt.Sprintf("%s %s %s", tag, name, val)
		}

		taggedLogs = append(taggedLogs, taggedLog)
	}

	return taggedLogs
}
