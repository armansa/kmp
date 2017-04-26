package main

import (
  "testing"
)

func CompareArrays(a []int, b []int) bool {
  if len(a) != len(b) {
    return false
  }
  for i := range a {
    if a[i] != b[i] {
      return false
    }
  }
  return true
}

func TestComputePrefix(t *testing.T) {
  prefix_indexes := ComputePrefix([]byte("ABCABCAB"))
  if ! CompareArrays(prefix_indexes, []int{-1, -1, -1, 0, 1, 2, 3, 4}) {
    t.Error("error")
  }
}

func TestKmp(t *testing.T) {
  a := []byte{1, 2, 3, 4, 1, 2, 3, 5, 1, 2, 3, 6, 1, 2, 3}
  b := []byte{1, 2, 3, 6, 1, 2, 3}
  result := Kmp(a, b)
  if ! CompareArrays(result, []int{8}) {
    t.Error("error", result)
  }
}

func TestStringSearch(t *testing.T) {
  text := "ABCDABCEABCFABC"
  pattern := "ABCFA"
  result := StringSearch(text, pattern)
  if ! CompareArrays(result, []int{9}) {
    t.Error("error", result)
  }

  text = "ABCDABCEABCFABC"
  pattern = "ABC"
  result = StringSearch(text, pattern)
  if ! CompareArrays(result, []int{1, 5, 9, 13}) {
    t.Error("error", result)
  }
}

