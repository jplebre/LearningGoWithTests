package iteration

import (
	. "github.com/onsi/gomega"

	"testing"
)

func TestRepeat(t *testing.T) {
	g := NewGomegaWithT(t)

	repeated := Repeat("a")
	expected := "aaaaa"

	g.Expect(repeated).To(Equal(expected))
}

func TestRepeatWithNoInitialization(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(RepeatWithNoInitialization(5)).To(Equal(5))
}

// Built in benchmark tool in Go
// To run do `go test -bench=.`
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
