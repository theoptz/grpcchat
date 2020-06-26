package utils

import "testing"

func TestGetHash(t *testing.T) {
	cases := []struct {
		v   string
		res string
	}{
		{
			v:   "example",
			res: "50d858e0985ecc7f60418aaf0cc5ab587f42c2570a884095a9e8ccacd0f6545c",
		},
	}

	for _, tc := range cases {
		res := GetHash(tc.v)
		if res != tc.res {
			t.Errorf("Hash of %s must be %s but got %s", tc.v, tc.res, res)
		}
	}
}
