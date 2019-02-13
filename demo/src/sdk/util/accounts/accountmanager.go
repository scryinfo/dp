package accounts

import (
    "../../interface/accountinterface"
    "context"
    "fmt"
    "github.com/pkg/errors"
    "google.golang.org/grpc"
    "sync"
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
    Online bool
}

type AccountManager struct {
    client scryinfo.KeyServiceClient
    accounts []*Account
}

func (am *AccountManager) Initialize(asNodeAddr string) error {
    cn, err := grpc.Dial(asNodeAddr, grpc.WithInsecure())
    if err != nil {
        fmt.Println("failed to connect to node:" + asNodeAddr)
        return err
    }

    am.client = scryinfo.NewKeyServiceClient(cn)
    if am.client == nil {
        return errors.New("failed to create account service client")
    }

    return nil
}

func (am *AccountManager) CreateAccount(password string) (*Account, error) {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("failed to create account, error:", err)
        }
    }()

    if am.client == nil {
        return nil, errors.New("failed to create account, error: null account service")
    }

    addr, err := am.client.GenerateAddress(context.Background(), &scryinfo.AddressParameter{Password:password})
    if err != nil {
        fmt.Println("failed to create account, error:", err)
        return nil, err
    }

    newAccount := &Account{addr.Address, false}
    am.accounts = append(am.accounts, newAccount)

    return newAccount, nil
}

func (am AccountManager) AuthAccount(address string, password string) (bool, error) {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("failed to authenticate account, error:", err)
        }
    }()

    if am.client == nil {
        return false, errors.New("failed to authenticate account, error: null account service")
    }

    addr, err := am.client.VerifyAddress(context.Background(),
        &scryinfo.AddressParameter{Password:password, Address: address})
    if err != nil {
        fmt.Println("failed to authenticate account, error:", err)
        return false, err
    }

    rv := addr.Status == scryinfo.Status_OK

    return rv, nil
}

func (am AccountManager) GetAccounts() ([]*Account) {
    return am.accounts
}

func (am AccountManager) AccountValid(address string) (bool) {
    for _, v := range am.accounts {
        if v.Address == address {
            return true
        }
    }

    return false
}

func (am AccountManager) Encrypt(plainText []byte, address string, password string) ([]byte, error) {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("failed to encrypt, error:", err)
        }
    }()

    if am.client == nil {
        return nil, errors.New("failed to encrypt, error: null account service")
    }

    in := scryinfo.CipherParameter{Message: plainText}
    out, err := am.client.ContentEncrypt(context.Background(), &in)
    if err != nil {
        fmt.Println("failed to encrypt, error:", err)
        return nil, err
    }

    return out.Data, nil
}

func (am AccountManager) ReEncrypt(cipherText []byte, address1 string,
                                    address2 string, password string) ([]byte, error) {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("failed to reencrypt, error:", err)
        }
    }()

    if am.client == nil {
        return nil, errors.New("failed to encrypt, error: null account service")
    }

    in := scryinfo.CipherParameter{Message: cipherText}
    out, err := am.client.ContentDecrypt(context.Background(), &in)
    if err != nil {
        fmt.Println("failed to decrypt, error:", err)
        return nil, err
    }

    in = scryinfo.CipherParameter{Message: out.Data}
    out, err = am.client.ContentEncrypt(context.Background(), &in)
    if err != nil {
        fmt.Println("failed to encrypt, error:", err)
        return nil, err
    }

    return out.Data, nil
}

func (am AccountManager) SignTransaction(message []byte, address string, password string) ([]byte, error) {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("failed to encrypt, error:", err)
        }
    }()

    if am.client == nil {
        return nil, errors.New("failed to encrypt, error: null client")
    }

    in := scryinfo.CipherParameter{
        Message: message,
        Address:address,
        Password:password,
    }

    out, err := am.client.Signature(context.Background(), &in)
    if err != nil {
        fmt.Println("failed to signature, error:", err)
        return nil, err
    }

    return out.Data, nil
}