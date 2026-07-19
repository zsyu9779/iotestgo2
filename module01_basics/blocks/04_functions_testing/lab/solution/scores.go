package scores

func Filter(values []int, keep func(int) bool) []int {
	filtered := make([]int, 0, len(values))
	for _, value := range values {
		if keep(value) {
			filtered = append(filtered, value)
		}
	}
	return filtered
}

func AtLeast(min int) func(int) bool {
	return func(value int) bool { return value >= min }
}

func WithAudit(name string, record func(string), operation func()) {
	record("start:" + name)
	defer record("end:" + name)
	operation()
}
