package custom

import(
  "fmt"
)

func main() {
  slice_data := []int{1,2,3,4,5,6,6}
  fmt.Println("series sum: ", mySum(slice_data...))
  fmt.Println("2 + 3 =", mySum(2, 3))
  fmt.Println("4 + 7 =", mySum(4, 7))
  fmt.Println("5 + 9 =", mySum(5, 9))
}

func mySum(xi ...int) int {
  sum := 0
  for _, v := range xi {
    sum += v
  }
  return sum
}