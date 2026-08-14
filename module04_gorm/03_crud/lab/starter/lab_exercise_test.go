//go:build exercise

package starter

import "testing"

func TestStockUpdate(t *testing.T) {
	got := StockUpdate(0)
	if value, ok := got["stock"]; !ok || value != 0 {
		t.Fatalf("StockUpdate(0)=%v", got)
	}
}
