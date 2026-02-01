package logs

import "strings"

// Application identifies the application emitting the given log.
func Application(log string) string {
	// panic("Please implement the Application() function")
    for _,r := range log{
		switch r{
		case '❗':
			return "recommendation"
		case '🔍':
			return "search"
		case '☀':
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	// panic("Please implement the Replace() function")
    return strings.ReplaceAll(log,string(oldRune),string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	// panic("Please implement the WithinLimit() function")
    cnt := 0
    for range log{
        cnt++
    }
    return cnt <= limit
}
