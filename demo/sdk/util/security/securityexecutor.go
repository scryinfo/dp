package security

import (
    "github.com/ethereum/go-ethereum/accounts/keystore"
    "github.com/ethereum/go-ethereum/crypto"
)

type CryptExecutor struct {
}

func (ce CryptExecutor) SignTransaction(message []byte, pubkey string, password string) ([]byte, error) {
    //hardcode here, will call jeremy's library later
    keyJson := `{"version":3,"id":"80d7b778-e617-4b35-bb09-f4b224984ed6","address":"d280b60c38bc8db9d309fa5a540ffec499f0a3e8","crypto":{"ciphertext":"58ac20c29dd3029f4d374839508ba83fc84628ae9c3f7e4cc36b05e892bf150d","cipherparams":{"iv":"9ab7a5f9bcc9df7d796b5022023e2d14"},"cipher":"aes-128-ctr","kdf":"scrypt","kdfparams":{"dklen":32,"salt":"63a364b8a64928843708b5e9665a79fa00890002b32833b3a9ff99eec78dbf81","n":262144,"r":8,"p":1},"mac":"3a38f91234b52dd95d8438172bca4b7ac1f32e6425387be4296c08d8bddb2098"}}`
    key, err := keystore.DecryptKey([]byte(keyJson), "12345")
    if err != nil {
        return nil, err
    }

    sig, err := crypto.Sign(message, key.PrivateKey)
    return sig, err
}

func (ce CryptExecutor) Encrypt(plainText []byte, pubkey string, password string) ([]byte, error) {
    //hardcode here, will call jeremy's library later
    return plainText, nil
}

func (ce CryptExecutor) ReEncryptWithAnotherPubKey(cipherText []byte, pubKey1 string,
                                        pubKey2 string, password string) ([]byte, error) {
    //hardcode here, will call jeremy's library later
    return cipherText, nil
}

