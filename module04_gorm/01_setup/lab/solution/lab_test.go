package solution

import "testing"

func TestValidatePool(t *testing.T) {
	if ValidatePool(10, 5) != nil || ValidatePool(2, 3) == nil || ValidatePool(0, 0) == nil {
		t.Fatal("pool validation contract failed")
	}
}
