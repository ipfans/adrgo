package adrgo_test

import (
	"testing"

	"github.com/ipfans/adrgo"
)

func Test_headerInfo(t *testing.T) {
	type args struct {
		title string
	}
	tests := []struct {
		name     string
		args     args
		wantId   int
		wantName string
	}{
		{
			name: "Test_headerInfo",
			args: args{
				title: "1. hello ",
			},
			wantId:   1,
			wantName: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, gotName := adrgo.HeaderInfo(tt.args.title)
			if gotId != tt.wantId {
				t.Errorf("headerInfo() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotName != tt.wantName {
				t.Errorf("headerInfo() gotName = %v, want %v", gotName, tt.wantName)
			}
		})
	}
}
