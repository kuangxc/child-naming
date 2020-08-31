package types

import (
	"fmt"
	"strconv"
)

//NameInfo name and related info
type NameInfo struct {
	Name         string   `json:"name,omitempty"`
	RelatedNames []string `json:"related_names,omitempty"`
	From         string   `json:"from,omitempty"`
	Meaning      string   `json:"meaning,omitempty"`
}

//Poetry 诗词
type Poetry struct {
	Category   string   `json:"category,omitempty"`   //类别，唐诗宋词等
	Author     string   `json:"author,omitempty"`     //作者
	Title      string   `json:"title,omitempty"`      //名称
	Paragraphs []string `json:"paragraphs,omitempty"` //内容
}

//Pinyin chinese pinyin
type Pinyin struct {
	Word   string `json:"word,omitempty"`   //黄
	Pinyin string `json:"pinyin,omitempty"` //huang
	Tone   int    `json:"tone,omitempty"`   //1 2 3 4
}

func ParsePinyin(s string) (*Pinyin, error) {
	rs := []rune(s)
	if len(rs) < 3 {
		return nil, fmt.Errorf("pinyin should size>=3,now:%d", len(rs))
	}
	tone, _ := strconv.Atoi(string(rs[len(rs)-1]))
	p := &Pinyin{
		Word:   string(rs[0]),
		Pinyin: string(rs[1 : len(rs)-1]),
		Tone:   tone,
	}
	if p.Tone <= 0 || p.Tone > 5 {
		return nil, fmt.Errorf("tone %d is illegal", p.Tone)
	}
	return p, nil
}
