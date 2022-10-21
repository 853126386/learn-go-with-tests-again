package pointer

import "testing"

func TestWallet(t *testing.T)  {

	assertBalance := func(t *testing.T,wallet Wallet,want Btc) {
		got :=wallet.Blance()
		if want != got{
			t.Errorf("got %s but want %s",got,want)
		}
	}

	assertError := func(t *testing.T,got error,want error) {
		if got == nil{
			t.Error("wanted an error but didnt get one")
		}

		if got != want{
			t.Errorf("got '%s', want '%s'", got, want)
		}

	}

	t.Run("Wallet Deposit ", func(t *testing.T) {
		wallet :=Wallet{}
		wallet.Deposit(Btc(10))
		want :=Btc(10)
		assertBalance(t,wallet,want)
	})


	t.Run("Wallet Withdraw ", func(t *testing.T) {
		wallet :=Wallet{blance: Btc(10)}
		err :=wallet.Withdraw(Btc(101))
		want :=Btc(1)
		assertBalance(t,wallet,want)
		assertError(t, err, InsufficientFundsError)
	})

}