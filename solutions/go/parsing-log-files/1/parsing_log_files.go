package parsinglogfiles

import (
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	// panic("Please implement the IsValidLine function")
	valid := []string{"[TRC]", "[DBG]", "[INF]", "[WRN]", "[ERR]", "[FTL]"}
	for _, v := range valid {
		if strings.HasPrefix(text, v) {
			return true
		}
	}
	return false
}

func SplitLogLine(text string) []string {
	// panic("Please implement the SplitLogLine function")
	re := regexp.MustCompile(`<[~*=\-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	// panic("Please implement the CountQuotedPasswords function")
	count := 0
	re := regexp.MustCompile(`"[^"]*(?i)password[^"]*"`)
	for _, line := range lines {
		if re.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	// panic("Please implement the RemoveEndOfLineText function")
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	// panic("Please implement the TagWithUserName function")
	out := make([]string, len(lines))
	re := regexp.MustCompile(`User\s+(\S+)`)
	for i, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) > 1 {
			out[i] = "[USR] " + matches[1] + " " + line
		} else {
			out[i] = line
		}
	}
	return out
}
