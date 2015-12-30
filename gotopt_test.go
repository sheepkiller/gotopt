package gotopt

import (
	//   "fmt"
	"errors"
	"testing"
)

// 1451436572

func TestTOTP(test *testing.T) {
	testCases := []struct {
		secret   string
		digits   int
		shaX     string
		ts       uint64
		e_token  string
		e_remain uint64
		e_err    error
		success  bool
	}{
		{"AAAAAAAAAAAAAAAA", 6, "sha1", 1451472133, "197007", 17, nil, true},
		{"AAAAAAAAAAAAAAAA", 8, "sha1", 1451472133, "82197007", 17, nil, true},
		{"AAAAAAAAAAAAAAAA", 5, "sha1", 1451472133, "97007", 17, nil, true},
		{"AAAAAAAAAAAAAAAA", 5, "sha1", 1451472133, "12345", 17, nil, false},
		{"AAAAAAAAAAAAAAAA", 5, "sha1", 1451472133, "97007", 4, nil, false},
		{"AAAAAAAAAAAAAAAA", 3, "sha1", 1451472133, "97007", 4, nil, false},
		{"AAAAAAAAAAAAAAA", 6, "sha1", 1451472133, "97007", 17, errors.New("here we go"), false},
	}
	for i := range testCases {
		tc := testCases[i]
		t, err := newTOPT(tc.secret, tc.digits, tc.shaX)
		if err != nil {
			if tc.success != false {
				test.Errorf("Expected '%s' got '%s'\n", tc.e_err, err)
			}
			continue
		}
		t.ts = tc.ts
		t_token, t_remain, t_err := t.TOPT()
		if t_token != tc.e_token && tc.success != false {
			test.Errorf("Token : expected %s got %s for %#v\n", tc.e_token, t_token, tc)
		}
		if t_remain != tc.e_remain && tc.success != false {
			test.Errorf("Remain: expected %d got %d %#v\n", tc.e_remain, t_remain, tc)
		}
		if t_err != tc.e_err && tc.success != false {
			test.Errorf("Error: expected %s got %s for %#v\n", tc.e_err, t_err, tc)
		}
	}
}
