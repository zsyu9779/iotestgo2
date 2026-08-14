//go:build exercise

package starter

import "testing"

func TestTransfer(t *testing.T) {
	from, to, err := Transfer(100, 50, 30)
	if err != nil || from != 70 || to != 80 {
		t.Fatalf("Transfer()=%d,%d,%v", from, to, err)
	}
	from, to, err = Transfer(10, 50, 30)
	if err == nil || from != 10 || to != 50 {
		t.Fatal("failed transfer must preserve both balances")
	}
}
