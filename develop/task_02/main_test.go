package main

import (
	"reflect"
	"testing"
)

func TestUnpacker(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Empty string", args{""}, ""},
		{"Single character", args{"abcd"}, "abcd"},
		{"Simple repetition", args{"a6b2c"}, "aaaaaabbc"},
		{"Invalid input: digit first", args{"5ra"}, ""},
		{"Complex repetition0", args{"a4b2c3d"}, "aaaabbcccd"},
		{"Complex repetition1", args{"a4bc2d5e"}, "aaaabccddddde"},
		{"Complex repetition2", args{"f3cn20"}, "fffcnnnnnnnnnnnnnnnnnnnn"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Unpacker(tt.args.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unpacker() = %v, want %v", got, tt.want)
			}
		})
	}
}
