package taskmanager

import "testing"

func TestManagerDemo(t *testing.T) {
	m := NewManager()
	first, err := m.Add("write tests")
	if err != nil {
		t.Fatal(err)
	}
	second, err := m.Add("review code")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("添加任务后: %#v", m.List())

	completed, err := m.Complete(first.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("完成任务: %#v", completed)
	t.Logf("完成后列表: %#v", m.List())

	if err := m.Delete(second.ID); err != nil {
		t.Fatal(err)
	}
	t.Logf("删除第二个任务后: %#v", m.List())
}
