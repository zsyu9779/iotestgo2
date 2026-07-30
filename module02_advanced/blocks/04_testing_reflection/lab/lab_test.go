package lab

import "testing"

func TestReadFieldName(t *testing.T) {
	got, err := ReadFieldName(User{ID: 1, Name: "Alice"}, "Name")
	if err != nil || got != "Alice" {
		t.Fatalf("ReadFieldName() = %q, %v; want Alice, nil", got, err)
	}
}

func TestReadFieldNameRejectsMissingField(t *testing.T) {
	if _, err := ReadFieldName(User{}, "Missing"); err == nil {
		t.Fatal("ReadFieldName() error = nil, want an error")
	}
}
