package naming

import (
	"fmt"

	"github.com/kuangxc/child-naming/pkg/types"
	"github.com/kuangxc/child-naming/pkg/util"
)

func filterExclude(srcNameChn chan *types.NameInfo, excludeWords []string,
	excludeNames []string) chan *types.NameInfo {
	filterExcludeNameChn := make(chan *types.NameInfo, 100)
	go func() {
		for {
			n, ok := <-srcNameChn
			if !ok {
				break
			}
			if util.IncludeStr(excludeNames, n.Name) {
				continue
			}
			rs := []rune(n.Name)
			i := 1
			for ; i < 3; i++ {
				if util.IncludeStr(excludeWords, string(rs[i])) {
					break
				}
			}
			if i >= 3 {
				filterExcludeNameChn <- n
			}
		}
	}()
	return filterExcludeNameChn
}

func filterInclude(srcNameChn chan *types.NameInfo, includeWords []string) chan *types.NameInfo {
	filterIncludeNameChn := make(chan *types.NameInfo, 100)
	go func() {
		for {
			n, ok := <-srcNameChn
			if !ok {
				break
			}
			rs := []rune(n.Name)
			i := 1
			for ; i < 3; i++ {
				if !util.IncludeStr(includeWords, string(rs[i])) {
					break
				}
			}
			if i >= 3 {
				filterIncludeNameChn <- n
			}
		}
	}()
	return filterIncludeNameChn
}

//音调相关的问题
func filtedPinyinDisharmony(srcNameChn chan *types.NameInfo, ps []*types.Pinyin) chan *types.NameInfo {
	resultNameChn := make(chan *types.NameInfo, 100)
	go func() {
		for {
			n, ok := <-srcNameChn
			if !ok {
				break
			}
			var pys []*types.Pinyin
			rs := []rune(n.Name)
			for i := 0; i < len(rs); i++ {
				py, err := getWordPinyin(string(rs[i]), ps)
				if nil != err {
					fmt.Printf("%s get pinyin err:%v \r\n", n.Name, err)
					break
				}
				pys = append(pys, py)
			}
			if 3 != len(pys) {
				continue
			}
			//声母相同，过滤
			if util.GetInitial(pys[0].Pinyin) == util.GetInitial(pys[1].Pinyin) ||
				util.GetInitial(pys[1].Pinyin) == util.GetInitial(pys[2].Pinyin) {
				continue
			}
			//韵母相同，过滤
			if util.GetVowel(pys[0].Pinyin) == util.GetVowel(pys[1].Pinyin) ||
				util.GetVowel(pys[1].Pinyin) == util.GetVowel(pys[2].Pinyin) {
				continue
			}
			//音调相同，过滤
			if pys[0].Tone == pys[1].Tone || pys[1].Tone == pys[2].Tone {
				continue
			}

			resultNameChn <- n
		}
	}()

	return resultNameChn
}
