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
    cn *grpc.ClientConn
    client authStub.KeyServiceClient
}

type UserAccount struct {
    Addr string
}

func GetAccIns() *Account {
    logger := dot.Logger()
    l := dot.GetDefaultLine()
    if l == nil {
        logger.Errorln("the line do not create, do not call it")
        return nil
    }
    d, err := l.ToInjecter().GetByLiveId(AccountTypeId)
    if err != nil {
        logger.Errorln(err.Error())
        return nil
    }

    if g, ok := d.(*Account); ok {
        return g
    }

    logger.Errorln("do not get the Account dot")
    return nil
}

//construct dot
func newAccountDot(_ interface{}) (dot.Dot, error) {
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

func (c *Account) Destroy(ignore bool) error {
    dot.Logger().Debugln("closing grpc connection...")
    err := c.cn.Close()
    if err != nil {
        dot.Logger().Errorln("failed to close grpc connection", zap.Error(err))
        return err
    }

    dot.Logger().Debugln("grpc connection closed")

    return nil
}


func (c *Account) Initialize(authServiceAddr string) error {
    var err error
    c.cn, err = grpc.Dial(authServiceAddr, grpc.WithInsecure())
    if err != nil {
        dot.Logger().Errorln("failed to Connect to node:" + authServiceAddr)
        return err
    }

    c.client = authStub.NewKeyServiceClient(c.cn)
    if c.client == nil {
        dot.Logger().Errorln("failed to create interface service client")
        return errors.New("failed to create interface service client")
    }

    return nil
}

func (c *Account) CreateUserAccount(password string) (*UserAccount, error) {
    defer func() {
        if er := recover(); er != nil {
            dot.Logger().Errorln("", zap.Any("failed to create user account, error:", er))
        }
    }()

    if c.client == nil {
        return nil, errors.New("failed to create user account, error: null grpc client")
    }

    addr, err := c.client.GenerateAddress(
        context.Background(),
        &authStub.AddressParameter{Password: password},
    )

    if err != nil {
        err = errors.Wrap(err, "failed to create user account")
    } else if addr != nil && addr.Status != authStub.Status_OK {
        err = errors.New("failed to create user account, status is not ok, error:" + addr.Msg)
    } else if addr == nil {
        err = errors.New("failed to create user account, returned address is null")
    }

    if err != nil {
        dot.Logger().Errorln("failed to create user account", zap.Error(err))
        return nil, err
    }

    newAccount := &UserAccount{addr.Address}

    return newAccount, nil
}

func (c *Account) AuthUserAccount(address string, password string) (bool, error) {
    defer func() {
        if er := recover(); er != nil {
            dot.Logger().Errorln("", zap.Any("failed to authenticate, error:", er))
        }
    }()

    if c.client == nil {
        return false, errors.New("failed to authenticate interface, error: null grpc client")
    }

    addr, err := c.client.VerifyAddress(
        context.Background(),
        &authStub.AddressParameter{Password: password, Address: address},
    )
    if err != nil {
        err = errors.Wrap(err, "failed to authenticate user account")
    } else if addr == nil {
        err = errors.New("failed to authenticate user account, returned address is null")
    }
    if err != nil {
        dot.Logger().Errorln("failed to authenticate user account", zap.Error(err))
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
        return nil, errors.New("failed to encrypt, error: client is null")
    }

    in := authStub.CipherParameter{Message: plainText, Address: address}


    dot.Logger().Debugln("Node: show encrypt params. ", zap.Any("in", in))


    out, err := c.client.ContentEncrypt(context.Background(), &in)
    if err != nil {
        err = errors.Wrap(err, "failed to encrypt data")
    } else if out == nil {
        err = errors.New("failed to encrypt data, error: result is null")
    } else if out.Status != authStub.Status_OK {
        err = errors.New("failed to encrypt data, status is not okk, error:" + out.Msg)
    }
    if err != nil {
        dot.Logger().Errorln("failed to encrypt data", zap.Error(err))
        return nil, err
    }

    return out.Data, nil
}

func (c *Account) Decrypt(
    cipherText []byte,
    address string,
    password string,
) ([]byte, error) {

    defer func() {
        if er := recover(); er != nil {
            dot.Logger().Errorln("", zap.Any("failed to decrypt, error:", er))
        }
    }()

    if c.client == nil {
        return nil, errors.New("failed to decrypt, error: client is null")
    }

    in := authStub.CipherParameter{
        Message:  cipherText,
        Address:  address,
        Password: password,
    }


    dot.Logger().Debugln("Node: show decrypt params. ", zap.Any("in", in))


    out, err := c.client.ContentDecrypt(context.Background(), &in)
    if err != nil {
        err = errors.Wrap(err, "failed to decrypt data")
    } else if out == nil {
        err = errors.New("failed to encrypt data, error: result is null")
    } else if out.Status != authStub.Status_OK {
        err = errors.New("failed to decrypt, status is not ok, error:" + out.Msg)
    }
    if err != nil {
        dot.Logger().Errorln("failed to decrypt", zap.Error(err))
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
            dot.Logger().Errorln("", zap.Any("failed to re-encrypt, error:", er))
        }
    }()

    if c.client == nil {
        return nil, errors.New("failed to re-encrypt, error: client is null")
    }

    in := authStub.CipherParameter{Message: cipherText, Address: address1, Password: password}
    out, err := c.client.ContentDecrypt(context.Background(), &in)
    if err != nil {
        err = errors.Wrap(err, "failed to encrypt data")
    } else if out == nil {
        err = errors.New("failed to encrypt data, error: result is null")
    } else if out.Status != authStub.Status_OK {
        err = errors.New("failed to decrypt, status is not ok, error:" + out.Msg)
    }
    if err != nil {
        dot.Logger().Errorln("failed to decrypt", zap.Error(err))
        return nil, err
    }

    in = authStub.CipherParameter{Message: out.Data, Address: address2}
    out, err = c.client.ContentEncrypt(context.Background(), &in)
    if err != nil {
        err = errors.Wrap(err, "failed to encrypt data")
    } else if out == nil {
        err = errors.New("failed to encrypt data, error: result is null")
    } else if out.Status != authStub.Status_OK {
        err = errors.New("failed to encrypt data, status is not ok, error:" + out.Msg)
    }
    if err != nil {
        dot.Logger().Errorln("failed to encrypt data", zap.Error(err))
        return nil, err
    }

    return out.Data, nil
}

func (c *Account) SignTransaction(message []byte, address string, password string) ([]byte, error) {
    defer func() {
        if er := recover(); er != nil {
            dot.Logger().Errorln("", zap.Any("failed to signature transaction, error:", er))
        }
    }()

    if c.client == nil {
        return nil, errors.New("failed to signature transaction, error: client is null")
    }

    in := authStub.CipherParameter{
        Message:  message,
        Address:  address,
        Password: password,
    }

    out, err := c.client.Signature(context.Background(), &in)
    if err != nil {
        err = errors.Wrap(err, "failed to signature transaction, error:")
    } else if out == nil {
        err = errors.New("failed to signature transaction, error: result is nil")
    } else if out.Status != authStub.Status_OK {
        err = errors.New("failed to signature transaction, status is not ok, error:" + out.Msg)
    }
    if err != nil {
        dot.Logger().Errorln("failed to signature transaction", zap.Error(err))
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
            dot.Logger().Errorln("", zap.Any("failed to import user account, error:", er))
        }
    }()

    if c.client == nil {
        return "", errors.New("failed to import user account, client is null")
    }

    in := authStub.ImportParameter{ContentPassword: oldPassword, ImportPsd: newPassword, Content: keyJson}
    out, err := c.client.ImportKeystore(context.Background(), &in)
    if err != nil {
        err = errors.Wrap(err, "failed to import user account, error:")
    } else if out == nil {
        err = errors.New("failed to import user account, error: result is nil")
    } else if out.Status != authStub.Status_OK {
        err = errors.New("failed to import user account, error:" + out.Msg)
    }
    if err != nil {
        dot.Logger().Errorln("failed to import user account", zap.Error(err))
        return "", err
    }

    return out.Address, nil
}
