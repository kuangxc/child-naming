package types

import (
	"testing"
)

func TestParsePinyin(t *testing.T) {
	legalPinyin := "邝kuang4"
	p, err := ParsePinyin(legalPinyin)
	if nil != err {
		t.Errorf("expect no err,now:%v", err)
	} else if p.Tone != 4 {
		t.Errorf("expect tone=4,now:%d", p.Tone)
	} else if p.Pinyin != "kuang" {
		t.Errorf("expect pinyin=kuang,now:%s", p.Pinyin)
	}
	illegalPinyin := "邝kuang"
	_, err = ParsePinyin(illegalPinyin)
	if nil == err {
		t.Errorf("expect err,now:%v", err)
	}
}
