package sum

import "testing"

func BenchmarkSum10(b *testing.B) {
    // run the Sum function b.N times
    for n := 0; n < b.N; n++ {
            Sum(1, 10)
    }
}
