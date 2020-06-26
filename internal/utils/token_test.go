package utils

import "testing"

func TestGenerateToken(t *testing.T) {
	v := GenerateToken("example")

	if len(v) != 64 {
		t.Errorf("Invalid token")
	}
}
