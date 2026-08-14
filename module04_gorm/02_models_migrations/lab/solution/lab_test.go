package solution

import "testing"

func TestValidateRelation(t *testing.T) {
	if ValidateRelation(Relation{1, 2, 3}) != nil || ValidateRelation(Relation{1, 0, 3}) == nil {
		t.Fatal("relation validation contract failed")
	}
}
