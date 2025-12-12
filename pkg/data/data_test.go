package data

import (
	"path/filepath"
	"runtime"
	"testing"
)

func repoRootForTest(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("failed to locate test file path")
	}
	// /path/to/repo/pkg/data/data_test.go -> /path/to/repo
	return filepath.Clean(filepath.Join(filepath.Dir(thisFile), "..", ".."))
}

func TestParsePoetryData(t *testing.T) {
	root := repoRootForTest(t)
	poetryRootPath = filepath.Join(root, "data", "poetry")
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
	root := repoRootForTest(t)
	commonWordPath = filepath.Join(root, "data", "common", "common.txt")
	words, err := ParseCommonWord()
	if nil != err || len(words) != 3500 {
		t.Errorf("expect words 3500 and no err,now: words %d err:%v", len(words), err)
	}
}

func TestParsePinyin(t *testing.T) {
	root := repoRootForTest(t)
	pinyinPath = filepath.Join(root, "data", "pinyin", "pinyin.txt")
	words, err := ParsePinyin()
	if nil != err || len(words) != 20519 {
		t.Errorf("expect pinyin 20519 and no err,now: pinyin %d err:%v", len(words), err)
	}
}
