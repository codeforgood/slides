package sum

// START
func Sum(l int, u int) int{
    sum := 0
    for n := l; n < u; n++ {
        sum += n
    }
    return sum
}
//END
