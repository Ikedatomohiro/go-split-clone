package util_test

import (
	"testing"

	ut "split-clone/util"
)

func TestGetFilename(t *testing.T) {
	var tests = []struct {
		name    string
		args    string
		want    string
		wantErr bool
	}{
		{
			name:    "success",
			args:    "xaa",
			want:    "xab",
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
