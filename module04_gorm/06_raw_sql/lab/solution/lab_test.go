package solution

import "testing"

func TestFindByName(t *testing.T) {
	query, args := FindByName("Alice' OR 1=1 --")
	if query != "SELECT id, name FROM m04_l06_users WHERE name = ?" || len(args) != 1 {
		t.Fatalf("query=%q args=%v", query, args)
	}
}
