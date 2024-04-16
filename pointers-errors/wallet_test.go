package pointers_errors

import (
	"fmt"
	"testing"
)

func BenchmarkWallet_Deposit(b *testing.B) {
	wallet := Wallet{Bitcoin(10)}
	for i := 0; i < b.N; i++ {
		wallet.Deposit(10)
	}
}

func ExampleWallet_Deposit() {
	wallet := Wallet{Bitcoin(10)}
	wallet.Deposit(10)
	fmt.Println(wallet.Balance())
	// Output: 20 BTC
}

func BenchmarkWallet_Withdraw(b *testing.B) {
	wallet := Wallet{Bitcoin(10)}
	for i := 0; i < b.N; i++ {
		err := wallet.Withdraw(10)
		if err != nil {
			return
		}
	}
}

func ExampleWallet_Withdraw() {
	wallet := Wallet{Bitcoin(10)}
	err := wallet.Withdraw(10)
	if err != nil {
		return
	}
	fmt.Println(wallet.Balance())
	// Output: 0 BTC
}

func BenchmarkWallet_Balance(b *testing.B) {
	wallet := Wallet{Bitcoin(10)}
	for i := 0; i < b.N; i++ {
		wallet.Balance()
	}
}

func ExampleWallet_Balance() {
	wallet := Wallet{Bitcoin(10)}
	fmt.Println(wallet.Balance())
	// Output: 10 BTC
}

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}
	t.Run("deposits Bitcoin into the wallet", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraws Bitcoin from the wallet", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(0))
	})
	t.Run("does not withdraw more funds than available", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "cannot withdraw, insufficient funds")
		assertBalance(t, wallet, startingBalance)
	})
}
