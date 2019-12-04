package main

import ("fmt"
        "math"
)

var seen map[int]bool

func init () {
  seen = make(map[int]bool)
}

func SquareHappy(n int) (result int) {
  for n > 0 {
    result += int(math.Pow(math.Mod(float64(n), 10), 2))
    n /= 10
  }

  return
}

func isHappy(n int) bool {
  if value, saw := seen[n]; saw {
    return value
  }

  initN := n
  cyclePreventer := make(map[int] bool)

  for n != 1 {
    n = SquareHappy(n)
    if value, saw := seen[n]; saw {
      seen[initN] = value
      return value
    }

    if n == 0 {
      return false
    }

    if _, saw := cyclePreventer[n]; saw {
      seen[initN] = false
      return false
    }

    cyclePreventer[n] = true

  }

  seen[initN] = true
  return true
}

func main() {
  fmt.Printf("is Happy %v\n", isHappy(2))
}