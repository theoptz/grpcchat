package utils

import "testing"

import "strings"

func TestNewRandomIntString(t *testing.T) {
	cases := []int{-1, 0, 1, 10}

	for _, l := range cases {
		v := NewRandomIntString(l)

		if l < 1 {
			if v != "" {
				t.Errorf("Res must be empty but got %s", v)
			}
			continue
		}

		if len(v) != l {
			t.Errorf("Len must be %d but got %d", l, len(v))
			continue
		}

		for i := 0; i < l; i++ {
			if !strings.Contains(digits, string(v[i])) {
				t.Errorf("Res must be contains only %s but have %s", digits, string(v[i]))
				break
			}
		}
	}
}
