package accounts

import (
    "sync"
    "time"
)

var accountManager *AccountManager = nil
var once sync.Once

func GetAMInstance() *AccountManager {
    once.Do(func() {
        accountManager = &AccountManager{}
    })

    return accountManager
}

type Account struct {
    Address string
}

type AccountManager struct {
    accounts []*Account
}

func (am *AccountManager) CreateAccount(password string) (*Account, error) {
    //will use jeremy's library later...
    newAccount := &Account{time.Now().String()}
    am.accounts = append(am.accounts, newAccount)

    return newAccount, nil
}

func (am *AccountManager) DestroyAccount(pubkey string, password string) (error) {
    //will use jeremy's library later...
    keyword := pubkey

    for i, v := range am.accounts {
        if v.Address == keyword {
            am.accounts = append(am.accounts[:i], am.accounts[i+1:]...)
        }
    }

    return nil
}

func (am *AccountManager) AuthAccount(pubkey string, password string) (bool, error) {
    return true, nil
}

func (am *AccountManager) GetAccountList() ([]*Account) {
    return am.accounts
}

func (am *AccountManager) AccountValid(pubkey string) (bool) {
    for _, v := range am.accounts {
        if v.Address == pubkey {
            return true
        }
    }

    return false
}