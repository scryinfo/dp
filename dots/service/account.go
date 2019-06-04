// Scry Info.  All rights reserved.
// license that can be found in the license file.

package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/scryinfo/dot/dot"
	"github.com/scryinfo/dp/dots/binary/sdk/interface/account"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type AccountManager struct {
	client   account.KeyServiceClient
	accounts []*Account
	config   amconfig
}

type Account struct {
	Address string
}

type amconfig struct {
	NodeAddr string `json:"nodeAddr"`
}

const (
	AccTypeId = "cf8856e3-75aa-4ea5-b3bf-0e9eec347e56"
	AccLiveId = "cf8856e3-75aa-4ea5-b3bf-0e9eec347e56"
)

func (c *AccountManager) Create(l dot.Line) error {
	cn, err := grpc.Dial(c.config.NodeAddr, grpc.WithInsecure())
	if err != nil {
		dot.Logger().Errorln("failed to connect to node:" + c.config.NodeAddr)
		return err
	}

	c.client = account.NewKeyServiceClient(cn)
	if c.client == nil {
		return errors.New("failed to create account service client")
	}

	return nil
}

func GetAMIns() *AccountManager {
	logger := dot.Logger()
	l := dot.GetDefaultLine()
	if l == nil {
		logger.Errorln("the line do not create, do not call it")
		return nil
	}
	d, err := l.ToInjecter().GetByLiveId(AccLiveId)
	if err != nil {
		logger.Errorln(err.Error())
		return nil
	}

	if g, ok := d.(*AccountManager); ok {
		return g
	}

	logger.Errorln("do not get the AM dot")
	return nil
}

func GetAMConfig() *amconfig {
	logger := dot.Logger()
	l := dot.GetDefaultLine()
	if l == nil {
		logger.Errorln("the line do not create, do not call it")
		return nil
	}
	d, err := l.ToInjecter().GetByLiveId(AccLiveId)
	if err != nil {
		logger.Errorln(err.Error())
		return nil
	}

	if g, ok := d.(*AccountManager); ok {
		return &g.config
	}

	logger.Errorln("do not get the AM dot")
	return nil
}

func newAccountDot(conf interface{}) (dot.Dot, error) {
	var err error
	var bs []byte
	if bt, ok := conf.([]byte); ok {
		bs = bt
	} else {
		return nil, dot.SError.Parameter
	}

	dConf := &amconfig{}
	err = dot.UnMarshalConfig(bs, dConf)
	if err != nil {
		return nil, err
	}

	d := &AccountManager{config: *dConf}

	return d, err
}

func AccountTypeLive() *dot.TypeLives {
	return &dot.TypeLives{
		Meta: dot.Metadata{TypeId: AccTypeId,
			NewDoter: func(conf interface{}) (dot dot.Dot, err error) {
				return newAccountDot(conf)
			}},
	}
}

func (c *AccountManager) CreateAccount(password string) (*Account, error) {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("failed to create account. ", zap.Any("", er))
		}
	}()

	if c.client == nil {
		return nil, errors.New("failed to create account, error: null account service")
	}

	addr, err := c.client.GenerateAddress(context.Background(), &account.AddressParameter{Password: password})
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

	newAccount := &Account{Address: addr.Address}
	c.accounts = append(c.accounts, newAccount)

	return newAccount, nil
}

func (c AccountManager) AuthAccount(address string, password string) (bool, error) {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("failed to authenticate account. ", zap.Any("", er))
		}
	}()

	if c.client == nil {
		return false, errors.New("failed to authenticate account, error: null account service")
	}

	addr, err := c.client.VerifyAddress(context.Background(),
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

func (c AccountManager) GetAccounts() []*Account {
	return c.accounts
}

func (c AccountManager) AccountValid(address string) bool {
	for _, v := range c.accounts {
		if v.Address == address {
			return true
		}
	}

	return false
}

func (c AccountManager) Encrypt(
	plainText []byte,
	address string,
) ([]byte, error) {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("failed to encrypt. ", zap.Any("", er))
		}
	}()

	if c.client == nil {
		return nil, errors.New("failed to encrypt, error: null account service")
	}

	in := account.CipherParameter{Message: plainText, Address: address}
	out, err := c.client.ContentEncrypt(context.Background(), &in)
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

func (c AccountManager) Decrypt(
	cipherText []byte,
	address string,
	password string) ([]byte, error) {

	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("failed to decrypt. ", zap.Any("", er))
		}
	}()

	if c.client == nil {
		return nil, errors.New("failed to decrypt, error: null account service")
	}

	in := account.CipherParameter{
		Message:  cipherText,
		Address:  address,
		Password: password,
	}
	out, err := c.client.ContentDecrypt(context.Background(), &in)
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

func (c AccountManager) ReEncrypt(
	cipherText []byte,
	address1 string,
	address2 string,
	password string,
) ([]byte, error) {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("failed to reencrypt. ", zap.Any("", er))
		}
	}()

	if c.client == nil {
		return nil, errors.New("failed to encrypt, error: null account service")
	}

	in := account.CipherParameter{Message: cipherText, Address: address1, Password: password}
	out, err := c.client.ContentDecrypt(context.Background(), &in)
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
	out, err = c.client.ContentEncrypt(context.Background(), &in)
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

func (c AccountManager) SignTransaction(message []byte, address string, password string) ([]byte, error) {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("failed to encrypt. ", zap.Any("", er))
		}
	}()

	if c.client == nil {
		return nil, errors.New("failed to encrypt, error: null client")
	}

	in := account.CipherParameter{
		Message:  message,
		Address:  address,
		Password: password,
	}

	out, err := c.client.Signature(context.Background(), &in)
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

func (c AccountManager) ImportAccount(
	keyJson []byte,
	oldPassword string,
	newPassword string) (string, error) {
	defer func() {
		if er := recover(); er != nil {
			dot.Logger().Errorln("failed to import account. ", zap.Any("", er))
		}
	}()

	if c.client == nil {
		return "", errors.New("failed to import account, null account service")
	}

	in := account.ImportParameter{ContentPassword: oldPassword, ImportPsd: newPassword, Content: keyJson}
	out, err := c.client.ImportKeystore(context.Background(), &in)
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
