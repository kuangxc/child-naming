package split

import (
	"regexp"
	"strings"

	"github.com/yanyiwu/gojieba"
)

func splitSetence(s string) (result []string) {
	reg := regexp.MustCompile(`(；|。|，)`)
	ss := reg.ReplaceAllString(s, `;`)
	result = strings.Split(ss, ";")
	return result
}

func SplitPoetry(s string) (words []string, err error) {
	x := gojieba.NewJieba()
	defer x.Free()
	ss := splitSetence(s)
	for _, s := range ss {
		ws := x.Cut(s, true)
		//	fmt.Println(words)
		words = append(words, ws...)
	}
	return words, nil
}
