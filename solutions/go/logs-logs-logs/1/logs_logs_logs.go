package logs

import (
	"maps"
	"slices"
	"strings"
	"unicode/utf8"
)

const (
	recomm = "recommendation"
	srch   = "search"
	wthr   = "weather"
	deft   = "default"
)

var logChars = map[string]rune{
	recomm: '\u2757',
	srch:   '\U0001F50D',
	wthr:   '\u2600',
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	emittedApps := map[int]string{}
	runeIdx := -1
	emittedApp := deft

	if strings.ContainsRune(log, logChars[recomm]) {
		runeIdx = strings.IndexRune(log, logChars[recomm])
		emittedApps[runeIdx] = recomm
	}

	if strings.ContainsRune(log, logChars[srch]) {
		runeIdx = strings.IndexRune(log, logChars[srch])
		emittedApps[runeIdx] = srch
	}

	if strings.ContainsRune(log, logChars[wthr]) {
		runeIdx = strings.IndexRune(log, logChars[wthr])
		emittedApps[runeIdx] = wthr
	}

	if len(emittedApps) > 0 {
		sortedEmition := slices.Sorted(maps.Keys(emittedApps))
		emittedApp = emittedApps[sortedEmition[0]]
	}

	return emittedApp
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
