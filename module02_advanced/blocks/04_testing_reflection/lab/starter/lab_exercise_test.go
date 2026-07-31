//go:build exercise

package lab

import "testing"

func TestReadFieldName(t *testing.T) {
	tests := []struct {
		name      string
		fieldName string
		want      string
		wantErr   bool
	}{
		{name: "string field", fieldName: "Name", want: "Alice"},
		{name: "non-string field", fieldName: "ID", wantErr: true},
		{name: "missing field", fieldName: "Missing", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFieldName(User{ID: 1, Name: "Alice"}, tt.fieldName)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ReadFieldName() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Fatalf("ReadFieldName() = %q, want %q", got, tt.want)
			}
		})
	}
}
