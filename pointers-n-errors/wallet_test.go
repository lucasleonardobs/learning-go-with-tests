package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		validateBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw in sufficient balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		validateBalance(t, wallet, Bitcoin(10))
		validateInexistentError(t, err)
	})

	t.Run("Withdraw in insufficient balance", func(t *testing.T) {
		startBalance := Bitcoin(10)
		wallet := Wallet{balance: startBalance}
		err := wallet.Withdraw(Bitcoin(20))

		validateBalance(t, wallet, startBalance)
		validateError(t, err, ErrorInsufficientBalance)
	})
}

func validateBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()

	result := wallet.Balance()

	if result != expected {
		t.Errorf("result %s, expected %s", result, expected)
	}
}

func validateInexistentError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("unexpected error received.")
	}
}

func validateError(t *testing.T, err error, expected error) {
	t.Helper()

	if err == nil {
		t.Fatalf("An error is expected, but was received nil.")
	}

	if err != expected {
		t.Errorf("result %s, expected %s", err, expected)
	}
}
