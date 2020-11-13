package main

import (
	"reflect"
	"testing"
)

type numReverseTest struct {
	number int
	want   int
}

func TestNumReverse(t *testing.T) {
	tgroup := []numReverseTest{
		{123, 321},
		{-123, -321},
		{700, 7},
		{1534236469, 0},
	}
	for _, i := range tgroup {
		got := NumReverse(i.number)
		if !reflect.DeepEqual(got, i.want) {
			t.Fatalf("got: %v, want: %v", got, i.want)
		}
	}
}
