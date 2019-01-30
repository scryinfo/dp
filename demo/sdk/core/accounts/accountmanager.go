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
}

type Connector struct {
    conn *grpc.ClientConn
}

type AccountManager struct {
    client scryinfo.KeyServiceClient
    accounts []*Account
}

func NewConnector(asNodeAddr string) (*Connector, error) {
    cn, err := grpc.Dial(asNodeAddr)
    if err != nil {
        fmt.Println("failed to connect to node:" + asNodeAddr)
        return nil, err
    }

    return &Connector{
        conn: cn,
    }, nil
}

func (am *AccountManager) Initialize(asNodeAddr string) error {
    c, err := NewConnector(asNodeAddr)
    if err != nil {
        return err
    }

    am.client = scryinfo.NewKeyServiceClient(c.conn)
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

    newAccount := &Account{addr.Address}
    am.accounts = append(am.accounts, newAccount)

    return newAccount, nil
}

func (am *AccountManager) AuthAccount(address string, password string) (bool, error) {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("failed to authenticate account, error:", err)
        }
    }()

    if am.client == nil {
        return false, errors.New("failed to authenticate account, error: null account service")
    }

    addr, err := am.client.VerifyAddress(context.Background(), &scryinfo.AddressParameter{Password:password, Address: address})
    if err != nil {
        fmt.Println("failed to authenticate account, error:", err)
        return false, err
    }

    rv := addr.Status == scryinfo.Status_OK

    return rv, nil
}

func (am *AccountManager) GetAccountList() ([]*Account) {
    return am.accounts
}

func (am *AccountManager) AccountValid(address string) (bool) {
    for _, v := range am.accounts {
        if v.Address == address {
            return true
        }
    }

    return false
}