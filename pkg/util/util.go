package util

import (
	"fmt"
	"strings"
)

var (
	initials = []string{"b", "p", "m", "f", "d", "t", "n", "l", "g", "k",
		"h", "j", "q", "x", "zh", "ch", "sh", "r", "z", "c", "s", "y", "w"}
	vowel = []string{"a", "o", "e", "i", "u", "ü", "ai", "ei", "ui", "ao", "ou", "iu", "ie", "üe", "er",
		"an", "en", "in", "un", "ün", "ang", "eng", "ing", "ong"}
)

func IncludeStr(ss []string, candicate string) bool {
	for _, s := range ss {
		if s == candicate {
			return true
		}
	}
	return false
}

func GetInitial(pinyin string) string {
	if len(pinyin) < 2 {
		fmt.Println("illegal pinyin:", pinyin)
		return pinyin
	}
	init := ""
	for _, j := range initials {
		if strings.Contains(pinyin[0:2], j) {
			if len(j) > len(init) {
				init = j
			}
		}
	}

	return init
}
func GetVowel(pinyin string) string {
	var vw string
	for _, v := range vowel {
		if strings.Contains(pinyin, v) && len(v) > len(vw) {
			vw = v
		}
	}
	return vw
}
