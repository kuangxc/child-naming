package naming

import (
	"fmt"

	"github.com/kuangxc/child-naming/pkg/types"
)

func getWordPinyin(word string, ps []*types.Pinyin) (*types.Pinyin, error) {
	for _, p := range ps {
		if word == p.Word {
			return p, nil
		}
	}
	return nil, fmt.Errorf("%s pinyin not find.", word)
}
