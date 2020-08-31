package naming

import (
	"fmt"

	split "github.com/kuangxc/child-naming/pkg/alg/splitword"
	"github.com/kuangxc/child-naming/pkg/conf"
	"github.com/kuangxc/child-naming/pkg/data"
	"github.com/kuangxc/child-naming/pkg/types"
)

func GenerateNames(surname string) chan *types.NameInfo {
	resultNameChn := make(chan *types.NameInfo, 100)
	poetryMap, err := data.ParsePoetryData(conf.GetExcludeDirs())
	if nil != err {
		fmt.Println("parse poetry data err:", err)
		return nil
	}
	go func() {
		for _, ps := range poetryMap {
			for _, p := range ps {
				for _, pa := range p.Paragraphs {
					words, err := split.SplitPoetry(pa)
					if nil != err {
						fmt.Println("split poetry err:", err)
					}
					for _, w := range words {
						cs := []rune(w)
						if 2 == len(cs) {
							resultNameChn <- &types.NameInfo{
								Name:    surname + w,
								From:    p.Category + " " + p.Title + " " + p.Author,
								Meaning: pa,
							}
						}
					}
				}
			}
		}
	}()

	return resultNameChn
}
