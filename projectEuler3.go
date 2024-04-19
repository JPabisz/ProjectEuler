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
    var x float64 = math.Sqrt(b)
    var y int = int(x) - 1
    for i:=y; i > 0; i-- {
        if a%i == 0 && largestPrime(i) == 1 {
            return i
        }
    }
    return 1
}
