package auth

import (
    "context"
    "github.com/pkg/errors"
    "github.com/scryinfo/dot/dot"
    authStub "github.com/scryinfo/dp/dots/auth/stub"
    "go.uber.org/zap"
    "google.golang.org/grpc"
)

const (
    AccountTypeId = "ca1c6ce4-182b-430a-9813-caeccf83f8ab"
)

type Account struct {
    client authStub.KeyServiceClient
}

type UserAccount struct {
    Addr string
}

//construct dot
func newAccountDot(conf interface{}) (dot.Dot, error) {
    d := &Account{}
    return d, nil
}

//Data structure needed when generating newer component
func AccountTypeLive() *dot.TypeLives {
    return &dot.TypeLives{
            Meta: dot.Metadata{TypeId: AccountTypeId, NewDoter: func(conf interface{}) (dot.Dot, error) {
                return newAccountDot(conf)
            },
       },
    }
}

func (c *Account) Create(l dot.Line) error {

    return nil
}

func (c *Account) Start(ignore bool) error {
    return nil
}

func (c *Account) Stop(ignore bool) error {
    return nil
}

func (c *Account) Destroy(ignore bool) error {
    return nil
}


func (c *Account) Initialize(authServiceAddr string) error {
    cn, err := grpc.Dial(authServiceAddr, grpc.WithInsecure())
    if err != nil {
        dot.Logger().Errorln("failed to Connect to node:" + authServiceAddr)
        return err
    }

    c.client = authStub.NewKeyServiceClient(cn)
    if c.client == nil {
        dot.Logger().Errorln("failed to create interface service client")
        return errors.New("failed to create interface service client")
    }

    return nil
}

func (c *Account) CreateUserAccount(password string) (*UserAccount, error) {
    defer func() {
        if er := recover(); er != nil {
            dot.Logger().Errorln("", zap.Any("failed to create interface, error:", er))
        }
    }()

    if c.client == nil {
        return nil, errors.New("failed to create interface, error: null interface service")
    }

    addr, err := c.client.GenerateAddress(context.Background(), &authStub.AddressParameter{Password: password})
    if err != nil {
        err = errors.Wrap(err, "failed to create interface, error:")
    } else if addr != nil && addr.Status != authStub.Status_OK {
        err = errors.New("failed to create interface, error:" + addr.Msg)
    } else if addr == nil {
        err = errors.New("failed to create interface, error: addr is nil")
    }
    if err != nil {
        dot.Logger().Errorln("", zap.NamedError("", err))
        return nil, err
    }

    newAccount := &UserAccount{addr.Address}

    return newAccount, nil
}

func (c *Account) AuthUserAccount(address string, password string) (bool, error) {
    defer func() {
        if er := recover(); er != nil {
            dot.Logger().Errorln("", zap.Any("failed to authenticate interface, error:", er))
        }
    }()

    if c.client == nil {
        return false, errors.New("failed to authenticate interface, error: null interface service")
    }

    addr, err := c.client.VerifyAddress(context.Background(),
        &authStub.AddressParameter{Password: password, Address: address})
    if err != nil {
        err = errors.Wrap(err, "failed to authenticate user, error:")
    } else if addr == nil {
        err = errors.New("failed to authenticate user, error: addr is nil")
    }
    if err != nil {
        dot.Logger().Errorln("", zap.NamedError("", err))
        return false, err
    }

    rv := addr.Status == authStub.Status_OK

    return rv, nil
}


func (c *Account) Encrypt(
    plainText []byte,
    address string,
) ([]byte, error) {
    defer func() {
        if er := recover(); er != nil {
            dot.Logger().Errorln("", zap.Any("failed to encrypt, error:", er))
        }
    }()

    if c.client == nil {
        return nil, errors.New("failed to encrypt, error: null interface service")
    }

    in := authStub.CipherParameter{Message: plainText, Address: address}
    out, err := c.client.ContentEncrypt(context.Background(), &in)
    if err != nil {
        err = errors.Wrap(err, "failed to encrypt data, error:")
    } else if out == nil {
        err = errors.New("failed to encrypt data, error: addr is nil")
    } else if out.Status != authStub.Status_OK {
        err = errors.New("failed to encrypt data, error:" + out.Msg)
    }
    if err != nil {
        dot.Logger().Errorln("", zap.NamedError("", err))
        return nil, err
    }

    return out.Data, nil
}

func (c *Account) Decrypt(
    cipherText []byte,
    address string,
    password string) ([]byte, error) {

    defer func() {
        if er := recover(); er != nil {
            dot.Logger().Errorln("", zap.Any("failed to decrypt, error:", er))
        }
    }()

    if c.client == nil {
        return nil, errors.New("failed to decrypt, error: null interface service")
    }

    in := authStub.CipherParameter{
        Message:  cipherText,
        Address:  address,
        Password: password,
    }
    out, err := c.client.ContentDecrypt(context.Background(), &in)
    if err != nil {
        err = errors.Wrap(err, "failed to decrypt data, error:")
    } else if out == nil {
        err = errors.New("failed to encrypt data, error: addr is nil")
    } else if out.Status != authStub.Status_OK {
        err = errors.New("failed to decrypt, error:" + out.Msg)
    }
    if err != nil {
        dot.Logger().Errorln("", zap.Any("", err))
        return nil, err
    }

    return out.Data, nil
}

func (c *Account) ReEncrypt(
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

    if c.client == nil {
        return nil, errors.New("failed to encrypt, error: null interface service")
    }

    in := authStub.CipherParameter{Message: cipherText, Address: address1, Password: password}
    out, err := c.client.ContentDecrypt(context.Background(), &in)
    if err != nil {
        err = errors.Wrap(err, "failed to encrypt data, error:")
    } else if out == nil {
        err = errors.New("failed to encrypt data, error: addr is nil")
    } else if out.Status != authStub.Status_OK {
        err = errors.New("failed to decrypt, error:" + out.Msg)
    }
    if err != nil {
        dot.Logger().Errorln("", zap.NamedError("", err))
        return nil, err
    }

    in = authStub.CipherParameter{Message: out.Data, Address: address2}
    out, err = c.client.ContentEncrypt(context.Background(), &in)
    if err != nil {
        err = errors.Wrap(err, "failed to encrypt data, error:")
    } else if out == nil {
        err = errors.New("failed to encrypt data, error: addr is nil")
    } else if out.Status != authStub.Status_OK {
        err = errors.New("failed to encrypt data, error:" + out.Msg)
    }
    if err != nil {
        dot.Logger().Errorln("", zap.NamedError("", err))
        return nil, err
    }

    return out.Data, nil
}

func (c *Account) SignTransaction(message []byte, address string, password string) ([]byte, error) {
    defer func() {
        if er := recover(); er != nil {
            dot.Logger().Errorln("", zap.Any("failed to encrypt, error:", er))
        }
    }()

    if c.client == nil {
        return nil, errors.New("failed to encrypt, error: null client")
    }

    in := authStub.CipherParameter{
        Message:  message,
        Address:  address,
        Password: password,
    }

    out, err := c.client.Signature(context.Background(), &in)
    if err != nil {
        err = errors.Wrap(err, "failed to signature, error:")
    } else if out == nil {
        err = errors.New("failed to signature, error: addr is nil")
    } else if out.Status != authStub.Status_OK {
        err = errors.New("failed to signature, error:" + out.Msg)
    }
    if err != nil {
        dot.Logger().Errorln("", zap.NamedError("", err))
        return nil, err
    }

    return out.Data, nil
}

func (c *Account) ImportUserAccount(
    keyJson []byte,
    oldPassword string,
    newPassword string,
) (string, error) {
    defer func() {
        if er := recover(); er != nil {
            dot.Logger().Errorln("", zap.Any("failed to import interface, error:", er))
        }
    }()

    if c.client == nil {
        return "", errors.New("failed to import interface, null interface service")
    }

    in := authStub.ImportParameter{ContentPassword: oldPassword, ImportPsd: newPassword, Content: keyJson}
    out, err := c.client.ImportKeystore(context.Background(), &in)
    if err != nil {
        err = errors.Wrap(err, "failed to import interface, error:")
    } else if out == nil {
        err = errors.New("failed to import interface, error: addr is nil")
    } else if out.Status != authStub.Status_OK {
        err = errors.New("failed to import interface, error:" + out.Msg)
    }
    if err != nil {
        dot.Logger().Errorln("", zap.NamedError("", err))
        return "", err
    }

    return out.Address, nil
}
