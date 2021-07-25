package custom

import(
  "testing"
)

func TestMySum(t *testing.T) {
  x := mySum(2, 3)
  if x != 5 {
    t.Error("Expected", 5, "Got", x)
  }
}

// cd /Users/username/workspace/backend/go/src/go_workspace/src/9-testing/custom
// go test
// go test v (v stands for verbose)