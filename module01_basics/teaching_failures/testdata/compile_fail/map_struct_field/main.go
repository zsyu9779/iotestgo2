package mapstructfield

type Item struct {
	Value int
}

func demonstrate() {
	items := map[string]Item{"one": {Value: 1}}
	items["one"].Value = 2
}
