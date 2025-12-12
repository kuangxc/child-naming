package conf

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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
	f, confPath, err := openConfFile()
	if err != nil {
		fmt.Printf("read conf file %s fail! err:%v", confFile, err)
		os.Exit(-1)
	}
	defer func() { _ = f.Close() }()

	err = yaml.NewDecoder(f).Decode(defaultConf)
	if nil != err {
		fmt.Printf("decode conf file %s err:%v \r\n", confPath, err)
		os.Exit(-1)
	}
	bytes, err := json.Marshal(defaultConf)
	fmt.Printf("get config:%s,err:%v\r\n", string(bytes), err)
}

func openConfFile() (*os.File, string, error) {
	if p := strings.TrimSpace(os.Getenv("CHILD_NAMING_CONF")); p != "" {
		f, err := os.Open(p)
		return f, p, err
	}

	// 1) Try current working directory.
	if f, err := os.Open(confFile); err == nil {
		return f, confFile, nil
	}

	// 2) Try repository root inferred from this source file location.
	// This makes `go test ./...` work regardless of working directory.
	_, thisFile, _, ok := runtime.Caller(0)
	if ok {
		// /path/to/repo/pkg/conf/conf.go -> /path/to/repo
		repoRoot := filepath.Clean(filepath.Join(filepath.Dir(thisFile), "..", ".."))
		p := filepath.Join(repoRoot, confFile)
		if f, err := os.Open(p); err == nil {
			return f, p, nil
		}
	}

	return nil, "", fmt.Errorf("config file not found: %s (cwd or repo root); you can set env CHILD_NAMING_CONF", confFile)
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
