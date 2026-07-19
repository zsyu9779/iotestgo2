package textstats

type Stats struct {
	Bytes       int
	Runes       int
	Words       int
	Frequencies map[string]int
}

func Analyze(text string) Stats {
	return Stats{Frequencies: map[string]int{}}
}
