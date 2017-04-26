package main

import (
  "strings"
  "fmt"
)

func ComputePrefix(pattern []byte) []int {
  pi := make([]int, len(pattern))

  k := -1
  pi[0] = k
  for i := 1; i < len(pattern); i++ {
    for  k > -1 && pattern[k+1] != pattern[i] {
      k = pi[k]
    }
    if pattern[i] == pattern[k+1] {
      k++
    }
    pi[i] = k
  }
  return pi
}

func Kmp(target []byte, pattern []byte) (result []int) {
  result = make([]int, 0, len(target))
  pi := ComputePrefix(pattern)
  k := -1
  for i := 0; i < len(target); i++ {
    for k > -1 && pattern[k+1] != target[i] {
      k = pi[k]
    }
    if target[i] == pattern[k+1] {
      k++
    }
    if k == len(pattern) - 1 {
      result = append(result, i-k)
      k = -1
    }
  }
  return
}

func ToArray(str string) []byte {
    a := make([]byte, len(str))
   copy(a[:], str)
   return a
}

func StringSearch(target string, pattern string) []int {
  t := ToArray(strings.ToLower(target))
  p := ToArray(strings.ToLower(pattern))
  result := Kmp(t, p)
  for i := range result {
    result[i] = result[i]+1
  }
  return result
}

//func main() {
//  target := "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"
//  pattern := "Peterz"
//  result := StringSearch(target, pattern)
//  fmt.Printf("%v\n", result)
//}
