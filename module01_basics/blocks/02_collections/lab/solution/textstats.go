package textstats

import (
	"strings"
	"unicode/utf8"
)

type Stats struct {
	Bytes       int
	Runes       int
	Words       int
	Frequencies map[string]int
}

func Analyze(text string) Stats {
	words := strings.Fields(text)
	frequencies := make(map[string]int, len(words))
	for _, word := range words {
		frequencies[strings.ToLower(word)]++
	}
	return Stats{
		Bytes:       len(text),
		Runes:       utf8.RuneCountInString(text),
		Words:       len(words),
		Frequencies: frequencies,
	}
}
