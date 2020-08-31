package conf

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

const (
	confFile = "naming.yaml"
)

type RequireConf struct {
	Surname string `json:"surname,omitempty"`
}

type ExcludeConf struct {
	Words string `json:"words,omitempty"`
	Names string `json:"names,omitempty"`
	Dirs  string `json:"dirs,omitempty"`
}

type OutputConf struct {
	Format string `json:"format,omitempty"`
	Number int    `json:"number,omitempty"`
}

//config config info for naming
type config struct {
	Require RequireConf `json:"require,omitempty"`
	Exclude ExcludeConf `json:"exclude,omitempty"`
	Output  OutputConf  `json:"output,omitempty"`
}

var defaultConf *config

func init() {
	defaultConf = &config{}
	f, err := os.Open(confFile)
	if err != nil {
		fmt.Printf("read conf file %s fail!", confFile)
		os.Exit(-1)
	}
	err = yaml.NewDecoder(f).Decode(defaultConf)
	if nil != err {
		fmt.Printf("decode conf file %s err:%v \r\n", confFile, err)
		os.Exit(-1)
	}
	bytes, err := json.Marshal(defaultConf)
	fmt.Printf("get config:%s,err:%v\r\n", string(bytes), err)
}

func GetExcludeDirs() []string {
	return strings.Split(defaultConf.Exclude.Dirs, ";")
}

func GetSurname() string {
	return defaultConf.Require.Surname
}

func GetExcludeNames() []string {
	return strings.Split(defaultConf.Exclude.Names, ";")
}
func GetExcludeWords() []string {
	return strings.Split(defaultConf.Exclude.Words, ";")
}

func GetOutputNum() int {
	return defaultConf.Output.Number
}
