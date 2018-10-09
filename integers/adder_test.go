package integers

import (
	"fmt"
	. "github.com/onsi/gomega"
	"testing"
)

func TestAdd(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(Add(2, 2)).To(Equal(4))
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
