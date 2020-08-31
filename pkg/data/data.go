package data

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/kuangxc/child-naming/pkg/types"
	"github.com/kuangxc/child-naming/pkg/util"
)

var (
	poetryRootPath = "./data/poetry/"
	commonWordPath = "./data/common/common.txt"
	pinyinPath     = "./data/pinyin/pinyin.txt"
)

func ParsePoetryData(excludeDirs []string) (map[string][]*types.Poetry, error) {
	poetryMap := make(map[string][]*types.Poetry)
	//1.get dir list and exclude exclude dirs
	rd, err := ioutil.ReadDir(poetryRootPath)
	if nil != err {
		return nil, fmt.Errorf("read dir err:%v", err)
	}
	for _, fi := range rd {
		if fi.IsDir() {
			if util.IncludeStr(excludeDirs, fi.Name()) {
				fmt.Printf("exclude dir %s \n", fi.Name())
				continue
			}
			fmt.Println("start parse files of dir ", fi.Name())
			//2.range all file and unmarshal
			ps, err := parsePoetryDataFromDir(poetryRootPath + "/" + fi.Name())
			if nil != err {
				fmt.Printf("%s parse file err:%v", fi.Name(), err)
			}
			for _, p := range ps {
				p.Category = fi.Name()
			}
			poetryMap[fi.Name()] = ps
		} else {
			fmt.Println(fi.Name())
		}
	}

	return poetryMap, nil
}

func parsePoetryDataFromDir(dirPath string) ([]*types.Poetry, error) {
	var ps []*types.Poetry
	rd, err := ioutil.ReadDir(dirPath)
	if nil != err {
		return nil, fmt.Errorf("read dir err:%v", err)
	}
	for _, fi := range rd {
		if fi.IsDir() || !strings.Contains(fi.Name(), ".json") {
			continue
		}
		fmt.Printf("start parse file %s/%s ", dirPath, fi.Name())
		var fps []*types.Poetry
		bytes, err := ioutil.ReadFile(dirPath + "/" + fi.Name())
		if err != nil {
			fmt.Printf("read file err:%v", err)
			continue
		}
		err = json.Unmarshal(bytes, &fps)
		if nil != err {
			fmt.Printf("%s unmarshal err:%v", fi.Name(), err)
			continue
		}
		ps = append(ps, fps...)
	}
	fmt.Printf("%s read poetry size:%d", dirPath, len(ps))
	return ps, nil
}

func ParseCommonWord() ([]string, error) {
	var words []string
	f, err := os.Open(commonWordPath)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReader(f)
	for {
		word, err := buf.ReadString('\n')
		word = strings.TrimSpace(word)
		words = append(words, word)
		if err == io.EOF {
			break
		}
	}
	return words, nil
}

func ParsePinyin() ([]*types.Pinyin, error) {
	var pys []*types.Pinyin
	f, err := os.Open(pinyinPath)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReader(f)
	for {
		word, err := buf.ReadString('\n')
		if nil != err {
			fmt.Println("read pinyin file err:", err)
			break
		}
		word = strings.TrimSpace(word)
		var py *types.Pinyin
		py, err = types.ParsePinyin(word)
		if nil != err {
			fmt.Printf("%s parse pinyin err:%v \r\n", word, err)
			continue
		}
		pys = append(pys, py)
		if err == io.EOF {
			break
		}
	}
	return pys, nil
}
