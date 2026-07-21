package student

import "testing"

func TestStudentDemo(t *testing.T) {
	s, err := New(1, " Alice ", 80)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("创建学生: %#v", *s)

	if err := s.Rename(" Bob "); err != nil {
		t.Fatal(err)
	}
	t.Logf("改名后: %#v", *s)

	if err := s.UpdateScore(95); err != nil {
		t.Fatal(err)
	}
	t.Logf("更新成绩后: %#v", *s)
	t.Logf("快照副本: %#v", s.Snapshot())
}
