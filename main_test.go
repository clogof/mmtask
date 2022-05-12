package main

import (
	"testing"
)

func TestFindMaxSubstr(t *testing.T) {
	testCase := map[string]struct {
		Val int
		Err error
	}{
		"abcdef":         {Val: 2, Err: nil},
		"aaaaa":          {Val: 5, Err: nil},
		"abaccab":        {Val: 4, Err: nil},
		"":               {Val: 0, Err: nil},
		"ababababa":      {Val: 9, Err: nil},
		"ababcababa":     {Val: 5, Err: nil},
		"caaacbbbczzzzc": {Val: 6, Err: nil},
		"ab cd":          {Val: 0, Err: myError("can not use space characters")},
		"   ":            {Val: 0, Err: myError("can not use space characters")},
	}

	for k, v := range testCase {
		val, err := FindMaxSubstr(k)
		if val != v.Val || err != v.Err {
			t.Errorf("Error FindMaxSubstr(\"%s\"): Val: %d Err: %v\n", k, val, err)
		}
	}
}
