package classroomdb

import (
	"testing"
)

func TestDSNUsesClassroomMySQL(t *testing.T) {
	want := "root:password@tcp(127.0.0.1:3306)/gorm_demo?charset=utf8mb4&parseTime=True&loc=Local"
	if DSN != want {
		t.Fatalf("DSN = %q, want %q", DSN, want)
	}
}
