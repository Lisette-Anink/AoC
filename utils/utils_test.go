package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncludesAll(t *testing.T) {
	type args struct {
		collection []string
		test       []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"case 1", args{[]string{"c", "d", "f", "g", "e", "b"}, four}, false},
		{"case 1", args{[]string{"c", "e", "f", "a", "b", "d"}, four}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IncludesAll(tt.args.collection, tt.args.test); got != tt.want {
				t.Errorf("IncludesAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

var four = []string{"e", "f", "a", "b"}

func TestDifference(t *testing.T) {
	type args struct {
		base      []string
		substract []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"case 1", args{[]string{"c", "d", "f", "g", "e", "b"}, four}, []string{"c", "d", "g"}},
		{"case 2", args{[]string{"c", "e", "f", "a", "b", "d"}, four}, []string{"c", "d"}},
		{"case 3", args{four, []string{"a", "b"}}, []string{"e", "f"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Difference(tt.args.base, tt.args.substract)
			assert.ElementsMatch(t, tt.want, got)
			// assert.Equal(t, []string{"e", "f", "a", "b"}, tt.args.base)

		})
	}
}
