package main

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	str := "a:b:c:d"
	sep := ":"
	rep := Split(str, sep)
	want := []string{"a", "b", "c", "d"}
	if !reflect.DeepEqual(rep, want) {
		t.Errorf("rep:%v, want:%v", rep, want)
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("a:b:c:d", ":")
	}
}
