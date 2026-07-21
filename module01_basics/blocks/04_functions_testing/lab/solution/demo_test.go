package scores

import "testing"

func TestScoresDemo(t *testing.T) {
	input := []int{59, 60, 75, 90}
	filtered := Filter(input, AtLeast(60))
	t.Logf("输入成绩: %v", input)
	t.Logf("筛选条件: 至少 60 分")
	t.Logf("筛选结果: %v", filtered)

	var events []string
	WithAudit("average", func(event string) {
		events = append(events, event)
	}, func() {
		events = append(events, "operation")
	})
	t.Logf("审计事件: %v", events)
}
