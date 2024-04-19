package main
import (
    "fmt"
    "math"
)

func main() {
  fmt.Println(largestPrime(600851475143))
}

func largestPrime(a int)int {
    var b float64 = float64(a)
