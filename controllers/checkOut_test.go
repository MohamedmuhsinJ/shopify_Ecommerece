package controllers

import "testing"

func TestOID(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"muhsin", len(OID())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := OID()
			if len(got) != tt.want {
				t.Errorf("OID() = %v, want %v", got, tt.want)
			}
		})
	}
}
