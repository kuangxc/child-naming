package split

import (
	"testing"
)

func TestSplitPoetry(t *testing.T) {
	s := "秦川雄帝宅，函谷壮皇居。"
	words, err := SplitPoetry(s)
	if err != nil {
		t.Errorf("expect no err,now err:%v", err)
	}
	if len(words) == 0 {
		t.Errorf("expect words size > 0,now:%d", len(words))
	}
}
