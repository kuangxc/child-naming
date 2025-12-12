package naming

import (
	"testing"

	"bou.ke/monkey"

	"github.com/kuangxc/child-naming/pkg/data"
	"github.com/kuangxc/child-naming/pkg/types"
)

func TestGenerateNames(t *testing.T) {
	dataGuard := monkey.Patch(data.ParsePoetryData, func([]string) (map[string][]*types.Poetry, error) {
		return map[string][]*types.Poetry{
			"tang": {
				&types.Poetry{
					Category: "tang",
					Title:    "帝京篇十首 二",
					Author:   "太宗皇帝",
					Paragraphs: []string{
						"岩廊罢机务，崇文聊驻辇。",
						"玉匣启龙图，金绳披凤篆。",
						"韦编断仍续，缥帙舒还卷。",
						"对此乃淹留，欹案观坟典。",
					},
				},
			},
		}, nil
	})
	defer dataGuard.Restore()
	names := GenerateNames("邝")
	if len(names) != 11 {
		t.Errorf("expect 5 names,now:%d", len(names))
	}
}
