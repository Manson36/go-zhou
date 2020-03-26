package main

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	ret := Split("babedberdbd", "b")
	want := []string{"", "a", "ed", "erd", "d"}
	if !reflect.DeepEqual(ret, want) {
		//测试用例失败了
		t.Errorf("want:%v, but got:%v", want, ret)
	}
}

//Benchmark 基准测试
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("babedberdbd", "b")
	}
}
