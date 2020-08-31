package naming

import (
	"fmt"

	"github.com/kuangxc/child-naming/pkg/alg/naming"
	"github.com/kuangxc/child-naming/pkg/conf"
	"github.com/kuangxc/child-naming/pkg/data"
	"github.com/kuangxc/child-naming/pkg/output"
	"github.com/kuangxc/child-naming/pkg/types"
)

func Run() {
	//1.加载常用字符、拼音库
	fmt.Println("start get common words ...")
	comWords, err := data.ParseCommonWord()
	if nil != err {
		fmt.Println("parse common word data err:", err)
		return
	}
	fmt.Println("start parse pinyin ...")
	var ps []*types.Pinyin
	ps, err = data.ParsePinyin()
	if nil != err {
		fmt.Println("parse pinyin data err:", err)
		return
	}
	//2.生成名字列表
	surnamePy, err := getWordPinyin(conf.GetSurname(), ps)
	if nil != err {
		fmt.Printf("%s pinyin not find,err:%v", conf.GetSurname(), err)
		return
	}
	fmt.Println("start generate name ...")
	rawNameChn := naming.GenerateNames(surnamePy.Word)
	//3.根据配置，剔除不符合要求的名字
	filterExcludeNameChn := filterExclude(rawNameChn, conf.GetExcludeWords(), conf.GetExcludeNames())
	//4. 剔除包含不常见字的名字
	filteUncommonNameChn := filterInclude(filterExcludeNameChn, comWords)
	//5.剔除与姓氏同声母的名字，例如吴王(wu wang)
	//6.剔除连续韵母相同，例如于玉秋(yuyuqiu)
	//7.根据拼音筛选平仄,避免同声，例如黄强
	filterDisHanmonyNameChn := filtedPinyinDisharmony(filteUncommonNameChn, ps)
	//8.输出姓名
	output.SaveExcel(filterDisHanmonyNameChn, conf.GetOutputNum())
}
