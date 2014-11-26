package main

import (
  "fmt"
)

func sum(a []int) int{
    sum := 0
    for _, v := range a {
        sum += v
    }
    return sum
}

func main() {
    a := []int{10,45,99,1000000}
    fmt.Println(sum(a))
}
