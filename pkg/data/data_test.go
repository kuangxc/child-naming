package data

import (
	"testing"
)

func TestParsePoetryData(t *testing.T) {
	poetryRootPath = "./poetry/"
	pm, err := ParsePoetryData([]string{"song"})
	if nil != err {
		t.Errorf("parse poetry data expect no err,now:%v", err)
	}
	if len(pm) != 2 {
		t.Errorf("expect 2 types poetry,now:%d", len(pm))
	}
	for name, ps := range pm {
		if len(ps) <= 0 {
			t.Errorf("expect %s poetry size >0 now:%d", name, len(ps))
		}
	}
}

func TestParseCommonWordFromDir(t *testing.T) {
	commonWordPath = "./common/common.txt"
	words, err := ParseCommonWord()
	if nil != err || len(words) != 3500 {
		t.Errorf("expect words 3500 and no err,now: words %d err:%v", len(words), err)
	}
}

func TestParsePinyin(t *testing.T) {
	pinyinPath = "./pinyin/pinyin.txt"
	words, err := ParsePinyin()
	if nil != err || len(words) != 20519 {
		t.Errorf("expect pinyin 20519 and no err,now: pinyin %d err:%v", len(words), err)
	}
}
