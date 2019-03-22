package accounts

import (
	"context"
	"github.com/pkg/errors"
	"github.com/scryinfo/iscap/demo/src/sdk/interface/accountinterface"
	rlog "github.com/sirupsen/logrus"
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

func ResetAMInstance() {
	accountManager.accounts = nil
}

type Account struct {
	Address string
}

type AccountManager struct {
	client   scryinfo.KeyServiceClient
	accounts []*Account
}

func (am *AccountManager) Initialize(asNodeAddr string) error {
	cn, err := grpc.Dial(asNodeAddr, grpc.WithInsecure())
	if err != nil {
		rlog.Error("failed to connect to node:" + asNodeAddr)
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
			rlog.Error("failed to create account, error:", err)
		}
	}()

	if am.client == nil {
		return nil, errors.New("failed to create account, error: null account service")
	}

	addr, _ := am.client.GenerateAddress(context.Background(), &scryinfo.AddressParameter{Password: password})
	if addr.Status != scryinfo.Status_OK {
		rlog.Error("failed to create account, error:", addr.Msg)
		return nil, errors.New("failed to create account, error:" + addr.Msg)
	}

	newAccount := &Account{addr.Address}
	am.accounts = append(am.accounts, newAccount)

	return newAccount, nil
}

func (am AccountManager) AuthAccount(address string, password string) (bool, error) {
	defer func() {
		if err := recover(); err != nil {
			rlog.Error("failed to authenticate account, error:", err)
		}
	}()

	if am.client == nil {
		return false, errors.New("failed to authenticate account, error: null account service")
	}

	addr, _ := am.client.VerifyAddress(context.Background(),
		&scryinfo.AddressParameter{Password: password, Address: address})
	rv := addr.Status == scryinfo.Status_OK

	return rv, nil
}

func (am AccountManager) GetAccounts() []*Account {
	return am.accounts
}

func (am AccountManager) AccountValid(address string) bool {
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
			rlog.Error("failed to encrypt, error:", err)
		}
	}()

	if am.client == nil {
		return nil, errors.New("failed to encrypt, error: null account service")
	}

	in := scryinfo.CipherParameter{Message: plainText, Address: address}
	out, _ := am.client.ContentEncrypt(context.Background(), &in)
	if out.Status != scryinfo.Status_OK {
		rlog.Error("failed to encrypt, error:", out.Msg)
		return nil, errors.New("failed to encrypt, error:" + out.Msg)
	}

	return out.Data, nil
}

func (am AccountManager) ReEncrypt(cipherText []byte, address1 string, address2 string, password string) ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			rlog.Error("failed to re-encrypt, error:", err)
		}
	}()

	if am.client == nil {
		return nil, errors.New("failed to encrypt, error: null account service")
	}

	in := scryinfo.CipherParameter{Message: cipherText, Address: address1, Password: password}
	out, _ := am.client.ContentDecrypt(context.Background(), &in)
	if out.Status != scryinfo.Status_OK {
		rlog.Error("failed to decrypt, error:", out.Msg)
		return nil, errors.New("failed to encrypt, error:" + out.Msg)
	}

	in = scryinfo.CipherParameter{Message: out.Data, Address: address2}
	out, _ = am.client.ContentEncrypt(context.Background(), &in)
	if out.Status != scryinfo.Status_OK {
		rlog.Error("failed to encrypt, error:", out.Msg)
		return nil, errors.New("failed to encrypt, error:" + out.Msg)
	}

	return out.Data, nil
}

func (am AccountManager) Decrypt(cipherText []byte, address string, password string)([]byte, error) {

	defer func() {
		if err := recover(); err != nil {
			rlog.Error("failed to decrypt, error:", err)
		}
	}()

	if am.client == nil {
		return nil, errors.New("failed to decrypt, error: null account service")
	}

	in := scryinfo.CipherParameter{
		Message: cipherText,
		Address: address,
		Password: password,
	}
	out, _ := am.client.ContentDecrypt(context.Background(), &in)
	if out.Status != scryinfo.Status_OK {
		rlog.Error("failed to decrypt, error:", out.Msg)
		return nil, errors.New("failed to decrypt, error:" + out.Msg)
	}

	return out.Data, nil
}

func (am AccountManager) SignTransaction(message []byte, address string, password string) ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			rlog.Error("failed to encrypt, error:", err)
		}
	}()

	if am.client == nil {
		return nil, errors.New("failed to encrypt, error: null client")
	}

	in := scryinfo.CipherParameter{
		Message:  message,
		Address:  address,
		Password: password,
	}

	out, _ := am.client.Signature(context.Background(), &in)
	if out.Status != scryinfo.Status_OK {
		rlog.Error("failed to signature, error:", out.Msg)
		return nil, errors.New("failed to signature, error:" + out.Msg)
	}

	return out.Data, nil
}

func (am AccountManager) ImportAccount(keyJson []byte, oldPassword string, newPassword string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			rlog.Error("failed to import account, error:", err)
		}
	}()

	if am.client == nil {
		return "", errors.New("failed to import account, null account service")
	}

	in := scryinfo.ImportParameter{ContentPassword: oldPassword, ImportPsd: newPassword, Content: keyJson}
	out, _ := am.client.ImportKeystore(context.Background(), &in)
	if out.Status != scryinfo.Status_OK {
		rlog.Error("failed to import account, error:", out.Msg)
		return "", errors.New("failed to import account, error:" + out.Msg)
	}

	return out.Address, nil
}
