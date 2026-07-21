package scorebook

import "testing"

func TestScorebookDemo(t *testing.T) {
	book := New()
	alice, err := book.Add("Alice", 90)
	if err != nil {
		t.Fatal(err)
	}
	bob, err := book.Add("Bob", 70)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("添加学生: %#v, %#v", alice, bob)

	if err := book.UpdateScore(bob.ID, 80); err != nil {
		t.Fatal(err)
	}
	t.Logf("更新 Bob 成绩: 70 -> 80")

	average, err := book.Average()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("平均分: %.1f", average)

	counts := book.CountByGrade()
	t.Logf("等级统计: A=%d B=%d C=%d D=%d F=%d", counts["A"], counts["B"], counts["C"], counts["D"], counts["F"])
}
