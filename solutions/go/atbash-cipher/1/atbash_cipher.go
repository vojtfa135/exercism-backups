package atbash

import (
	"fmt"
	"slices"
	"sort"
	"strings"
	"unicode"
)

func alphabetMapping() []string {
	return []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
}

func reverseAlphabetMapping() []string {
	alphabetMapping := alphabetMapping()
	sort.Sort(sort.Reverse(sort.StringSlice(alphabetMapping)))

	return alphabetMapping
}

type AtbashCipher struct {
	alphabetMapping, reverseAlphabetMapping []string
	groupSize                               int
}

func newAtbashCipher() AtbashCipher {
	return AtbashCipher{
		alphabetMapping:        alphabetMapping(),
		reverseAlphabetMapping: reverseAlphabetMapping(),
		groupSize:              5,
	}
}

func (ab *AtbashCipher) normalizePlaintText(plainText string) string {
	var (
		normalizedPlaintText strings.Builder
		stripped             = strings.ReplaceAll(strings.ToLower(plainText), " ", "")
	)

	for _, r := range stripped {
		if !unicode.IsPunct(r) {
			normalizedPlaintText.WriteRune(r)
		}
	}

	return normalizedPlaintText.String()
}

func (ab *AtbashCipher) encrypt(plainText string) string {
	var (
		cipherText, normalizedChar, encChar string
		charAlphabetIdx                     int
		normalizedPlaintText                = ab.normalizePlaintText(plainText)
	)

	for idx, char := range normalizedPlaintText {
		normalizedChar = fmt.Sprintf("%c", char)
		charAlphabetIdx = slices.Index(ab.alphabetMapping, normalizedChar)

		if charAlphabetIdx != -1 {
			encChar = ab.reverseAlphabetMapping[charAlphabetIdx]
		} else {
			encChar = normalizedChar
		}

		cipherText += encChar

		if newIdx := idx + 1; newIdx%ab.groupSize == 0 && newIdx != len(normalizedPlaintText) {
			cipherText += " "
		}
	}

	return cipherText
}

func Atbash(s string) string {
	atbash := newAtbashCipher()

	return atbash.encrypt(s)
}
