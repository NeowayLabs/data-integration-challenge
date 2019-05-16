package util

import "testing"

func TestIsBlank(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"Empty", "", true},
		{"Blank", " ", true},
		{"Tab", "\t", true},
		{"Non empty", "a", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBlank(tt.arg); got != tt.want {
				t.Errorf("IsBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}
