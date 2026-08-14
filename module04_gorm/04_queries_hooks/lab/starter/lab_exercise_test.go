//go:build exercise

package starter

import "testing"

func TestAvoidsNPlusOne(t *testing.T) {
	if !AvoidsNPlusOne(10, 2) || AvoidsNPlusOne(10, 11) {
		t.Fatal("Preload contract should accept two queries, not 1+N")
	}
}
