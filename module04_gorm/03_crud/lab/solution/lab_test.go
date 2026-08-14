package solution

import "testing"

func TestStockUpdate(t *testing.T) {
	if value := StockUpdate(0)["stock"]; value != 0 {
		t.Fatalf("stock=%v", value)
	}
}
