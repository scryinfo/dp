// Scry Info.  All rights reserved.
// license that can be found in the license file.

package accounts

import (
	"context"
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dp/dots/binary/sdk/interface/account"
	"go.uber.org/zap"
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
	client   account.KeyServiceClient
	accounts []*Account
}

func (am *AccountManager) Initialize(asNodeAddr string) error {
	cn, err := grpc.Dial(asNodeAddr, grpc.WithInsecure())
	if err != nil {
		dot.Logger().Errorln("failed to WSConnect to node:" + asNodeAddr)
		return err
	}

	am.client = account.NewKeyServiceClient(cn)
	if am.client == nil {
		return errors.New("failed to create account service client")
	}

	return nil
}

func (am *AccountManager) CreateAccount(password string) (*Account, error) {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("", zap.Any("failed to create account, error:", er))
		}
	}()

	if am.client == nil {
		return nil, errors.New("failed to create account, error: null account service")
	}

	addr, err := am.client.GenerateAddress(context.Background(), &account.AddressParameter{Password: password})
	if err != nil {
		err = errors.Wrap(err, "failed to create account, error:")
	} else if addr != nil && addr.Status != account.Status_OK {
		err = errors.New("failed to create account, error:" + addr.Msg)
	} else if addr == nil {
		err = errors.New("failed to create account, error: addr is nil")
	}
	if err != nil {
		dot.Logger().Errorln("", zap.NamedError("", err))
		return nil, err
	}

	newAccount := &Account{addr.Address}
	am.accounts = append(am.accounts, newAccount)

	return newAccount, nil
}

func (am AccountManager) AuthAccount(address string, password string) (bool, error) {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("", zap.Any("failed to authenticate account, error:", er))
		}
	}()

	if am.client == nil {
		return false, errors.New("failed to authenticate account, error: null account service")
	}

	addr, err := am.client.VerifyAddress(context.Background(),
		&account.AddressParameter{Password: password, Address: address})
	if err != nil {
		err = errors.Wrap(err, "failed to authenticate user, error:")
	} else if addr == nil {
		err = errors.New("failed to authenticate user, error: addr is nil")
	}
	if err != nil {
		dot.Logger().Errorln("", zap.NamedError("", err))
		return false, err
	}

	rv := addr.Status == account.Status_OK

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
	address string,
) ([]byte, error) {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("", zap.Any("failed to encrypt, error:", er))
		}
	}()

	if am.client == nil {
		return nil, errors.New("failed to encrypt, error: null account service")
	}

	in := account.CipherParameter{Message: plainText, Address: address}
	out, err := am.client.ContentEncrypt(context.Background(), &in)
	if err != nil {
		err = errors.Wrap(err, "failed to encrypt data, error:")
	} else if out == nil {
		err = errors.New("failed to encrypt data, error: addr is nil")
	} else if out.Status != account.Status_OK {
		err = errors.New("failed to encrypt data, error:" + out.Msg)
	}
	if err != nil {
		dot.Logger().Errorln("", zap.NamedError("", err))
		return nil, err
	}

	return out.Data, nil
}

func (am AccountManager) Decrypt(
	cipherText []byte,
	address string,
	password string) ([]byte, error) {

	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("", zap.Any("failed to decrypt, error:", er))
		}
	}()

	if am.client == nil {
		return nil, errors.New("failed to decrypt, error: null account service")
	}

	in := account.CipherParameter{
		Message:  cipherText,
		Address:  address,
		Password: password,
	}
	out, err := am.client.ContentDecrypt(context.Background(), &in)
	if err != nil {
		err = errors.Wrap(err, "failed to decrypt data, error:")
	} else if out == nil {
		err = errors.New("failed to encrypt data, error: addr is nil")
	} else if out.Status != account.Status_OK {
		err = errors.New("failed to decrypt, error:" + out.Msg)
	}
	if err != nil {
		dot.Logger().Errorln("", zap.Any("", err))
		return nil, err
	}

	return out.Data, nil
}

func (am AccountManager) ReEncrypt(
	cipherText []byte,
	address1 string,
	address2 string,
	password string,
) ([]byte, error) {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("", zap.Any("failed to reencrypt, error:", er))
		}
	}()

	if am.client == nil {
		return nil, errors.New("failed to encrypt, error: null account service")
	}

	in := account.CipherParameter{Message: cipherText, Address: address1, Password: password}
	out, err := am.client.ContentDecrypt(context.Background(), &in)
	if err != nil {
		err = errors.Wrap(err, "failed to encrypt data, error:")
	} else if out == nil {
		err = errors.New("failed to encrypt data, error: addr is nil")
	} else if out.Status != account.Status_OK {
		err = errors.New("failed to decrypt, error:" + out.Msg)
	}
	if err != nil {
		dot.Logger().Errorln("", zap.NamedError("", err))
		return nil, err
	}

	in = account.CipherParameter{Message: out.Data, Address: address2}
	out, err = am.client.ContentEncrypt(context.Background(), &in)
	if err != nil {
		err = errors.Wrap(err, "failed to encrypt data, error:")
	} else if out == nil {
		err = errors.New("failed to encrypt data, error: addr is nil")
	} else if out.Status != account.Status_OK {
		err = errors.New("failed to encrypt data, error:" + out.Msg)
	}
	if err != nil {
		dot.Logger().Errorln("", zap.NamedError("", err))
		return nil, err
	}

	return out.Data, nil
}

func (am AccountManager) SignTransaction(message []byte, address string, password string) ([]byte, error) {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("", zap.Any("failed to encrypt, error:", er))
		}
	}()

	if am.client == nil {
		return nil, errors.New("failed to encrypt, error: null client")
	}

	in := account.CipherParameter{
		Message:  message,
		Address:  address,
		Password: password,
	}

	out, err := am.client.Signature(context.Background(), &in)
	if err != nil {
		err = errors.Wrap(err, "failed to signature, error:")
	} else if out == nil {
		err = errors.New("failed to signature, error: addr is nil")
	} else if out.Status != account.Status_OK {
		err = errors.New("failed to signature, error:" + out.Msg)
	}
	if err != nil {
		dot.Logger().Errorln("", zap.NamedError("", err))
		return nil, err
	}

	return out.Data, nil
}

func (am AccountManager) ImportAccount(
	keyJson []byte,
	oldPassword string,
	newPassword string) (string, error) {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("", zap.Any("failed to import account, error:", er))
		}
	}()

	if am.client == nil {
		return "", errors.New("failed to import account, null account service")
	}

	in := account.ImportParameter{ContentPassword: oldPassword, ImportPsd: newPassword, Content: keyJson}
	out, err := am.client.ImportKeystore(context.Background(), &in)
	if err != nil {
		err = errors.Wrap(err, "failed to import account, error:")
	} else if out == nil {
		err = errors.New("failed to import account, error: addr is nil")
	} else if out.Status != account.Status_OK {
		err = errors.New("failed to import account, error:" + out.Msg)
	}
	if err != nil {
		dot.Logger().Errorln("", zap.NamedError("", err))
		return "", err
	}

	return out.Address, nil
}
