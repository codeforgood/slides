package main

import (
  "fmt"
)

func sum2(a []int, c chan int) {
    sum := 0
    for _, v := range a {
        sum += v
    }
    c <- sum // send sum to c
}

func main() {
    a := []int{10,45,99,1000000}
    c := make(chan int)

    go sum2(a[:len(a)/2], c) // first half
    go sum2(a[len(a)/2:], c) // second half

    x, y := <-c, <-c // receive from c
    fmt.Println(x, y, x+y)
}
