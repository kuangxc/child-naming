package output

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/kuangxc/child-naming/pkg/types"
)

func TestSaveExcel(t *testing.T) {
	// SaveExcel calls os.Exit(0) by design, which would terminate `go test`.
	// We intentionally skip it here.
	t.Skip("SaveExcel exits the process (os.Exit), not suitable for unit tests")
}

func TestSaveText(t *testing.T) {
	tmp := t.TempDir()
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd err:%v", err)
	}
	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("chdir err:%v", err)
	}
	defer func() { _ = os.Chdir(cwd) }()

	names := []*types.NameInfo{
		{
			Name: "邝帅哥",
			RelatedNames: []string{
				"潘安", "楚留香",
			},
			From:    "诗经",
			Meaning: "就是很帅，没其他的。",
		},
	}
	want := time.Now().Format("2006-01-02") + ".txt"

	SaveText(names, 10)

	if _, err := os.Stat(filepath.Join(tmp, want)); err != nil {
		t.Fatalf("expected output file %s, stat err:%v", want, err)
	}
}
