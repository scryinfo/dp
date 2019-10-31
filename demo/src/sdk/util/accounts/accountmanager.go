// Scry Info.  All rights reserved.
// license that can be found in the license file.

package accounts

import (
	"context"
	"github.com/pkg/errors"
	"github.com/scryinfo/dp/demo/src/sdk/interface/accountinterface"
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
		rlog.Error("failed to WSConnect to node:" + asNodeAddr)
		return err
	}

	am.client = scryinfo.NewKeyServiceClient(cn)
	if am.client == nil {
		return errors.New("failed to create interface service client")
	}

	return nil
}

func (am *AccountManager) CreateAccount(password string) (*Account, error) {

	defer func() {
		if err := recover(); err != nil {
			rlog.Error("failed to create interface, error:", err)
		}
	}()

	if am.client == nil {
		return nil, errors.New("failed to create interface, error: null interface service")
	}

	addr, err := am.client.GenerateAddress(context.Background(), &scryinfo.AddressParameter{Password: password})
	if err != nil {
		rlog.Error("failed to create interface, error:", err)
		return nil, errors.New("failed to create interface, error:" + err.Error())
	} else if addr != nil && addr.Status != scryinfo.Status_OK {
		rlog.Error("failed to create interface, error:", addr.Msg)
		return nil, errors.New("failed to create interface, error:" + addr.Msg)
	} else if addr == nil {
		rlog.Error("failed to create interface, error: addr is nil")
		return nil, errors.New("failed to create interface, error: addr is nil")
	}

	newAccount := &Account{addr.Address}
	am.accounts = append(am.accounts, newAccount)

	return newAccount, nil
}

func (am AccountManager) AuthAccount(address string, password string) (bool, error) {
	defer func() {
		if err := recover(); err != nil {
			rlog.Error("failed to authenticate interface, error:", err)
		}
	}()

	if am.client == nil {
		return false, errors.New("failed to authenticate interface, error: null interface service")
	}

	addr, err := am.client.VerifyAddress(context.Background(),
		&scryinfo.AddressParameter{Password: password, Address: address})
	if err != nil {
		rlog.Error("failed to authenticate user, error:", err)
		return false, errors.New("failed to authenticate user, error:" + err.Error())
	} else if addr == nil {
		rlog.Error("failed to authenticate user, error: addr is nil")
		return false, errors.New("failed to authenticate user, error: addr is nil")
	}

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

func (am AccountManager) Encrypt(
	plainText []byte,
	address string) ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			rlog.Error("failed to encrypt, error:", err)
		}
	}()

	if am.client == nil {
		return nil, errors.New("failed to encrypt, error: null interface service")
	}

	in := scryinfo.CipherParameter{Message: plainText, Address: address}
	out, err := am.client.ContentEncrypt(context.Background(), &in)
	if err != nil {
		rlog.Error("failed to encrypt data, error:", err)
		return nil, errors.New("failed to encrypt data, error:" + err.Error())
	} else if out == nil {
		rlog.Error("failed to encrypt data, error: addr is nil")
		return nil, errors.New("failed to encrypt data, error: addr is nil")
	} else if out.Status != scryinfo.Status_OK {
		rlog.Error("failed to encrypt data, error:", out.Msg)
		return nil, errors.New("failed to encrypt data, error:" + out.Msg)
	}

	return out.Data, nil
}

func (am AccountManager) Decrypt(
	cipherText []byte,
	address string,
	password string) ([]byte, error) {

	defer func() {
		if err := recover(); err != nil {
			rlog.Error("failed to decrypt, error:", err)
		}
	}()

	if am.client == nil {
		return nil, errors.New("failed to decrypt, error: null interface service")
	}

	in := scryinfo.CipherParameter{
		Message:  cipherText,
		Address:  address,
		Password: password,
	}
	out, err := am.client.ContentDecrypt(context.Background(), &in)
	if err != nil {
		rlog.Error("failed to decrypt data, error:", err)
		return nil, errors.New("failed to encrypt data, error:" + err.Error())
	} else if out == nil {
		rlog.Error("failed to decrypt data, error: addr is nil")
		return nil, errors.New("failed to encrypt data, error: addr is nil")
	} else if out.Status != scryinfo.Status_OK {
		rlog.Error("failed to decrypt, error:", out.Msg)
		return nil, errors.New("failed to decrypt, error:" + out.Msg)
	}

	return out.Data, nil
}

func (am AccountManager) ReEncrypt(
	cipherText []byte,
	address1 string,
	address2 string,
	password string) ([]byte, error) {
	defer func() {
		if err := recover(); err != nil {
			rlog.Error("failed to reencrypt, error:", err)
		}
	}()

	if am.client == nil {
		return nil, errors.New("failed to encrypt, error: null interface service")
	}

	in := scryinfo.CipherParameter{Message: cipherText, Address: address1, Password: password}
	out, err := am.client.ContentDecrypt(context.Background(), &in)
	if err != nil {
		rlog.Error("failed to decrypt data, error:", err)
		return nil, errors.New("failed to encrypt data, error:" + err.Error())
	} else if out == nil {
		rlog.Error("failed to decrypt data, error: addr is nil")
		return nil, errors.New("failed to encrypt data, error: addr is nil")
	} else if out.Status != scryinfo.Status_OK {
		rlog.Error("failed to decrypt, error:", out.Msg)
		return nil, errors.New("failed to decrypt, error:" + out.Msg)
	}

	in = scryinfo.CipherParameter{Message: out.Data, Address: address2}
	out, err = am.client.ContentEncrypt(context.Background(), &in)
	if err != nil {
		rlog.Error("failed to encrypt data, error:", err)
		return nil, errors.New("failed to encrypt data, error:" + err.Error())
	} else if out == nil {
		rlog.Error("failed to encrypt data, error: addr is nil")
		return nil, errors.New("failed to encrypt data, error: addr is nil")
	} else if out.Status != scryinfo.Status_OK {
		rlog.Error("failed to encrypt data, error:", out.Msg)
		return nil, errors.New("failed to encrypt data, error:" + out.Msg)
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

	out, err := am.client.Signature(context.Background(), &in)
	if err != nil {
		rlog.Error("failed to signature, error:", err)
		return nil, errors.New("failed to signature, error:" + err.Error())
	} else if out == nil {
		rlog.Error("failed to signature, error: addr is nil")
		return nil, errors.New("failed to signature, error: addr is nil")
	} else if out.Status != scryinfo.Status_OK {
		rlog.Error("failed to signature, error:", out.Msg)
		return nil, errors.New("failed to signature, error:" + out.Msg)
	}

	return out.Data, nil
}

func (am AccountManager) ImportAccount(
	keyJson []byte,
	oldPassword string,
	newPassword string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			rlog.Error("failed to import interface, error:", err)
		}
	}()

	if am.client == nil {
		return "", errors.New("failed to import interface, null interface service")
	}

	in := scryinfo.ImportParameter{ContentPassword: oldPassword, ImportPsd: newPassword, Content: keyJson}
	out, err := am.client.ImportKeystore(context.Background(), &in)
	if err != nil {
		rlog.Error("failed to import interface, error:", err)
		return "", errors.New("failed to import interface, error:" + err.Error())
	} else if out == nil {
		rlog.Error("failed to import interface, error: addr is nil")
		return "", errors.New("failed to import interface, error: addr is nil")
	} else if out.Status != scryinfo.Status_OK {
		rlog.Error("failed to import interface, error:", out.Msg)
		return "", errors.New("failed to import interface, error:" + out.Msg)
	}

	return out.Address, nil
}
