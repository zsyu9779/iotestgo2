package scores

func Filter(values []int, keep func(int) bool) []int {
	return nil
}

func AtLeast(min int) func(int) bool {
	return func(value int) bool { return false }
}

func WithAudit(name string, record func(string), operation func()) {
	operation()
}
