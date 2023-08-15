package util_test

import (
	"testing"

	in "split-clone/input"
	ut "split-clone/util"
)

type GetParamArgs struct {
	args []string
	ap   in.ArgPosition
}

func TestGetFilename(t *testing.T) {
	var tests = []struct {
		name    string
		args    string
		want    string
		wantErr bool
	}{
		{
			name:    "success",
			args:    "aa",
			want:    "ab",
			wantErr: false,
		},
		{
			name:    "success: zの処理",
			args:    "az",
			want:    "ba",
			wantErr: false,
		},
		{
			name:    "success: zの処理",
			args:    "zzz",
			want:    "aaaa",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ut.GetFilename(tt.args)
			if got != tt.want {
				t.Errorf("GetFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetParam(t *testing.T) {
	var tests = []struct {
		name    string
		args    GetParamArgs
		want    in.Input
		wantErr bool
	}{
		{
			name: "success: オプションなし、プレフィックスなし（-l 1000のオプションをデフォルトとする）",
			args: GetParamArgs{
				args: []string{"split", "test.txt"},
				ap:   in.ArgPosition{Option: 0, Prefix: 0, FileName: 1},
			},
			want:    in.Input{Option: "l", OptionValue: 1000, FileName: "test.txt", Prefix: "x"},
			wantErr: false,
		},
		{
			name: "success: オプションあり、プレフィックスなし",
			args: GetParamArgs{
				args: []string{"split", "-b", "3000", "test.txt"},
				ap:   in.ArgPosition{Option: 1, Prefix: 0, FileName: 3},
			},
			want:    in.Input{Option: "b", OptionValue: 3000, FileName: "test.txt", Prefix: "x"},
			wantErr: false,
		},
		{
			name: "success: オプションあり、プレフィックスなし",
			args: GetParamArgs{
				args: []string{"split", "-b", "300m", "test.txt"},
				ap:   in.ArgPosition{Option: 1, Prefix: 0, FileName: 3},
			},
			want:    in.Input{Option: "b", OptionValue: 300000000, FileName: "test.txt", Prefix: "x"},
			wantErr: false,
		},
		{
			name: "success: オプションあり、プレフィックスなし",
			args: GetParamArgs{
				args: []string{"split", "-b", "300k", "test.txt"},
				ap:   in.ArgPosition{Option: 1, Prefix: 0, FileName: 3},
			},
			want:    in.Input{Option: "b", OptionValue: 300000, FileName: "test.txt", Prefix: "x"},
			wantErr: false,
		},
		{
			name: "success: オプションあり、プレフィックスなし",
			args: GetParamArgs{
				args: []string{"split", "-b", "3G", "test.txt"},
				ap:   in.ArgPosition{Option: 1, Prefix: 0, FileName: 3},
			},
			want:    in.Input{Option: "b", OptionValue: 3000000000, FileName: "test.txt", Prefix: "x"},
			wantErr: false,
		},
		{
			name: "success: オプションあり、プレフィックスなし",
			args: GetParamArgs{
				args: []string{"split", "-n", "3", "test.txt"},
				ap:   in.ArgPosition{Option: 1, Prefix: 0, FileName: 3},
			},
			want:    in.Input{Option: "n", OptionValue: 3, FileName: "test.txt", Prefix: "x"},
			wantErr: false,
		},
		{
			name: "success: オプションなし、プレフィックスあり",
			args: GetParamArgs{
				args: []string{"split", "test.txt", "sample_"},
				ap:   in.ArgPosition{Option: 0, Prefix: 2, FileName: 1},
			},
			want:    in.Input{Option: "l", OptionValue: 1000, FileName: "test.txt", Prefix: "sample_"},
			wantErr: false,
		},
		{
			name: "success: オプションあり、プレフィックスあり",
			args: GetParamArgs{
				args: []string{"split", "-n", "3", "test.txt", "sample_"},
				ap:   in.ArgPosition{Option: 1, Prefix: 4, FileName: 3},
			},
			want:    in.Input{Option: "n", OptionValue: 3, FileName: "test.txt", Prefix: "sample_"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ut.GetParam(tt.args.args, tt.args.ap)
			if got != tt.want {
				t.Errorf("GetParam() got = %v, want %v", got, tt.want)
			}
		})
	}
}
