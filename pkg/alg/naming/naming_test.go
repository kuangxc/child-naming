package naming

import (
	"testing"
	"time"

	"bou.ke/monkey"

	"github.com/kuangxc/child-naming/pkg/data"
	split "github.com/kuangxc/child-naming/pkg/alg/splitword"
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

	// Make tokenization deterministic for test.
	splitGuard := monkey.Patch(split.SplitPoetry, func(string) ([]string, error) {
		// Only 2-rune words will be used as given name candidates.
		return []string{"崇文", "龙图", "金绳", "凤篆", "坟典", "岩"}, nil
	})
	defer splitGuard.Restore()

	names := GenerateNames("邝")

	// 4 paragraphs * 5 valid(2-rune) words = 20 results.
	want := 20
	got := 0
	timeout := time.After(2 * time.Second)
	for got < want {
		select {
		case n := <-names:
			if n == nil {
				t.Fatalf("unexpected nil NameInfo")
			}
			if len([]rune(n.Name)) != 3 {
				t.Fatalf("expect full name 3 runes, got %q", n.Name)
			}
			got++
		case <-timeout:
			t.Fatalf("timeout waiting names, got=%d want=%d", got, want)
		}
	}
}
