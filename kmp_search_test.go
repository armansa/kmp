package main

import (
  "strconv"
  "testing"
  "strings"
)

func TestIntersectionTick(t *testing.T) {
  changed := intersection.Tick()
  if ! changed {
    t.Error("Expected changed = true, got false")
  }
}

