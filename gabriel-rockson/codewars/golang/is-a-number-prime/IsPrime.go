package main

import (
	"fmt"
)

func IsPrime(n int) bool {
  if n <= 0 {
    return false
  }

  for a := 2; a * a <= n; a++ {
    if n % a == 0 {
      return false
    }
  }
  
  return true
}

func main() {
  fmt.Println(IsPrime(9))
  fmt.Println(IsPrime(1))
  fmt.Println(IsPrime(75))
  fmt.Println(IsPrime(5))
  fmt.Println(IsPrime(2))
  fmt.Println(IsPrime(3))
  fmt.Println(IsPrime(-20))
}
