package textstats

import (
	"strings"
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
	runeArr := []rune(text)
	return Stats{
		Bytes:       len(text),
		Runes:       len(runeArr),
		Words:       len(words),
		Frequencies: frequencies,
	}
}
