package output

import (
	"testing"

	"github.com/kuangxc/child-naming/pkg/types"
)

func TestSaveExcel(t *testing.T) {
	type args struct {
		names []*types.NameInfo
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-ok",
			args: args{
				names: []*types.NameInfo{
					&types.NameInfo{
						Name: "邝帅哥",
						RelatedNames: []string{
							"潘安", "楚留香",
						},
						From:    "诗经",
						Meaning: "就是很帅，没其他的。",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveExcel(tt.args.names)
		})
	}
}

func TestSaveText(t *testing.T) {
	type args struct {
		names []*types.NameInfo
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test-ok",
			args: args{
				names: []*types.NameInfo{
					&types.NameInfo{
						Name: "邝帅哥",
						RelatedNames: []string{
							"潘安", "楚留香",
						},
						From:    "诗经",
						Meaning: "就是很帅，没其他的。",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveText(tt.args.names)
		})
	}
}
