package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got %d want %d given %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}

	got := SumAll(slice1, slice2)
	want := []int{6, 15}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
