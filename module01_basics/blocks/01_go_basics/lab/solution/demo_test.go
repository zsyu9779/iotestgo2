package grade

import "testing"

func TestGradeDemo(t *testing.T) {
	for _, score := range []int{95, 82, 75, 60, 42} {
		got, err := Grade(score)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("分数 %d -> 等级 %s", score, got)
	}
}
