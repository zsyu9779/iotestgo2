//go:build exercise

package starter

import "testing"

func TestValidateRelation(t *testing.T) {
	if ValidateRelation(Relation{1, 2, 3}) != nil || ValidateRelation(Relation{1, 0, 3}) == nil {
		t.Fatal("all foreign keys must be non-zero")
	}
}
