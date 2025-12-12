package split

import (
	"testing"
)

func TestSplitPoetry(t *testing.T) {
	s := "秦川雄帝宅，函谷壮皇居。"
	words, err := SplitPoetry(s)
	if nil != err || len(words) != 5 {
		t.Errorf("expect no err and size=3,now err:%v,size:%d", err, len(words))
	}
}
