package solution

import "testing"

func TestTransfer(t *testing.T) {
	from, to, err := Transfer(100, 50, 30)
	if err != nil || from != 70 || to != 80 {
		t.Fatal("successful transfer contract failed")
	}
	from, to, err = Transfer(10, 50, 30)
	if err == nil || from != 10 || to != 50 {
		t.Fatal("rollback contract failed")
	}
}
