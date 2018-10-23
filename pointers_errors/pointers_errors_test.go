package pointers_errors

import (
	"fmt"
	. "github.com/onsi/gomega"

	"testing"
)

// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors#refactor-1

func TestWallet(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		fmt.Printf("address of balance in test is %v \n", &wallet.balance)

		got := wallet.Balance()
		want := Bitcoin(10)

		g.Expect(got).To(Equal(want), "got %s want %s", got, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		fmt.Printf("address of balance in test is %v \n", &wallet.balance)

		got := wallet.Balance()
		want := Bitcoin(10)

		g.Expect(got).To(Equal(want), "got %s want %s", got, want)
	})
}
