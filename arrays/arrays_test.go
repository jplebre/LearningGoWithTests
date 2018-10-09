package arrays

import (
	. "github.com/onsi/gomega"
	"reflect"
	"testing"
)

func TestSumWithArrays(t *testing.T) {
	g := NewGomegaWithT(t)
	numbers := [5]int{1, 2, 3, 4, 5}

	t.Run("Run the code using a full 'for' loop", func(t *testing.T) {
		g.Expect(SumWithArrays(numbers)).To(Equal(15))
	})

	t.Run("Run the code using a range in 'for' loop", func(t *testing.T) {
		g.Expect(SumWithArraysUsingRange(numbers)).To(Equal(15))
	})
}

func TestSumWithSlices(t *testing.T) {
	g := NewGomegaWithT(t)

	numbers := []int{1, 2, 3}

	t.Run("Run the code using a full 'for' loop", func(t *testing.T) {
		g.Expect(SumWithSlices(numbers)).To(Equal(6))
	})
}

func TestSumAll(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Sums one array correctly", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		// This won't work, you can't use equality operators with slices
		// if got != want {
		// 	t.Errorf("got %v want %v", got, want)
		// }

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Sums one array correctly", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		g.Expect(got).To(Equal(want))
	})

	t.Run("Sums more than one array correctly", func(t *testing.T) {
		g.Expect(SumAll([]int{1, 1, 1, 1})).To(Equal([]int{4}))
	})
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{0, 9})
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
