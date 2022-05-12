package main

import (
	"fmt"
)

type myError string

func (e myError) Error() string {
	return string(e)
}

func FindMaxSubstr(s string) (int, error) {
	max := 0
	// индекс левой границы искомой строки
	l := 0
	// индекс элемента, отличающегося от предыдущего
	g := 0
	m := make(map[byte]interface{})

	for i, ch := range []byte(s) {
		if ch == ' ' {
			return 0, myError("can not use space characters")
		}

		if len(m) > 1 {
			if _, ok := m[ch]; !ok {
				l = g
				for k := range m {
					if k != s[l] {
						delete(m, k)
						break
					}
				}
			}
		}

		if ch != s[g] {
			g = i
		}

		m[ch] = struct{}{}

		if c := i - l + 1; c > max {
			max = c
		}
	}

	return max, nil
}

func main() {
	s := "caaaaaacbbbczzzzzc"

	cnt, err := FindMaxSubstr(s)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(cnt)
}
