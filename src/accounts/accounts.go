package accounts

// Account - exchange account
type Account struct {
	Balance int `json:"balance"`

	Username string `json:"username"`
	Email    string `json:"email"`
	PassHash string `json:"passwordhash"`

	WalletAddresses []string `json:"walletaddresses"`
	WalletBalances  []int
}

// NewAccount - create, return new account
func NewAccount(username string, email string, passHash string, walletaddrs []string) Account {
	rAccount := Account{Username: username, Email: email, PassHash: passHash, WalletAddresses: walletaddrs}
	return rAccount
}
