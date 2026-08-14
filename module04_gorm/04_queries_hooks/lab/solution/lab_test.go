package solution

import "testing"

func TestAvoidsNPlusOne(t *testing.T) {
	if !AvoidsNPlusOne(10, 2) || AvoidsNPlusOne(10, 11) {
		t.Fatal("N+1 contract failed")
	}
}
