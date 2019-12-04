package main

import ("fmt"
        "math"
)

func SquareHappy(n int) (result int) {
  for n > 0 {
    result += int(math.Mod(float64(n), 10))
    n /= 10
  }

  return
}

func isHappy(n int) bool {

  if n < 10 {
    return false
  }

  seen := make(map[int]bool)
  seen[n] = true

  for n != 1 {
    n = SquareHappy(n)
    if _,saw := seen[n]; saw {
      return false
    }

    if n == 0 {
      return false
    }

    seen[n] = true
  }

  return true
}

func main() {
  fmt.Printf("is Happy %v\n", isHappy(19))
}